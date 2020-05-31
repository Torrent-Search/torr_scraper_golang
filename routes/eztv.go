package routes

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func Eztv(c *gin.Context) {
	url := fmt.Sprintf("https://eztv.io/search/%s", url.PathEscape(c.Query("search")))
	print(url)
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
	selector := doc.Find("tr.forum_header_border")

	if selector.Length() > 0 {
		infos := make([]TorrentInfo, 0)
		selector.Each(func(i int, s *goquery.Selection) {
			tr := TorrentInfo{}
			tr.Name = strings.Trim(s.Find("td:nth-child(2)").Text(), "\n")
			tr.Seeders = s.Find("td:nth-child(6)").Text()
			tr.Leechers = "N/A"
			tr.Date = s.Find("td:nth-child(5)").Text()
			tr.Size = s.Find("td:nth-child(4)").Text()
			tr.Magnet = s.Find("td:nth-child(3) a:nth-child(1)").AttrOr("href", "")
			tr.Url = "https://eztv.io" + s.Find("td:nth-child(2) a").AttrOr("href", "")
			tr.Website = "EZTV"
			tr.Uploader = "--"
			tr.TorrentFileUrl = s.Find("td:nth-child(3) a:nth-child(2)").AttrOr("href", "")
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
