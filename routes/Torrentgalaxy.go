package routes

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func Torrentgalaxy(ginCon *gin.Context) {
	// search := strings.ReplaceAll(strings.TrimSpace(ginCon.Query("search")), " ", "%20")
	param := url.Values{}
	param.Add("search", ginCon.Query("search"))
	url := fmt.Sprintf("https://torrentgalaxy.to/torrents.php?%s", param.Encode())
	c := colly.NewCollector()
	infos := make([]TorrentInfo, 0)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		ti := TorrentInfo{}
		print(e.DOM.Find("div.tgxtablerow").Length())
		if e.DOM.Find("div.tgxtablerow").Length() == 0 {
			return
		}
		e.ForEach("div.tgxtablerow", func(i int, e *colly.HTMLElement) {
			ti.Name = e.ChildText("div:nth-child(4)")
			ti.Seeders = e.ChildText("div:nth-child(11) span font:nth-child(1)")
			ti.Leechers = e.ChildText("div:nth-child(11) span font:nth-child(2)")
			ti.Date = strings.Split(e.ChildText("div:nth-child(12)"), " ")[0]
			ti.Size = e.ChildText("div:nth-child(8)")
			ti.Uploader = e.ChildText("div:nth-child(7)")
			ti.Magnet = e.DOM.Find("#click").Next().Find("a:nth-child(2)").AttrOr("href", "")
			ti.Url = "https://torrentgalaxy.to" + e.ChildAttr("div:nth-child(4) a", "href")
			ti.Website = "Torrent Galaxy"
			ti.TorrentFileUrl = e.DOM.Find("#click").Next().Find("a:nth-child(1)").AttrOr("href", "")
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
