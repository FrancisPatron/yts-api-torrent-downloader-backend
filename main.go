package main

import (
	"backend/app"
	"log"
	"net/http"

	"github.com/cenkalti/rain/torrent"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	router := mux.NewRouter()

	app.Downloads = make(map[string]*torrent.Torrent)

	router.HandleFunc("/api/movies/", app.GetMovies).Methods("GET")
	router.HandleFunc("/api/download/", app.DownloadTorrent).Methods("POST")

	logrus.Info("http://localhost:8080/api/")
	log.Fatal(http.ListenAndServe(":8080", router))
}
