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

func Torrentgalaxy(c *gin.Context) {
	search := c.Query("search")
	url := fmt.Sprint("https://torrentgalaxy.to/torrents.php?search=", strings.TrimSpace(search))
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

	selector := doc.Find("div.tgxtablerow")
	log.Println(selector.Length())
	if selector.Length() > 0 {
		var infos []TorrentInfo
		selector.Each(func(i int, s *goquery.Selection) {
			// if i == 1 {
			// 	return
			// }
			name := s.Find("div:nth-child(4)").Text()
			// Seeders
			seeders := s.Find("div:nth-child(11) span font:nth-child(1)").Text()
			//  Leechers
			leechers := s.Find("div:nth-child(11) span font:nth-child(2)").Text()
			//  Upload Date
			upload_date := strings.Split(s.Find("div:nth-child(12)").Text(), " ")[0]
			//  File Size
			file_size := s.Find("div:nth-child(8)").Text()

			// Uploader
			uploader := s.Find("div:nth-child(7)").Text()
			// Magnet
			// var magnet string
			// if s.Children().Length() == 12 {
			log.Println(s.Find("div.tgxtablecell#rounded").Length())
			magnet := s.Find("#click").Next().Find("a:nth-child(2)").AttrOr("href", "")
			// } else {
			// magnet = s.Find("div:nth-child(5) a:nth-child(2)").AttrOr("href", "")
			// }
			url, _ := s.Find("div:nth-child(4) a").Attr("href")
			website := "Torrent Galaxy"
			infos = append(infos, TorrentInfo{name, "https://torrentgalaxy.to" + url, seeders, leechers, upload_date, file_size, uploader, magnet, website})

		})
		repo := TorrentRepo{infos}
		c.JSON(200, repo)

	} else {
		c.AbortWithStatus(204)
	}
}
