package controller

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber"
	models "github.com/scraper_v2/models"
)

func LimetorrentsController(fibCon *fiber.Ctx) {
	var (
		search    string = fibCon.Query("search")
		searchUrl string
		pageNo    string             = fibCon.Query("page")
		c         *colly.Collector   = colly.NewCollector()
		infos                        = make([]models.TorrentInfo, 0)
		repo      models.TorrentRepo = models.TorrentRepo{}
		ti        models.TorrentInfo = models.TorrentInfo{}
	)
	search = url.PathEscape(search)
	if pageNo == "" {
		pageNo = "1"
	}
	searchUrl = fmt.Sprintf("https://www.limetorrents.pro/search/all/%s/%s/", search, pageNo)

	c.OnHTML("body", func(e *colly.HTMLElement) {

		e.ForEach("table.table2 tbody tr", func(i int, e *colly.HTMLElement) {
			if i == 0 {
				return
			}
			ti.Name = e.ChildText("td:nth-child(1)")
			ti.Seeders = e.ChildText("td:nth-child(4)")
			ti.Leechers = e.ChildText("td:nth-child(5)")
			ti.Date = strings.Split(e.ChildText("td:nth-child(2)"), " - ")[0]
			ti.Size = e.ChildText("td:nth-child(3)")
			ti.Magnet = ""
			ti.Url = "https://www.limetorrents.pro" + e.ChildAttr("td.tdleft div.tt-name a:nth-child(2)", "href")
			ti.Website = "Limetorrents"
			ti.Uploader = "N/A"
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
			fibCon.Status(200)
		}
	})
	c.Visit(searchUrl)

}
func LimetorrentsMgController(fibCon *fiber.Ctx) {
	url := fibCon.Query("url")
	c := colly.NewCollector()
	var magnet, torrentfile string
	c.OnHTML("body", func(e *colly.HTMLElement) {
		magnet = e.ChildAttr("#content > div:nth-child(6) > div:nth-child(1) > div > div:nth-child(13) > div > p > a", "href")
		torrentfile = e.ChildAttr("body > div > table:nth-child(6) > tbody > tr:nth-child(3) > td > span > a", "href")
		print(url)
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
