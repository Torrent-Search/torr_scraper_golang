package routes

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func Skytorrents(ginCon *gin.Context) {
	// search := strings.ReplaceAll(strings.TrimSpace(ginCon.Query("search")), " ", "%20")
	param := url.Values{}
	param.Add("query", ginCon.Query("search"))
	url := fmt.Sprint("https://www.skytorrents.lol/?", param.Encode())

	c := colly.NewCollector()
	infos := make([]TorrentInfo, 0)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		tr := TorrentInfo{}
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

			if isMagnet(magnet_selector_with_child.Eq(3).AttrOr("href", "")) {
				tr.TorrentFileUrl = magnet_selector_with_child.Eq(2).AttrOr("href", "")
				tr.Magnet = magnet_selector_with_child.Eq(3).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(4).AttrOr("href", "")) {
				tr.TorrentFileUrl = magnet_selector_with_child.Eq(3).AttrOr("href", "")
				tr.Magnet = magnet_selector_with_child.Eq(4).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(5).AttrOr("href", "")) {
				tr.TorrentFileUrl = magnet_selector_with_child.Eq(4).AttrOr("href", "")
				tr.Magnet = magnet_selector_with_child.Eq(5).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(6).AttrOr("href", "")) {
				tr.TorrentFileUrl = magnet_selector_with_child.Eq(5).AttrOr("href", "")
				tr.Magnet = magnet_selector_with_child.Eq(6).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(7).AttrOr("href", "")) {
				tr.TorrentFileUrl = magnet_selector_with_child.Eq(6).AttrOr("href", "")
				tr.Magnet = magnet_selector_with_child.Eq(7).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(8).AttrOr("href", "")) {
				tr.TorrentFileUrl = magnet_selector_with_child.Eq(7).AttrOr("href", "")
				tr.Magnet = magnet_selector_with_child.Eq(8).AttrOr("href", "")
			}
			tr.Url = "https://www.skytorrents.lol/" + e.ChildAttr("td:nth-child(1) a:nth-child(1)", "href")
			tr.Website = "Skytorrents"
			tr.Uploader = "--"
			tr.TorrentFileUrl = "https:" + tr.TorrentFileUrl
			infos = append(infos, tr)
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnScraped(func(r *colly.Response) {
		if len(infos) > 0 {
			ginCon.JSON(200, TorrentRepo{infos})
		} else {
			ginCon.AbortWithStatus(204)
		}
	})
	c.Visit(url)
}
