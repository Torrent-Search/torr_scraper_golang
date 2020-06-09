package controller

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/scraper_v2/helper"
	models "github.com/scraper_v2/models"
)

func YtsController(fibCon *fiber.Ctx) {
	param := url.Values{}
	param.Add("query_term", fibCon.Query("search"))
	url := fmt.Sprintf("https://yts.mx/api/v2/list_movies.json?%s", param.Encode())
	res, _ := helper.GetResponse(url)

	var data models.Result
	err := json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		defer res.Body.Close()
		fibCon.Status(500)
	}
	var infos = make([]models.TorrentInfo, 0)
	var repo models.TorrentRepo = models.TorrentRepo{}
	var tr models.TorrentInfo = models.TorrentInfo{}

	movies := data.Data.Movies
	if data.Data.MovieCount > 0 {
		for _, obj := range movies {
			for _, torr_obj := range obj.Torrents {
				tr.Name = obj.Title + " " + torr_obj.Quality
				tr.Seeders = strconv.Itoa(torr_obj.Seeds)
				tr.Leechers = strconv.Itoa(torr_obj.Peers)
				tr.Date = strings.Split(torr_obj.DateUploaded, " ")[0]
				tr.Url = torr_obj.URL
				tr.Uploader = "YTS"
				tr.Website = "YTS"
				tr.Size = torr_obj.Size
				tr.Magnet = helper.GenerateYtsMagnet(torr_obj.Hash)
				tr.TorrentFileUrl = ""
				infos = append(infos, tr)
			}
		}
		defer res.Body.Close()
		repo.Data = &infos
		fibCon.Status(200).JSON(repo)
	} else {
		defer res.Body.Close()
		fibCon.Status(204)
	}
}
