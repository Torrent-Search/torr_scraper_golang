package controller

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber"
	helper "github.com/scraper_v2/helper"
	models "github.com/scraper_v2/models"
)

func Controller1337x(fibCon *fiber.Ctx) {

	var (
		search string = fibCon.Query("search")
		page   string = fibCon.Query("page")
		url    string

		c     *colly.Collector     = colly.NewCollector()
		infos []models.TorrentInfo = make([]models.TorrentInfo, 0)
		repo  models.TorrentRepo   = models.TorrentRepo{}
		ti    models.TorrentInfo   = models.TorrentInfo{}
	)

	if page == "" {
		page = "1"
	}

	url = fmt.Sprintf("https://1337x.to/search/%s/%s/", search, page)

	c.OnHTML("body", func(e *colly.HTMLElement) {
		if e.DOM.Find("tr").Length() == 0 {
			return
		}
		e.ForEach("tr", func(i int, e *colly.HTMLElement) {
			if i == 0 {
				return
			}
			ti.Name = e.ChildText("td.coll-1.name")
			ti.Seeders = e.ChildText("td.coll-2.seeds")
			ti.Leechers = e.ChildText("td.coll-3.leeches")
			ti.Date = e.ChildText("td.coll-date")
			ti.Size = e.DOM.Find("td:nth-child(5)").Clone().Children().Remove().End().Text()
			ti.Uploader = e.ChildText("td:nth-child(6)")
			ti.Url =
				"https://1337x.to" +
					e.ChildAttr("td.coll-1.name > a:nth-child(2)", "href")
			ti.Website = "1337x"
			ti.TorrentFileUrl = ""
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
	c.SetProxyFunc(helper.GetProxy())
	c.Visit(url)
}

func Controller1337xMg(fibCon *fiber.Ctx) {
	url := fibCon.Query("url")
	var magnet, torrentfile string
	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		magnet = e.ChildAttr("div.clearfix ul li a", "href")
		torrentfile = e.ChildAttr("div.clearfix ul li.dropdown ul li:nth-child(1) a", "href")
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
	c.SetProxyFunc(helper.GetProxy())
	c.Visit(url)
}
