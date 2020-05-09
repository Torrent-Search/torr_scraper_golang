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

func PirateBay(c *gin.Context) {
	search := c.Query("search")
	url := fmt.Sprint("https://thepiratebay.asia/s/?q=", strings.TrimSpace(search))
	log.Println(url)
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
	if selector.Length() > 1 {
		var infos []TorrentInfo
		selector.Each(func(i int, s *goquery.Selection) {
			if i == 1 {
				return
			}
			name := s.Find("td:nth-child(2) div a").Text()
			// Seeders
			seeders := s.Find("td:nth-child(3)").Text()
			//  Leechers
			leechers := s.Find("td:nth-child(4)").Text()
			//  Upload Date
			file_info := s.Find("font.detDesc").Text()
			// upload_date_temp := strings.Split(file_info, ",")[0]
			//  Upload Date
			upload_date := "" //strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(upload_date_temp, " ", "-"), "Uploaded ", ""), " ", "-")
			//  File Size
			file_size := file_info //string(len(strings.Split(file_info, ", "))) //strings.ReplaceAll(strings.Split(file_info, ", ")[1], "size ", "")
			// Uploader
			uploader := "" //strings.ReplaceAll(strings.Split(file_info, ",")[2], "ULed by ", "")
			// Magnet
			magnet, _ := s.Find("td:nth-child(2) a:nth-child(2)").Attr("href")
			url, _ := s.Find("td:nth-child(2) div a").Attr("href")
			website := "Pirate Bay"
			infos = append(infos, TorrentInfo{name, url, seeders, leechers, upload_date, file_size, uploader, magnet, website})

		})
		repo := TorrentRepo{infos[1:len(infos)]}
		c.JSON(200, repo)

	} else {
		c.AbortWithStatus(204)
	}
}
