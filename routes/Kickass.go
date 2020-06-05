package routes

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func Kickass(ginCon *gin.Context) {
	search := strings.ReplaceAll(strings.TrimSpace(ginCon.Query("search")), " ", "%20")
	url := fmt.Sprint("https://kickasstorrents.to/usearch/", search)
	c := colly.NewCollector()
	infos := make([]TorrentInfo, 0)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		ti := TorrentInfo{}
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
			ginCon.JSON(200, TorrentRepo{infos})
		} else {
			ginCon.AbortWithStatus(204)
		}
	})
	c.Visit(url)
}

func Kickass_getMagnet(ginCon *gin.Context) {
	url := ginCon.Query("url")
	c := colly.NewCollector()
	var magnet, torrentfile string
	c.OnHTML("body", func(e *colly.HTMLElement) {
		magnet = e.ChildAttr("a.kaGiantButton", "href")
		torrentfile = getKickass_fileurl(magnet)
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
