package routes

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func Horriblesubs(ginCon *gin.Context) {
	param := url.Values{}
	param.Add("q", ginCon.Query("search"))
	url := fmt.Sprintf("https://nyaa.si/user/HorribleSubs?f=0&c=0_0&%s", param.Encode())
	c := colly.NewCollector()
	infos := make([]TorrentInfo, 0)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		ti := TorrentInfo{}
		e.ForEach("tr", func(i int, e *colly.HTMLElement) {
			if i == 0 {
				return
			}
			if e.DOM.Find("td:nth-child(2) a").Length() == 2 {
				ti.Name = e.DOM.Find("td:nth-child(2) a").Eq(1).Text()
			} else {
				ti.Name = e.ChildText("td:nth-child(2) a")
			}
			ti.Seeders = e.ChildText("td:nth-child(6)")
			ti.Leechers = e.ChildText("td:nth-child(7)")
			ti.Date = e.ChildText("td:nth-child(5)")
			ti.Size = e.ChildText("td:nth-child(4)")
			ti.Magnet = e.ChildAttr("td:nth-child(3) a:nth-child(2)", "href")
			ti.Url = "https://nyaa.si" + e.ChildAttr("td:nth-child(2) a", "href")
			ti.Website = "Nyaa"
			ti.Uploader = "N/A"
			ti.TorrentFileUrl = "https://nyaa.si" + e.ChildAttr("td:nth-child(3) a:nth-child(1)", "href")
			infos = append(infos, ti)
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
