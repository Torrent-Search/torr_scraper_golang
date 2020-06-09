package models

type Data struct {
	MovieCount int     `json:"movie_count"`
	PageNumber int     `json:"int"`
	Movies     []Movie `json:"movies"`
}

// Result represents the response from the API
type Result struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          Data   `json:"data"`
}
type Movie struct {
	DateUploaded     string    `json:"date_uploaded"`
	DateUploadedUnix int64     `json:"date_uploaded_unix"`
	Genres           []string  `json:"genres"`
	ID               int       `json:"id"`
	ImdbID           string    `json:"imdb_code"`
	Language         string    `json:"language"`
	MediumCover      string    `json:"medium_cover_image"`
	Rating           float64   `json:"rating"`
	Runtime          int       `json:"runtime"`
	SmallCover       string    `json:"small_cover_image"`
	State            string    `json:"state"`
	Title            string    `json:"title"`
	TitleLong        string    `json:"title_long"`
	Torrents         []Torrent `json:"torrents"`
	Year             int       `json:"year"`
}
type Torrent struct {
	DateUploaded     string `json:"date_uploaded"`
	DateUploadedUnix int64  `json:"date_uploaded_unix"`
	Hash             string `json:"hash"`
	Peers            int    `json:"peers"`
	Quality          string `json:"quality"`
	Seeds            int    `json:"seeds"`
	Size             string `json:"size"`
	SizeBytes        int    `json:"size_bytes"`
	URL              string `json:"url"`
}
