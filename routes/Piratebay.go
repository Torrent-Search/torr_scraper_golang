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

func PirateBay(c *gin.Context) {
	search := strings.ReplaceAll(strings.TrimSpace(c.Query("search")), " ", "%20")
	url := fmt.Sprintf("https://piratebaylive.com/search?q=%s&cat%5B%5D=&search=Pirate+Search", search)
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 20 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	var client = &http.Client{
		Timeout:   time.Second * 20,
		Transport: netTransport,
	}
	request, _ := http.NewRequest("GET", url, nil)
	res, _ := client.Do(request)
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	selector := doc.Find("#st")
	if selector.Length() > 0 {
		infos := make([]TorrentInfo, 0)
		selector.Each(func(i int, s *goquery.Selection) {
			if i == 0 || i == selector.Length()-1 {
				return
			}
			tr := TorrentInfo{}
			tr.Name = s.Find("span.list-item.item-name.item-title").Text()
			tr.Seeders = s.Find("span.list-item.item-seed").Text()
			tr.Leechers = s.Find("span.list-item.item-leech").Text()
			tr.Date = s.Find("span.list-item.item-uploaded").Text()
			tr.Size = s.Find("span.list-item.item-size").Text()
			tr.Uploader = s.Find("span.list-item.item-user").Text()
			tr.Magnet = s.Find("span.item-icons a").AttrOr("href", "")
			tr.Url = s.Find("span.list-item.item-name.item-title a").AttrOr("href", "")
			tr.Website = "Pirate Bay"
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
