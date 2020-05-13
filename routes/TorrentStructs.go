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
