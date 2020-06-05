package routes

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func Limetorrents(ginCon *gin.Context) {
	search := strings.ReplaceAll(strings.TrimSpace(ginCon.Query("search")), " ", "%20")
	url := fmt.Sprintf("https://www.limetorrents.info/search/all/%s/", search)
	c := colly.NewCollector()
	infos := make([]TorrentInfo, 0)
	c.OnHTML("body", func(e *colly.HTMLElement) {
		ti := TorrentInfo{}
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
			ti.Url = "https://www.limetorrents.info" + e.ChildAttr("td.tdleft div.tt-name a:nth-child(2)", "href")
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
			ginCon.JSON(200, TorrentRepo{infos})
		} else {
			ginCon.AbortWithStatus(204)
		}
	})
	c.Visit(url)
}

func Limetorrents_getMagnet(ginCon *gin.Context) {
	url := ginCon.Query("url")
	c := colly.NewCollector()
	var magnet, torrentfile string
	c.OnHTML("body", func(e *colly.HTMLElement) {
		magnet = e.ChildAttr("body > div > table:nth-child(6) > tbody > tr:nth-child(5) > td > span > a", "href")
		torrentfile = e.ChildAttr("body > div > table:nth-child(6) > tbody > tr:nth-child(3) > td > span > a", "href")
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
