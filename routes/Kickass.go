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

func Kickass(c *gin.Context) {
	search := strings.ReplaceAll(strings.TrimSpace(c.Query("search")), " ", "%20")
	url := fmt.Sprint("https://kickasstorrents.to/usearch/", search)
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
	// request.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1 RuxitSynthetic/1.0 v1094723656 t4690183951324214268 smf=0")
	res, err := client.Do(request)
	if err != nil {
		log.Print(err)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		print(err)
	}

	selector := doc.Find("tr.odd , tr.even")
	content_avail := doc.Find("span[itemprop=name]").Length()

	if content_avail == 0 {
		infos := make([]TorrentInfo, 0)
		selector.Each(func(i int, s *goquery.Selection) {
			tr := TorrentInfo{}
			tr.Name = strings.TrimPrefix(s.Find(".cellMainLink").Text(), "\n")
			// Seeders
			tr.Seeders = strings.TrimPrefix(s.Find("td:nth-child(5)").Text(), "\n")
			//  Leechers
			tr.Leechers = strings.TrimPrefix(s.Find("td:nth-child(6)").Text(), "\n")
			//  Upload Date
			tr.Date = strings.TrimPrefix(s.Find("td:nth-child(4)").Text(), "\n")
			//  File Size
			tr.Size = strings.TrimPrefix(s.Find("td:nth-child(2)").Text(), "\n")
			// Uploader
			tr.Uploader = strings.TrimPrefix(s.Find("td:nth-child(3)").Text(), "\n")
			// Magnet
			tr.Magnet = ""
			tr.Url = s.Find(".cellMainLink").AttrOr("href", "")
			tr.Website = "Kickass"
			tr.TorrentFileUrl = ""
			infos = append(infos, tr)
		})
		repo := TorrentRepo{infos}
		c.JSON(200, repo)

	} else {
		c.AbortWithStatus(204)
	}
}

func Kickass_getMagnet(c *gin.Context) {
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
	res, err := client.Do(request)
	if err != nil {
		log.Print(err)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Print(err)
	}
	magnet, _ := doc.Find("a.kaGiantButton").Attr("href")

	if strings.HasPrefix(magnet, "magnet") {
		torrentfile := getKickass_fileurl(magnet)
		c.JSON(200, gin.H{"magnet": magnet, "torrentFile": torrentfile})
		defer res.Body.Close()
	} else {
		c.AbortWithStatus(204)
		defer res.Body.Close()
	}
}
