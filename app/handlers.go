package app

import (
	"backend/yts"
	"encoding/json"
	"net/http"
	"time"

	"github.com/cenkalti/rain/torrent"
	log "github.com/sirupsen/logrus"
)

var Downloads map[string]*torrent.Torrent

func GetMovies(w http.ResponseWriter, r *http.Request) {
	searchParams := r.URL.Query()
	res := yts.ListMovies(searchParams)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func DownloadTorrent(w http.ResponseWriter, r *http.Request) {
	hash := r.URL.Query()["hash"][0]
	title := r.URL.Query()["title"][0]
	magnetLink := yts.CreateMagnetLink(hash, title)
	log.Infof("starting download for: %s(hash:%s) magnet: %s", title, hash, magnetLink)
	config := torrent.DefaultConfig
	config.DataDir = "/tmp/torrents"
	ses, err := torrent.NewSession(config)
	if err != nil {
		log.Error(err)
	}
	tor, err := ses.AddURI(magnetLink, nil)
	if err != nil {
		log.Error(err)
	}
	Downloads[hash] = tor // add to downloads map
	for range time.Tick(time.Second) {
		s := tor.Stats()
		down_percentage := float32(s.Bytes.Completed) / float32(s.Bytes.Total+1)
		log.Infof("status: %s Downloaded: %f %%", s.Status, down_percentage*100)
	}
	log.Infof("completed download for: %s(hash:%s) magnet: %s", title, hash, magnetLink)
	delete(Downloads, hash)
}

func DownloadSRT(w http.ResponseWriter, r *http.Request) {

}

func GetDownloads(w http.ResponseWriter, r *http.Request) {

}
