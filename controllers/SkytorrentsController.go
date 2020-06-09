package controller

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber"
	models "github.com/scraper_v2/models"
)

func SkytorrentsController(fibCon *fiber.Ctx) {
	// search := strings.ReplaceAll(strings.TrimSpace(fibCon.Query("search")), " ", "%20")
	param := url.Values{}
	param.Add("query", fibCon.Query("search"))
	url := fmt.Sprint("https://www.skytorrents.lol/?", param.Encode())

	c := colly.NewCollector()
	var infos = make([]models.TorrentInfo, 0)
	var repo models.TorrentRepo = models.TorrentRepo{}
	var tr models.TorrentInfo = models.TorrentInfo{}
	c.OnHTML("body", func(e *colly.HTMLElement) {

		if e.DOM.Find("tr").Length() == 0 {
			return
		}
		e.ForEach("tr", func(i int, e *colly.HTMLElement) {
			if i == 0 {
				return
			}
			tr.Name = e.ChildText("td:nth-child(1) a:nth-child(1)")
			tr.Seeders = e.ChildText("td:nth-child(5)")
			tr.Leechers = e.ChildText("td:nth-child(6)")
			tr.Date = e.ChildText("td:nth-child(4)")
			tr.Size = e.ChildText("td:nth-child(2)")

			magnet_selector_with_child := e.DOM.Find("td:nth-child(1)").Children()

			if strings.HasPrefix(magnet_selector_with_child.Eq(3).AttrOr("href", ""), "magnet") {
				tr.TorrentFileUrl = magnet_selector_with_child.Eq(2).AttrOr("href", "")
				tr.Magnet = magnet_selector_with_child.Eq(3).AttrOr("href", "")
			} else if strings.HasPrefix(magnet_selector_with_child.Eq(4).AttrOr("href", ""), "magnet") {
				tr.TorrentFileUrl = magnet_selector_with_child.Eq(3).AttrOr("href", "")
				tr.Magnet = magnet_selector_with_child.Eq(4).AttrOr("href", "")
			} else if strings.HasPrefix(magnet_selector_with_child.Eq(5).AttrOr("href", ""), "magnet") {
				tr.TorrentFileUrl = magnet_selector_with_child.Eq(4).AttrOr("href", "")
				tr.Magnet = magnet_selector_with_child.Eq(5).AttrOr("href", "")
			} else if strings.HasPrefix(magnet_selector_with_child.Eq(6).AttrOr("href", ""), "magnet") {
				tr.TorrentFileUrl = magnet_selector_with_child.Eq(5).AttrOr("href", "")
				tr.Magnet = magnet_selector_with_child.Eq(6).AttrOr("href", "")
			} else if strings.HasPrefix(magnet_selector_with_child.Eq(7).AttrOr("href", ""), "magnet") {
				tr.TorrentFileUrl = magnet_selector_with_child.Eq(6).AttrOr("href", "")
				tr.Magnet = magnet_selector_with_child.Eq(7).AttrOr("href", "")
			} else if strings.HasPrefix(magnet_selector_with_child.Eq(8).AttrOr("href", ""), "magnet") {
				tr.TorrentFileUrl = magnet_selector_with_child.Eq(7).AttrOr("href", "")
				tr.Magnet = magnet_selector_with_child.Eq(8).AttrOr("href", "")
			}
			tr.Url = "https://www.skytorrents.lol/" + e.ChildAttr("td:nth-child(1) a:nth-child(1)", "href")
			tr.Website = "Skytorrents"
			tr.Uploader = "N/A"
			tr.TorrentFileUrl = "https:" + tr.TorrentFileUrl
			infos = append(infos, tr)
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
