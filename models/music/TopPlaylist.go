package music

import "encoding/json"

func UnmarshalTopPlaylist(data []byte) (TopPlaylist, error) {
	var r TopPlaylist
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *TopPlaylist) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type TopPlaylist struct {
	TopPlaylistItem []TopPlaylistItem `json:"top_playlists"`
}

type TopPlaylistItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Image string `json:"image"`
}
