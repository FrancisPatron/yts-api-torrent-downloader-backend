package app

import (
	"encoding/json"
	"net/http"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	searchParams := r.URL.Query()
	res := ListMovies(searchParams)
	json.NewEncoder(w).Encode(res)
}
