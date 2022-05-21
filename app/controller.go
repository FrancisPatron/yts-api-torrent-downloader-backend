package app

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

	fmt.Println(err)

	return movies
}
