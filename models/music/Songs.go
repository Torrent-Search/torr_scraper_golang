package music

type Songs struct {
	Data     []SongsData `json:"data"`
	Position int64       `json:"position"`
}

type SongsData struct {
	ID       string           `json:"id"`
	Title    string           `json:"title"`
	Image    string           `json:"image"`
	Album    string           `json:"album"`
	URL      string           `json:"url"`
	Type     string           `json:"type"`
	Ctr      int64            `json:"ctr"`
	MoreInfo SongItemMoreInfo `json:"more_info"`
}

type SongItemMoreInfo struct {
	PrimaryArtists string `json:"primary_artists"`
	Singers        string `json:"singers"`
}

type SongsDataWithLink struct {
	ID        string `json:"id"`
	Title     string `json:"song"`
	Image     string `json:"image"`
	Album     string `json:"album"`
	AlbumID   string `json:"albumid"`
	URL       string `json:"encrypted_media_url"`
	Year      string `json:"year"`
	Duration  string `json:"duration"`
	Singers   string `json:"singers"`
	Is320kbps string `json:"320kbps"`
}
