package yts

type Torrent struct {
	Url              string  `json:"url"`
	Hash             string  `json:"hash"`
	Quality          string  `json:"quality"`
	Type             string  `json:"type"`
	Seeds            int     `json:"seeds"`
	Peers            int     `json:"peers"`
	Size             string  `json:"size"`
	SizeBytes        float64 `json:"size_bytes"`
	DateUploaded     string  `json:"date_uploaded"`
	DateUploadedUnix int     `json:"date_uploaded_unix"`
}

type Movie struct {
	ID                      int       `json:"id"`
	Uul                     string    `json:"url"`
	ImdbCode                string    `json:"imdb_code"`
	Title                   string    `json:"title"`
	TitleEng                string    `json:"title_english"`
	TitleLong               string    `json:"title_long"`
	Slug                    string    `json:"slug"`
	Year                    int       `json:"year"`
	Raiting                 float32   `json:"rating"`
	Runtime                 int       `json:"runtime"`
	Generes                 []string  `json:"genres"`
	Summary                 string    `json:"summary"`
	DescriptionFull         string    `json:"description_full"`
	Synopsis                string    `json:"synopsis"`
	YtTrailerCode           string    `json:"yt_trailer_code"`
	Language                string    `json:"language"`
	MpaRating               string    `json:"mpa_rating"`
	BackgroundImage         string    `json:"background_image"`
	BackgroundImageOriginal string    `json:"background_image_original"`
	SmallCoverImage         string    `json:"small_cover_image"`
	MediumCoverImage        string    `json:"medium_cover_image"`
	LargeCoverImage         string    `json:"large_cover_image"`
	State                   string    `json:"state"`
	Torrents                []Torrent `json:"torrents"`
}

type MetaData struct {
	ServerTime     int    `json:"server_time"`
	ServerTimezone string `json:"server_timezone"`
	ApiVersion     int    `json:"api_version"`
	ExecutionTime  string `json:"execution_time"`
}

type YTSResponse struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          struct {
		Movies []Movie `json:"movies"`
	} `json:"data"`
	Meta MetaData `json:"@meta"`
}
