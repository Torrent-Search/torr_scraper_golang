package routes

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Yts(c *gin.Context) {
	param := url.Values{}
	param.Add("query_term", c.Query("search"))
	url := fmt.Sprintf("https://yts.mx/api/v2/list_movies.json?%s", param.Encode())
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 20 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	var client = &http.Client{
		Timeout:   time.Second * 20,
		Transport: netTransport,
	}
	request, _ := http.NewRequest("GET", url, nil)
	res, _ := client.Do(request)

	var data Result
	err := json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		defer res.Body.Close()
		c.AbortWithStatus(500)
	}
	infos := make([]TorrentInfo, 0)

	movies := data.Data.Movies
	if data.Data.MovieCount > 0 {
		tr := TorrentInfo{}
		for _, obj := range movies {
			for _, torr_obj := range obj.Torrents {
				tr.Name = obj.Title + " " + torr_obj.Quality
				tr.Seeders = strconv.Itoa(torr_obj.Seeds)
				tr.Leechers = strconv.Itoa(torr_obj.Peers)
				tr.Date = torr_obj.DateUploaded
				tr.Url = torr_obj.URL
				tr.Uploader = "YTS"
				tr.Website = "YTS"
				tr.Size = torr_obj.Size
				tr.Magnet = getYts_mg(torr_obj.Hash)
				infos = append(infos, tr)
			}
		}
		defer res.Body.Close()
		c.JSON(200, TorrentRepo{infos})
	} else {
		defer res.Body.Close()
		c.AbortWithStatus(204)
	}
}
