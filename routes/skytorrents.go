package routes

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func Skytorrents(c *gin.Context) {
	search := c.Query("search")
	url := fmt.Sprint("https://www.skytorrents.lol/?query=", strings.TrimSpace(search))
	log.Println(url)
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
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1 RuxitSynthetic/1.0 v1094723656 t4690183951324214268 smf=0")
	res, err := client.Do(request)
	if err != nil {
		log.Print(err)
	}
	resBody := res.Body
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resBody)
	if err != nil {
		print(err)
	}

	selector := doc.Find("tr.result")
	if selector.Length() > 0 {
		var infos []TorrentInfo
		selector.Each(func(i int, s *goquery.Selection) {

			//  File Name
			name := s.Find("td:nth-child(1) a:nth-child(1)").Text()
			// Seeders
			seeders := s.Find("td:nth-child(5)").Text()
			//  Leechers
			leechers := s.Find("td:nth-child(6)").Text()
			//  Upload Date
			upload_date := s.Find("td:nth-child(4)").Text()
			//  File Size
			file_size := s.Find("td:nth-child(2)").Text()

			// Magnet
			magnet_selector_with_child := s.Find("td:nth-child(1)").Children()

			var magnet string
			if isMagnet(magnet_selector_with_child.Eq(3).AttrOr("href", "")) {
				magnet = magnet_selector_with_child.Eq(3).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(4).AttrOr("href", "")) {
				magnet = magnet_selector_with_child.Eq(4).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(5).AttrOr("href", "")) {
				magnet = magnet_selector_with_child.Eq(5).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(6).AttrOr("href", "")) {
				magnet = magnet_selector_with_child.Eq(6).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(7).AttrOr("href", "")) {
				magnet = magnet_selector_with_child.Eq(7).AttrOr("href", "")
			} else if isMagnet(magnet_selector_with_child.Eq(8).AttrOr("href", "")) {
				magnet = magnet_selector_with_child.Eq(8).AttrOr("href", "")
			}
			url := "https://www.skytorrents.lol/" + s.Find("td:nth-child(1) a:nth-child(1)").AttrOr("href", "")
			website := "Skytorrents"
			infos = append(infos, TorrentInfo{name, url, seeders, leechers, upload_date, file_size, "--", magnet, website})

		})
		repo := TorrentRepo{infos}
		c.JSON(200, repo)

	} else {
		c.AbortWithStatus(204)
	}
}
