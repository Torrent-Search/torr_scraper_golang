package music

import "encoding/json"

func UnmarshalCharts(data []byte) (Charts, error) {
	var r Charts
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Charts) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Charts struct {
	Items []ChartItem `json:"charts"`
}

type ChartItem struct {
	Image string `json:"image"`
	Title string `json:"title"`
	ID    string `json:"id"`
	Type  string `json:"type"`
}
