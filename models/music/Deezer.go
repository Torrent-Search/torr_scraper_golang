package music

import "encoding/json"

type Deezer []DeezerItem

func UnmarshalDeezer(data []byte) (Deezer, error) {
	var r Deezer
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Deezer) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DeezerItem struct {
	ID        int64  `json:"id"`       
	Title     string `json:"title"`    
	Source    string `json:"source"`   
	Duration  int64  `json:"duration"` 
	Thumbnail string `json:"thumbnail"`
	Artist    string `json:"artist"`   
	URL       string `json:"url"`      
}
