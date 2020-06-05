package routes

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func Torr_1337x(ginCon *gin.Context) {
	search := strings.ReplaceAll(strings.TrimSpace(ginCon.Query("search")), " ", "%20")
	url := fmt.Sprintf("https://1337x.to/search/%s/1/", search)
	c := colly.NewCollector()
	infos := make([]TorrentInfo, 0)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		ti := TorrentInfo{}
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
			ginCon.JSON(200, TorrentRepo{infos})
		} else {
			ginCon.AbortWithStatus(204)
		}
	})
	c.Visit(url)
}

func Torr_1337x_getMagnet(ginCon *gin.Context) {
	url := ginCon.Query("url")
	c := colly.NewCollector()
	var magnet, torrentfile string
	c.OnHTML("body", func(e *colly.HTMLElement) {
		magnet = e.ChildAttr("div.clearfix ul li a", "href")
		torrentfile = e.ChildAttr("div.clearfix ul li.dropdown ul li:nth-child(1) a", "href")
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnScraped(func(r *colly.Response) {
		if strings.HasPrefix(magnet, "magnet") {
			ginCon.JSON(200, gin.H{"magnet": magnet, "torrentFile": torrentfile})
		} else {
			ginCon.AbortWithStatus(204)
		}
	})
	c.Visit(url)

}
