package controller

import (
	"fmt"
	"net/url"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber"
	models "github.com/scraper_v2/models"
)

func PirateBayController(fibCon *fiber.Ctx) {
	param := url.Values{}
	param.Add("q", fibCon.Query("search"))
	url := fmt.Sprintf("https://piratebaylive.com/search?%s&cat%5B%5D=&search=Pirate+Search", param.Encode())
	c := colly.NewCollector()
	var infos = make([]models.TorrentInfo, 0)
	var repo models.TorrentRepo = models.TorrentRepo{}
	var ti models.TorrentInfo = models.TorrentInfo{}
	c.OnHTML("body", func(e *colly.HTMLElement) {
		selector := e.DOM.Find("#st")
		if selector.Length() > 1 {
			e.ForEach("#st", func(i int, e *colly.HTMLElement) {
				ti.Name = e.ChildText("span.list-item.item-name.item-title")
				ti.Seeders = e.ChildText("span.list-item.item-seed")
				ti.Leechers = e.ChildText("span.list-item.item-leech")
				ti.Date = e.ChildText("span.list-item.item-uploaded")
				ti.Size = e.ChildText("span.list-item.item-size")
				ti.Uploader = e.ChildText("span.list-item.item-user")
				ti.Magnet = e.ChildAttr("span.item-icons a", "href")
				ti.Url = e.ChildAttr("span.list-item.item-name.item-title a", "href")
				ti.Website = "Pirate Bay"
				ti.TorrentFileUrl = ""
				infos = append(infos, ti)
			})
		}
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
