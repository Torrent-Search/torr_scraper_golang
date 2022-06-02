package controller

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/mmcdole/gofeed"
	helper "github.com/scraper_v2/helper"
	models "github.com/scraper_v2/models"
)

func TorrentdownloadsController(fibCon *fiber.Ctx) {
	param := url.Values{}
	param.Add("q", fibCon.Query("search"))
	url := fmt.Sprintf("https://www.torrentdownload.info/feed?%s", param.Encode())
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)
	items := feed.Items
	var infos = make([]models.TorrentInfo, 0)
	var repo models.TorrentRepo = models.TorrentRepo{}
	var tr models.TorrentInfo = models.TorrentInfo{}
	if len(items) > 0 {
		for _, obj := range items {
			desc := strings.Split(obj.Description, " ")

			tr.Name = obj.Title
			tr.Url = obj.Link
			tr.Date = strings.Trim(fmt.Sprint(strings.Split(obj.Published, " ")[0:3]), "[]")
			tr.Size = strings.Trim(fmt.Sprint(desc[1:3]), "[]")
			tr.Seeders = desc[4]
			tr.Leechers = desc[7]
			tr.Uploader = "N/A"
			tr.Magnet = helper.GenerateTorrentDownloadMagnet(desc[9])
			tr.Website = "Torrent Download"
			tr.TorrentFileUrl = ""
			infos = append(infos, tr)
		}
		repo.Data = &infos
		fibCon.Status(200).JSON(repo)
	} else {
		fibCon.Status(201)
	}
}
