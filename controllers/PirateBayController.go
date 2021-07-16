package controller

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber"
	models "github.com/scraper_v2/models"
)

func PirateBayController(fibCon *fiber.Ctx) {

	var page string
	if fibCon.Query("page") == "" {
		page = "1"
	} else {
		page = fibCon.Query("page")
	}
	url := fmt.Sprintf("https://thepiratebay.zone/search/%s/%s/99/0", url.PathEscape(fibCon.Query("search")), page)

	c := colly.NewCollector()

	var infos = make([]models.TorrentInfo, 0)
	var repo models.TorrentRepo = models.TorrentRepo{}
	var ti models.TorrentInfo = models.TorrentInfo{}
	c.OnHTML("body", func(e *colly.HTMLElement) {
		selector := e.DOM.Find("tr")
		if selector.Length() > 1 {

			e.ForEach("tr", func(i int, e *colly.HTMLElement) {
				if i == 0 {
					return
				}
				ti.Name = e.ChildText("td:nth-child(2) div a")
				ti.Seeders = e.ChildText("td:nth-child(3)")
				ti.Leechers = e.ChildText("td:nth-child(4)")
				file_info := e.ChildText("font.detDesc")
				upload_date_temp := strings.Split(file_info, ",")
				if len(upload_date_temp) == 1 {
					return
				}
				ti.Date = replace(replace(upload_date_temp[0], "Uploaded ", ""), " ", "-")
				ti.Size = replace(upload_date_temp[1], " Size ", "")
				ti.Uploader = replace(upload_date_temp[2], " ULed by ", "")
				ti.Magnet = e.ChildAttr("td:nth-child(2) a:nth-child(2)", "href")
				ti.Url = e.ChildAttr("td:nth-child(2) div a", "href")
				ti.Website = "Pirate Bay"
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

func replace(str string, new string, old string) string {
	return strings.ReplaceAll(str, new, old)
}
