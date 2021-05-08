package main

import "net/http"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func getContentHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func storeContentHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
