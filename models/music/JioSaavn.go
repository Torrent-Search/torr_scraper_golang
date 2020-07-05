package music

import "encoding/json"

func UnmarshalJioSaavnQuery(data []byte) (JioSaavnQuery, error) {
	var r JioSaavnQuery
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *JioSaavnQuery) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type JioSaavnQuery struct {
	Albums Albums `json:"albums"`
	Songs  Songs  `json:"songs"`
}
