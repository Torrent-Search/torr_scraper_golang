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

func Torr_1337x(c *gin.Context) {
	search := c.Query("search")
	url := fmt.Sprintf("https://1337x.to/search/%s/1/", strings.TrimSpace(search))
	log.Println(search)
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

	selector := doc.Find("tr")
	log.Println(selector.Length())
	if selector.Length() > 0 {
		var infos []TorrentInfo
		selector.Each(func(i int, s *goquery.Selection) {
			if i == 0 {
				return
			}
			//  File Name
			name := s.Find("td.coll-1.name").Text()
			//  Seeders
			seeders := s.Find("td.coll-2.seeds").Text()
			//  Leechers
			leechers := s.Find("td.coll-3.leeches").Text()
			//  Upload Date
			upload_date := s.Find("td.coll-date").Text()
			//  File Size
			file_size := s.Find("td:nth-child(5)").Clone().Children().Remove().End().Text()

			//  Uploader Name
			uploader_name := s.Find("td:nth-child(6)").Text()

			//  url
			url =
				"https://1337x.to" +
					s.Find("td.coll-1.name > a:nth-child(2)").AttrOr("href", "")
			website := "1337x"

			infos = append(infos, TorrentInfo{name, url, seeders, leechers, upload_date, file_size, uploader_name, "", website})

		})
		repo := TorrentRepo{infos}
		c.JSON(200, repo)

	} else {
		c.AbortWithStatus(204)
	}
}

func Torr_1337x_getMagnet(c *gin.Context) {
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
	magnet, _ := doc.Find("div.clearfix ul li a").Attr("href")

	// log.Println(magnet)
	if strings.HasPrefix(magnet, "magnet") {
		c.JSON(200, gin.H{"magnet": magnet})
	} else {
		c.AbortWithStatus(204)
	}
}
