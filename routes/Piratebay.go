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
	search := c.Query("search")
	url := fmt.Sprint("https://thepiratebay.asia/s/?q=", strings.TrimSpace(search))
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
	selector := doc.Find("tr")
	if selector.Length() > 0 {
		infos := make([]TorrentInfo, 0)
		selector.Each(func(i int, s *goquery.Selection) {
			if i == 0 || i == selector.Length()-1 {
				return
			}
			tr := TorrentInfo{}
			tr.Name = s.Find("td:nth-child(2) div a").Text()
			tr.Seeders = s.Find("td:nth-child(3)").Text()
			tr.Leechers = s.Find("td:nth-child(4)").Text()
			file_info := s.Find("td:nth-child(2) font.detDesc").Text()
			upload_date_temp := strings.Split(file_info, ",")
			tr.Date = replace(replace(upload_date_temp[0], "Uploaded ", ""), " ", "-")
			tr.Size = replace(upload_date_temp[1], " Size ", "")
			tr.Uploader = replace(upload_date_temp[2], " ULed by ", "")
			tr.Magnet = s.Find("td:nth-child(2) a:nth-child(2)").AttrOr("href", "")
			tr.Url = s.Find("td:nth-child(2) div a").AttrOr("href", "")
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
