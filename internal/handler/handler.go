package handler

import (
	"fmt"
	"github.com/KaiserWerk/mininotes-server/internal/databaseservice"
	"io/ioutil"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func GetContentHandler(w http.ResponseWriter, r *http.Request) {
	id, secret, ok := r.BasicAuth()
	if !ok {
		fmt.Println("Basic auth failed for id " + id)
		http.Error(w, "basic auth failed", http.StatusUnauthorized)
		return
	}

	ds := databaseservice.Get()
	user, err := ds.GetUser(id)
	if err != nil {
		fmt.Printf("could not find user for id %s: %s\n", id, err.Error())
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if secret != user.Secret {
		fmt.Printf("wrong secret for user with id %s\n", id)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Write([]byte(user.Content))
}

func StoreContentHandler(w http.ResponseWriter, r *http.Request) {
	// TODO duplicate code
	id, secret, ok := r.BasicAuth()
	if !ok {
		fmt.Println("Basic auth failed for id " + id)
		http.Error(w, "basic auth failed", http.StatusUnauthorized)
		return
	}

	ds := databaseservice.Get()
	user, err := ds.GetUser(id)
	if err != nil {
		fmt.Printf("could not find user for id %s: %s\n", id, err.Error())
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if secret != user.Secret {
		fmt.Printf("wrong secret for user with id %s\n", id)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	cont, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not read request body: %s\n", err.Error())
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	_ = r.Body.Close()

	user.Content = string(cont)
	err = ds.UpdateUser(*user)
	if err != nil {
		fmt.Printf("could not update user: %s\n", err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
