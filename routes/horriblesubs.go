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

func Horriblesubs(c *gin.Context) {
	search := c.Query("search")
	url := fmt.Sprint("https://nyaa.si/user/HorribleSubs?f=0&c=0_0&q=", strings.TrimSpace(search))
	log.Println(url)
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
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalln(err)
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

	selector := doc.Find("tr")
	log.Println(selector.Length())
	if selector.Length() > 0 {
		var infos []TorrentInfo
		selector.Each(func(i int, s *goquery.Selection) {

			if i == 1 {
				return
			}
			var name string
			//  File Name
			if s.Find("td:nth-child(2) a").Length() == 2 {
				name = s.Find("td:nth-child(2) a").Eq(1).Text()
			} else {
				name = s.Find("td:nth-child(2) a").Text()
			}
			// Seeders
			seeders := s.Find("td:nth-child(6)").Text()
			//  Leechers
			leechers := s.Find("td:nth-child(7)").Text()
			//  Upload Date
			upload_date := s.Find("td:nth-child(5)").Text()
			//  File Size
			file_size := s.Find("td:nth-child(4)").Text()

			// Magnet
			magnet, _ := s.Find("td:nth-child(3) a:nth-child(2)").Attr("href")
			url, _ := s.Find("td:nth-child(2) a").Attr("href")
			website := "Horrible Subs"
			infos = append(infos, TorrentInfo{name, "https://nyaa.si" + url, seeders, leechers, upload_date, file_size, website, magnet, website})

		})
		repo := TorrentRepo{infos}
		c.JSON(200, repo)

	} else {
		c.AbortWithStatus(204)
	}
}
