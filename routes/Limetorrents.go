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

func Limetorrents(c *gin.Context) {
	search := c.Query("search")
	url := fmt.Sprintf("https://www.limetorrents.info/search/all/%s/", strings.TrimSpace(search))
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

	selector := doc.Find("table.table2 tbody tr")
	if selector.Length() > 1 {
		var infos []TorrentInfo
		selector.Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}
			name := s.Find("td:nth-child(1)").Text()
			// Seeders
			seeders := s.Find("td:nth-child(4)").Text()
			//  Leechers
			leechers := s.Find("td:nth-child(5)").Text()
			//  Upload Date
			upload_date := strings.Split(s.Find("td:nth-child(2)").Text(), " - ")[0]
			//  File Size
			file_size := s.Find("td:nth-child(3)").Text()
			// Uploader
			uploader := "--"
			// Magnet
			magnet := ""
			url, _ := s.Find("td.tdleft div.tt-name a:nth-child(2)").Attr("href")
			website := "Limetorrents"
			infos = append(infos, TorrentInfo{name, "https://www.limetorrents.info" + url, seeders, leechers, upload_date, file_size, uploader, magnet, website})

		})
		repo := TorrentRepo{infos}
		c.JSON(200, repo)

	} else {
		c.AbortWithStatus(204)
	}
}

func Limetorrents_getMagnet(c *gin.Context) {
	search_url := c.Query("url")
	log.Println(search_url)
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
	request, err := http.NewRequest("GET", search_url, nil)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(204)

	}
	// request.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1 RuxitSynthetic/1.0 v1094723656 t4690183951324214268 smf=0")
	res, err := client.Do(request)
	if err != nil {
		log.Print(err)
		c.AbortWithStatus(204)

	}
	resBody := res.Body
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resBody)
	if err != nil {
		log.Print(err)
		c.AbortWithStatus(204)

	}
	magnet, _ := doc.Find("#content > div:nth-child(6) > div:nth-child(1) > div > div:nth-child(13) > div > p > a").Attr("href")

	// log.Println(magnet)
	if strings.HasPrefix(magnet, "magnet") {
		c.JSON(200, gin.H{"magnet": magnet})
	} else {
		c.AbortWithStatus(204)
	}
}
