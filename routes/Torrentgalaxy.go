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

func Torrentgalaxy(c *gin.Context) {
	search := c.Query("search")
	url := fmt.Sprint("https://torrentgalaxy.to/torrents.php?search=", strings.TrimSpace(search))
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

	selector := doc.Find("div.tgxtablerow")
	if selector.Length() > 0 {
		infos := make([]TorrentInfo, 0)
		selector.Each(func(i int, s *goquery.Selection) {
			tr := TorrentInfo{}
			tr.Name = s.Find("div:nth-child(4)").Text()
			tr.Seeders = s.Find("div:nth-child(11) span font:nth-child(1)").Text()
			tr.Leechers = s.Find("div:nth-child(11) span font:nth-child(2)").Text()
			tr.Date = strings.Split(s.Find("div:nth-child(12)").Text(), " ")[0]
			tr.Size = s.Find("div:nth-child(8)").Text()
			tr.Uploader = s.Find("div:nth-child(7)").Text()
			tr.Magnet = s.Find("#click").Next().Find("a:nth-child(2)").AttrOr("href", "")
			tr.Url = "https://torrentgalaxy.to" + s.Find("div:nth-child(4) a").AttrOr("href", "")
			tr.Website = "Torrent Galaxy"
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
