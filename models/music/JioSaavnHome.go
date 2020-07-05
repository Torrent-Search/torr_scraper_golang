package music

import "encoding/json"

func UnmarshalJioSaavnHome(data []byte) (JioSaavnHome, error) {
	var r JioSaavnHome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *JioSaavnHome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type JioSaavnHome struct {
	Charts      []ChartItem       `json:"charts"`
	Trending    []TrendingItem    `json:"trending"`
	TopPlaylist []TopPlaylistItem `json:"trending"`
}
