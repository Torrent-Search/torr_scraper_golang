package routes

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func Horriblesubs(c *gin.Context) {
	param := url.Values{}
	param.Add("q", c.Query("search"))
	url := fmt.Sprintf("https://nyaa.si/user/HorribleSubs?f=0&c=0_0&%s", param.Encode())
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
	selector := doc.Find("tr")

	if selector.Length() > 0 {
		infos := make([]TorrentInfo, 0)
		selector.Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}
			tr := TorrentInfo{}
			if s.Find("td:nth-child(2) a").Length() == 2 {
				tr.Name = s.Find("td:nth-child(2) a").Eq(1).Text()
			} else {
				tr.Name = s.Find("td:nth-child(2) a").Text()
			}
			tr.Uploader = "Horrible Subs"
			tr.Seeders = s.Find("td:nth-child(6)").Text()
			tr.Leechers = s.Find("td:nth-child(7)").Text()
			tr.Date = s.Find("td:nth-child(5)").Text()
			tr.Size = s.Find("td:nth-child(4)").Text()
			tr.Magnet = s.Find("td:nth-child(3) a:nth-child(2)").AttrOr("href", "")
			tr.Url = "https://nyaa.si" + s.Find("td:nth-child(2) a").AttrOr("href", "")
			tr.Website = "Horrible Subs"
			tr.TorrentFileUrl = "https://nyaa.si" + s.Find("td:nth-child(3) a:nth-child(1)").AttrOr("href", "")
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
