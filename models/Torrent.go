package models

type TorrentInfo struct {
	Name           string `json:"name"`
	Url            string `json:"torrent_url"`
	Seeders        string `json:"seeders"`
	Leechers       string `json:"leechers"`
	Date           string `json:"upload_date"`
	Size           string `json:"size"`
	Uploader       string `json:"uploader"`
	Magnet         string `json:"magnet"`
	Website        string `json:"website"`
	TorrentFileUrl string `json:"torrent_file_url"`
}

type TorrentRepo struct {
	Data *[]TorrentInfo `json:"data"`
}
