package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"mininotes-server/assets"
	"mininotes-server/config"
	"mininotes-server/databaseservice"
	"mininotes-server/helper"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	port = "11100"
)

func main() {
	configFilePtr := flag.String("c", "", "The configuration file to load")
	portPtr := flag.String("p", "", "The port to use")
	flag.Parse()

	if *configFilePtr != "" {
		config.SetConfigFile(*configFilePtr)
	}
	if *portPtr != "" {
		port = *portPtr
	}

	if !helper.FileExists(*configFilePtr) {
		configAssets := assets.GetConfigFiles()
		cont, err := configAssets.ReadFile("config/config.dist.yml")
		if err != nil {
			panic("no embedded config dist file found: " + err.Error())
		}
		err = ioutil.WriteFile(*configFilePtr, cont, 0744)
		if err != nil {
			panic("could not write config dist file to disk: " + err.Error())
		}
	}

	ds := databaseservice.Get()
	err := ds.AutoMigrate()
	if err != nil {
		panic("AutoMigrate panic: " + err.Error())
	}

	host := fmt.Sprintf(":%s", *portPtr)

	router := mux.NewRouter()
	router.HandleFunc("/", helloHandler)
	router.HandleFunc("/get", getContentHandler)
	router.HandleFunc("/store", storeContentHandler)

	// catch ctrl+c for graceful shutdown
	notify := make(chan os.Signal)
	signal.Notify(notify, os.Interrupt)

	srv := &http.Server{
		Addr: 				host,
		Handler:            router,
		ReadTimeout:		2 * time.Second,
		WriteTimeout:       2 * time.Second,
		IdleTimeout:        3 * time.Second,
		ReadHeaderTimeout:  2 * time.Second,
	}

	go func() {
		<-notify
		fmt.Println("Initiating graceful shutdown...")
		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()
		// do stuff before exiting here

		srv.SetKeepAlivesEnabled(false)
		err := srv.Shutdown(ctx)
		if err != nil {
			panic("Could not gracefully shut down server: " + err.Error())
		}
	}()


	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("server error: %v\n", err.Error())
	}
	fmt.Println("Server shutdown complete. Have a nice day!")
}