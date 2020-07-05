package music

import "encoding/json"

func UnmarshalTrending(data []byte) (Trending, error) {
	var r Trending
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Trending) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Trending struct {
	TrendingItem []TrendingItem `json:"new_trending"`
}

type TrendingItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Image string `json:"image"`
}
