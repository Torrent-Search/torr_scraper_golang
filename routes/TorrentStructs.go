package routes

import (
	"encoding/json"
	"net/http"
	"time"
)

type TorrentInfo struct {
	Name     string `json:"name"`
	Url      string `json:"torrent_url"`
	Seeders  string `json:"seeders"`
	Leechers string `json:"leechers"`
	Date     string `json:"upload_date"`
	Size     string `json:"size"`
	Uploader string `json:"uploader"`
	Magnet   string `json:"magnet"`
	Website  string `json:"website"`
}

type TorrentRepo struct {
	Data []TorrentInfo `json:"data"`
}

type TorrentResult struct {
	Title    string `json:"title"`
	Filename string `json:"filename"`
	Category string `json:"category"`
	Download string `json:"download"`
	Seeders  int    `json:"seeders"`
	Leechers int    `json:"leechers"`
	Size     uint64 `json:"size"`
	PubDate  string `json:"pubdate"`
	Ranked   int    `json:"ranked"`
	InfoPage string `json:"info_page"`
}

type TorrentResults []TorrentResult

type APIResponse struct {
	Torrents  json.RawMessage `json:"torrent_results"`
	Error     string          `json:"error"`
	ErrorCode int             `json:"error_code"`
}

// Token keeps token and it's expiration date.
type Token struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"-"`
}
type API struct {
	client          *http.Client
	Query           string
	APIToken        Token
	categories      []int
	appID           string
	reqDelay        time.Duration
	tokenExpiration time.Duration
	url             string
	maxRetries      int
}

// Movie represents the movies
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

// Torrent represents the quality for a torrent
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

// Data represents the data inside the response body
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
