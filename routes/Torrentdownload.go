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

func Torrentdownloads(c *gin.Context) {
	search := c.Query("search")
	url := fmt.Sprint("https://www.torrentdownload.info/search?q=", strings.TrimSpace(search))
	log.Println(search)
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
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

	selector := doc.Find("body div table:nth-child(8) tr")
	log.Println(selector.Length())
	if selector.Length() > 1 {
		var infos []TorrentInfo
		selector.Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}
			//  File Name
			name := s.Find("td.tdleft").Find("a").Text()
			//  Seeders
			seeders := s.Find("td:nth-child(4)").Text()
			//  Leechers
			leechers := s.Find("td:nth-child(5)").Text()
			//  Upload Date
			upload_date := s.Find("td:nth-child(2)").Text()
			//  File Size
			file_size := s.Find("td:nth-child(3)").Text()

			//  Uploader Name
			uploader_name := "--"

			//  url
			url =
				"https://www.torrentdownload.info" +
					s.Find("td.tdleft").Find("a").AttrOr("href", "")
			website := "torrentdownloads"

			infos = append(infos, TorrentInfo{name, url, seeders, leechers, upload_date, file_size, uploader_name, "", website})

		})
		repo := TorrentRepo{infos}
		c.JSON(200, repo)

	} else {
		c.AbortWithStatus(204)
	}
}

func Torrrentdownload_getMagnet(c *gin.Context) {
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
		log.Fatalln(err)
	}
	// request.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1 RuxitSynthetic/1.0 v1094723656 t4690183951324214268 smf=0")
	res, err := client.Do(request)
	if err != nil {
		log.Print(err)
	}
	resBody := res.Body
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resBody)
	if err != nil {
		log.Print(err)
	}
	magnet, _ := doc.Find("tbody tr:nth-child(5) td span a").Attr("href")

	// log.Println(magnet)
	if strings.HasPrefix(magnet, "magnet") {
		c.JSON(200, gin.H{"magnet": magnet})
	} else {
		c.AbortWithStatus(204)
	}
}
