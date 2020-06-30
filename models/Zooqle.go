package models

import "encoding/json"

func UnmarshalZooqle(data []byte) ([]ZooqleSearch, error) {
	var r []ZooqleSearch
	err := json.Unmarshal(data, &r)
	var sorted []ZooqleSearch
	for _, obj := range r {
		if obj.Type == "t" {
			sorted = append(sorted, obj)
		}
	}
	return sorted, err
}

type ZooqleSearch struct {
	Type  string `json:"t"`
	Name  string `json:"n"`
	ID    string `json:"id"`
	Image int    `json:"i"`
}

type ZooqleItem struct {
	Seasons    int             `json:"seasons_count"`
	SeasonList []ZooqleEpisode `json:"zooqle_Season"`
}

type ZooqleSeason struct {
	Season_No     int             `json:"season_no"`
	SeasonEpisode []ZooqleEpisode `json:"zooqle_Season"`
}

type ZooqleEpisode struct {
	Episode_No   string              `json:"episode_no"`
	Data_url     string              `json:"data_url"`
	EpisodesData []ZooqleEpisodeData `json:"episode_data"`
}

type ZooqleEpisodeData struct {
	Name     string `json:"name"`
	Seeders  string `json:"seeders"`
	Leechers string `json:"leechers"`
	Date     string `json:"upload_date"`
	Size     string `json:"size"`
	Magnet   string `json:"magnet"`
}

type ZooqleData struct {
	Data []ZooqleSeason `json:"data"`
}
