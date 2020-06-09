package controller

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gofiber/fiber"
	models "github.com/scraper_v2/models"
)

func KickassController(fibCon *fiber.Ctx) {
	search := url.PathEscape(fibCon.Query("search"))
	url := fmt.Sprint("https://kickasstorrents.to/usearch/", search)
	c := colly.NewCollector()
	var infos = make([]models.TorrentInfo, 0)
	var repo models.TorrentRepo = models.TorrentRepo{}
	var ti models.TorrentInfo = models.TorrentInfo{}
	c.OnHTML("body", func(e *colly.HTMLElement) {

		if e.DOM.Find("span[itemprop=name]").Length() == 0 {
			e.ForEach("tr.odd , tr.even", func(i int, e *colly.HTMLElement) {
				ti.Name = e.ChildText(".cellMainLink")
				ti.Seeders = e.ChildText("td:nth-child(5)")
				ti.Leechers = e.ChildText("td:nth-child(6)")
				ti.Date = e.ChildText("td:nth-child(4)")
				ti.Size = e.ChildText("td:nth-child(2)")
				ti.Magnet = ""
				ti.Url = "https://kickasstorrents.to" + e.ChildAttr(".cellMainLink", "href")
				ti.Website = "Kickass"
				ti.Uploader = e.ChildText("td:nth-child(3)")
				ti.TorrentFileUrl = e.ChildAttr("td:nth-child(3) a:nth-child(2)", "href")
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
func KickassMgController(fibCon *fiber.Ctx) {
	url := fibCon.Query("url")
	c := colly.NewCollector()
	var magnet, torrentfile string
	c.OnHTML("body", func(e *colly.HTMLElement) {
		magnet = e.ChildAttr("a.kaGiantButton", "href")
		torrentfile = ""
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnScraped(func(r *colly.Response) {
		if strings.HasPrefix(magnet, "magnet") {
			fibCon.Status(200).JSON(fiber.Map{"magnet": magnet, "torrentFile": torrentfile})
		} else {
			fibCon.Status(204)
		}
	})
	c.Visit(url)
}
