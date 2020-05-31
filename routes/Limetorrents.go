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

func Limetorrents(c *gin.Context) {
	search := strings.ReplaceAll(strings.TrimSpace(c.Query("search")), " ", "%20")
	url := fmt.Sprintf("https://www.limetorrents.info/search/all/%s/", search)

	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
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

	selector := doc.Find("table.table2 tbody tr")
	infos := make([]TorrentInfo, 0)
	if selector.Length() > 1 {
		selector.Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}
			tr := TorrentInfo{}
			tr.Name = s.Find("td:nth-child(1)").Text()
			tr.Seeders = s.Find("td:nth-child(4)").Text()
			tr.Leechers = s.Find("td:nth-child(5)").Text()
			tr.Date = strings.Split(s.Find("td:nth-child(2)").Text(), " - ")[0]
			tr.Size = s.Find("td:nth-child(3)").Text()
			tr.Uploader = "--"
			tr.Magnet = gn_Lime_mg(s.Find("a.csprite_dl14").AttrOr("href", ""))
			tr.Url = "https://www.limetorrents.info" + s.Find("td.tdleft div.tt-name a:nth-child(2)").AttrOr("href", "")
			tr.Website = "Limetorrents"
			tr.TorrentFileUrl = ""
			infos = append(infos, tr)

		})
		repo := TorrentRepo{infos}
		c.JSON(200, repo)
		defer res.Body.Close()

	} else {
		c.AbortWithStatus(204)
		defer res.Body.Close()
	}
}

func Limetorrents_getMagnet(c *gin.Context) {
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
	magnet, _ := doc.Find("#content > div:nth-child(6) > div:nth-child(1) > div > div:nth-child(13) > div > p > a").Attr("href")
	torrentfile := doc.Find("#content > div:nth-child(6) > div:nth-child(1) > div > div:nth-child(7) > div > p > a").AttrOr("href", "")
	if strings.HasPrefix(magnet, "magnet") {
		c.JSON(200, gin.H{"magnet": magnet, "torrentFile": torrentfile})
		defer res.Body.Close()
	} else {
		c.AbortWithStatus(204)
		defer res.Body.Close()
	}
}
