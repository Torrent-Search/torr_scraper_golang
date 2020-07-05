package music

type Albums struct {
	Data []AlbumsData `json:"data"`
}

type AlbumsData struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Music       string `json:"music"`
	URL         string `json:"url"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Ctr         int64  `json:"ctr"`
	Position    int64  `json:"position"`
}
