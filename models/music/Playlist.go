package music

import "encoding/json"

func UnmarshalPlaylist(data []byte) (Playlist, error) {
	var r Playlist
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Playlist) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Playlist struct {
	Listid   string              `json:"listid"`
	Listname string              `json:"listname"`
	Image    string              `json:"image"`
	Share    string              `json:"share"`
	Songs    []SongsDataWithLink `json:"songs"`
}
