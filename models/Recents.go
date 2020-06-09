package models

type Recents struct {
	Name       string `json:"name"`
	Url        string `json:"torrent_url"`
	ImgFileUrl string `json:"img_url"`
	Imdb_code  string `json:"imdbcode"`
}
type RecentRepo struct {
	Data *[]Recents `string:"data"`
}
type Imdb struct {
	Title      string        `json:"Title"`
	Year       string        `json:"Year"`
	Rated      string        `json:"Rated"`
	Released   string        `json:"Released"`
	Runtime    string        `json:"Runtime"`
	Genre      string        `json:"Genre"`
	Director   string        `json:"Director"`
	Writer     string        `json:"Writer"`
	Actors     string        `json:"Actors"`
	Plot       string        `json:"Plot"`
	Language   string        `json:"Language"`
	Country    string        `json:"Country"`
	Awards     string        `json:"Awards"`
	Poster     string        `json:"Poster"`
	Ratings    []interface{} `json:"Ratings"`
	Metascore  string        `json:"Metascore"`
	ImdbRating string        `json:"imdbRating"`
	ImdbVotes  string        `json:"imdbVotes"`
	ImdbID     string        `json:"imdbID"`
	Type       string        `json:"Type"`
	DVD        string        `json:"DVD"`
	BoxOffice  string        `json:"BoxOffice"`
	Production string        `json:"Production"`
	Website    string        `json:"Website"`
	Response   string        `json:"Response"`
}
