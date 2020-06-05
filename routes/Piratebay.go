package routes

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func PirateBay(ginCon *gin.Context) {
	// search := strings.ReplaceAll(strings.TrimSpace(ginCon.Query("search")), " ", "%20")
	param := url.Values{}
	param.Add("q", ginCon.Query("search"))
	url := fmt.Sprintf("https://piratebaylive.com/search?%s&cat%5B%5D=&search=Pirate+Search", param.Encode())
	c := colly.NewCollector()
	infos := make([]TorrentInfo, 0)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		ti := TorrentInfo{}
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
			ginCon.JSON(200, TorrentRepo{infos})
		} else {
			ginCon.AbortWithStatus(204)
		}
	})
	c.Visit(url)
}
