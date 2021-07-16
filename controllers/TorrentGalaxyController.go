package controller

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber"
	models "github.com/scraper_v2/models"
)

func TorrentGalaxyController(fibCon *fiber.Ctx) {
	param := url.Values{}
	param.Add("search", fibCon.Query("search"))
	if fibCon.Query("page") == "" {
		param.Add("page", "0")
	} else {
		param.Add("page", fibCon.Query("page"))
	}
	url := fmt.Sprintf("https://torrentgalaxy.to/torrents.php?parent_cat=&sort=id&order=desc&%s", param.Encode())
	c := colly.NewCollector()
	var infos = make([]models.TorrentInfo, 0)
	var repo models.TorrentRepo = models.TorrentRepo{}
	var ti models.TorrentInfo = models.TorrentInfo{}
	c.OnHTML("body", func(e *colly.HTMLElement) {

		print(e.DOM.Find("div.tgxtable").Length())
		if e.DOM.Find("div.tgxtable").Length() == 0 {
			return
		}
		e.ForEach("div.tgxtable div", func(i int, e *colly.HTMLElement) {
			ti.Name = e.ChildText("div:nth-child(4)")
			if ti.Name == "" || i == 0 {
				return
			}
			ti.Seeders = e.ChildText("div:nth-child(11) span font:nth-child(1)")
			ti.Leechers = e.ChildText("div:nth-child(11) span font:nth-child(2)")
			ti.Date = strings.Split(e.ChildText("div:nth-child(12)"), " ")[0]
			ti.Size = e.ChildText("div:nth-child(8)")
			ti.Uploader = e.ChildText("div:nth-child(7)")
			ti.Magnet = e.DOM.Find("#click").Next().Find("a:nth-child(2)").AttrOr("href", "")
			ti.Url = "https://torrentgalaxy.to" + e.ChildAttr("div:nth-child(4) a", "href")
			ti.Website = "Torrent Galaxy"
			ti.TorrentFileUrl = e.DOM.Find("#click").Next().Find("a:nth-child(1)").AttrOr("href", "")
			infos = append(infos, ti)
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnScraped(func(r *colly.Response) {
		if len(infos) > 0 {
			repo.Data = &infos
			fibCon.Status(200).JSON(repo)
		} else {
			fibCon.Status(204)
		}
	})
	c.Visit(url)
}
