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

func Skytorrents(c *gin.Context) {
	// search := strings.ReplaceAll(strings.TrimSpace(c.Query("search")), " ", "%20")
	param := url.Values{}
	param.Add("query", c.Query("search"))
	url := fmt.Sprint("https://www.skytorrents.lol/?", param.Encode())
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	var client = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	request, _ := http.NewRequest("GET", url, nil)
	res, _ := client.Do(request)
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	selector := doc.Find("tr.result")
	if selector.Length() > 0 {
		infos := make([]TorrentInfo, 0)
		selector.Each(func(i int, s *goquery.Selection) {
			tr := TorrentInfo{}
			tr.Name = s.Find("td:nth-child(1) a:nth-child(1)").Text()
			tr.Seeders = s.Find("td:nth-child(5)").Text()
			tr.Leechers = s.Find("td:nth-child(6)").Text()
			tr.Date = s.Find("td:nth-child(4)").Text()
			tr.Size = s.Find("td:nth-child(2)").Text()

			magnet_selector_with_child := s.Find("td:nth-child(1)").Children()

			if isMagnet(magnet_selector_with_child.Eq(3).AttrOr("href", "")) {
				tr.Magnet = magnet_selector_with_child.Eq(3).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(4).AttrOr("href", "")) {
				tr.Magnet = magnet_selector_with_child.Eq(4).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(5).AttrOr("href", "")) {
				tr.Magnet = magnet_selector_with_child.Eq(5).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(6).AttrOr("href", "")) {
				tr.Magnet = magnet_selector_with_child.Eq(6).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(7).AttrOr("href", "")) {
				tr.Magnet = magnet_selector_with_child.Eq(7).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(8).AttrOr("href", "")) {
				tr.Magnet = magnet_selector_with_child.Eq(8).AttrOr("href", "")
			}
			tr.Url = "https://www.skytorrents.lol/" + s.Find("td:nth-child(1) a:nth-child(1)").AttrOr("href", "")
			tr.Website = "Skytorrents"
			tr.Uploader = "--"
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
