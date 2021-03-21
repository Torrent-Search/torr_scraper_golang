package music

import "encoding/json"

type Yt []YtElement

func UnmarshalYt(data []byte) (Yt, error) {
	var r Yt
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Yt) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type YtElement struct {
	ID         string   `json:"id"`
	Thumbnails []string `json:"thumbnails"`
	Title      string   `json:"title"`
	URLSuffix  string   `json:"url_suffix"`
	Duration   string   `json:"duration"`
}
