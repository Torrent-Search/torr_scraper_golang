package controller

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber"
	"github.com/scraper_v2/models"
)

func ZooqleController(fibCon *fiber.Ctx) {
	param := url.Values{}
	param.Add("q", fibCon.Query("search"))
	var url string = fmt.Sprintf("https://zooqle.com/search?%s", param.Encode())
	var infos = make([]models.TorrentInfo, 0)
	var repo models.TorrentRepo = models.TorrentRepo{}
	var ti models.TorrentInfo = models.TorrentInfo{}
	var c *colly.Collector = colly.NewCollector()
	var seedLeechString string = ""
	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, e *colly.HTMLElement) {
			if i == 0 {
				return
			}
			ti.Name = e.ChildText("td:nth-child(2) a")
			seedLeechString = e.ChildAttr("td:nth-child(6) div", "title")
			ti.Seeders = strings.Split(seedLeechString, " ")[1]
			ti.Leechers = strings.Split(seedLeechString, " ")[4]
			ti.Date = e.ChildText("td:nth-child(5)")
			ti.Size = e.ChildText("td:nth-child(4) div div")
			ti.Magnet = e.ChildAttr("td:nth-child(3) ul li:nth-child(2) a", "href")
			infos = append(infos, ti)
		})
		// }
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
