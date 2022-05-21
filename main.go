package main

import (
	"backend/app"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/", app.GetMovies).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
