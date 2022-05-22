package yts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func ListMovies(params url.Values) YTSResponse {
	req, err := http.NewRequest("GET", "https://yts.torrentbay.to/api/v2/list_movies.json", nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	req.URL.RawQuery = params.Encode()

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var movies YTSResponse

	err = json.Unmarshal(body, &movies)
	if err != nil {
		log.Fatal(err)
	}

	return movies
}

func CreateMagnetLink(hash, title string) string {
	link := fmt.Sprintf("magnet:?xt=urn:btih:%s&dn=%s+%282022%29+%5B720p%5D+%5Byts.torrentbay.t%5D"+
		"&tr=udp://open.demonii.com:1337/announce"+
		"&tr=udp://tracker.openbittorrent.com:80"+
		"&tr=udp://tracker.coppersurfer.tk:6969"+
		"&tr=udp://glotorrents.pw:6969/announce"+
		"&tr=udp://tracker.opentrackr.org:1337/announce"+
		"&tr=udp://torrent.gresille.org:80/announce"+
		"&tr=udp://p4p.arenabg.com:1337"+
		"&tr=udp://tracker.leechers-paradise.org:6969", hash, title)
	return link
}
