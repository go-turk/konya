package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World, I'm Lab 1"))
	})
	http.ListenAndServe(":8080", router)
}
