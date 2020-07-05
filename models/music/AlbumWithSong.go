package music

import "encoding/json"

type AlbumWithSongs struct {
	Title            string              `json:"title"`
	Name             string              `json:"name"`
	Year             string              `json:"year"`
	ReleaseDate      string              `json:"release_date"`
	PrimaryArtists   string              `json:"primary_artists"`
	PrimaryArtistsID string              `json:"primary_artists_id"`
	Albumid          string              `json:"albumid"`
	PermaURL         string              `json:"perma_url"`
	Image            string              `json:"image"`
	Songs            []SongsDataWithLink `json:"songs"`
}

func UnmarshalAlbumWithSongs(data []byte) (AlbumWithSongs, error) {
	var r AlbumWithSongs
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *AlbumWithSongs) Marshal() ([]byte, error) {
	return json.Marshal(r)
}
