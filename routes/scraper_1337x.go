package routes

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func Torr_1337x(c *gin.Context) {
	search := strings.ReplaceAll(strings.TrimSpace(c.Query("search")), " ", "%20")
	url := fmt.Sprintf("https://1337x.to/search/%s/1/", search)
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var client = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	request, _ := http.NewRequest("GET", url, nil)

	res, _ := client.Do(request)
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	selector := doc.Find("tr")
	if selector.Length() > 0 {
		infos := make([]TorrentInfo, 0)
		selector.Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}
			tr := TorrentInfo{}
			tr.Name = s.Find("td.coll-1.name").Text()
			tr.Seeders = s.Find("td.coll-2.seeds").Text()
			tr.Leechers = s.Find("td.coll-3.leeches").Text()
			tr.Date = s.Find("td.coll-date").Text()
			tr.Size = s.Find("td:nth-child(5)").Clone().Children().Remove().End().Text()
			tr.Uploader = s.Find("td:nth-child(6)").Text()
			tr.Url =
				"https://1337x.to" +
					s.Find("td.coll-1.name > a:nth-child(2)").AttrOr("href", "")
			tr.Website = "1337x"
			tr.TorrentFileUrl = ""
			infos = append(infos, tr)

		})
		defer res.Body.Close()
		repo := TorrentRepo{infos}
		c.JSON(200, repo)

	} else {
		defer res.Body.Close()
		c.AbortWithStatus(204)
	}
}

func Torr_1337x_getMagnet(c *gin.Context) {
	search_url := c.Query("url")
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var client = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	request, _ := http.NewRequest("GET", search_url, nil)
	res, _ := client.Do(request)
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	magnet, _ := doc.Find("div.clearfix ul li a").Attr("href")
	torrentfile := doc.Find("div.clearfix ul li.dropdown ul li:nth-child(1) a").AttrOr("href", "")
	if strings.HasPrefix(magnet, "magnet") {
		defer res.Body.Close()
		c.JSON(200, gin.H{"magnet": magnet, "torrentFile": torrentfile})
	} else {
		defer res.Body.Close()
		c.AbortWithStatus(204)
	}
}
