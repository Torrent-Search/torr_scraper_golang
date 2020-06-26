package controller

import (
	"fmt"
	"net/url"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber"
	models "github.com/scraper_v2/models"
)

func EztvController(fibCon *fiber.Ctx) {
	url := fmt.Sprintf("https://eztv.io/search/%s", url.PathEscape(fibCon.Query("search")))
	c := colly.NewCollector()
	var infos = make([]models.TorrentInfo, 0)
	var repo models.TorrentRepo = models.TorrentRepo{}
	var ti models.TorrentInfo = models.TorrentInfo{}
	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("tr.forum_header_border", func(i int, e *colly.HTMLElement) {

			ti.Name = e.ChildText("td:nth-child(2)")
			ti.Seeders = e.ChildText("td:nth-child(6)")
			ti.Leechers = "N/A"
			ti.Date = e.ChildText("td:nth-child(5)")
			ti.Size = e.ChildText("td:nth-child(4)")
			ti.Magnet = e.ChildAttr("td:nth-child(3) a:nth-child(1)", "href")
			ti.Url = "https://eztv.io" + e.ChildAttr("td:nth-child(2) a", "href")
			ti.Website = "EZTV"
			ti.Uploader = "N/A"
			ti.TorrentFileUrl = e.ChildAttr("td:nth-child(3) a:nth-child(2)", "href")
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
