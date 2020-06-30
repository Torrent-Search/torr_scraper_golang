package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber"
	helper "github.com/scraper_v2/helper"
	"github.com/scraper_v2/models"
)

func ZooqleController(fibCon *fiber.Ctx) {
	param := url.Values{}
	param.Add("q", fibCon.Query("search"))
	var url string = fmt.Sprintf("https://zooqle.unblockit.id/search?%s", param.Encode())
	var infos = make([]models.TorrentInfo, 0)
	var repo models.TorrentRepo = models.TorrentRepo{}
	var ti models.TorrentInfo = models.TorrentInfo{}
	var c *colly.Collector = colly.NewCollector()
	var seedLeechString string = ""
	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(i int, e *colly.HTMLElement) {
			if i == 0 {
				return
			}
			ti.Name = e.ChildText("td:nth-child(2) a")
			seedLeechString = e.ChildAttr("td:nth-child(6) div", "title")
			ti.Seeders = strings.Split(seedLeechString, " ")[1]
			ti.Leechers = strings.Split(seedLeechString, " ")[4]
			ti.Date = e.ChildText("td:nth-child(5)")
			ti.Size = e.ChildText("td:nth-child(4) div div")
			ti.Magnet = e.ChildAttr("td:nth-child(3) ul li:nth-child(2) a", "href")
			infos = append(infos, ti)
		})
		// }
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnScraped(func(r *colly.Response) {
		if len(infos) > 0 {
			repo.Data = &infos
			fibCon.Status(200).JSON(repo)
		} else {
			fibCon.Status(204)
		}
	})
	c.Visit(url)

}

func ZooqleSeriesController(fibCon *fiber.Ctx) {
	var (
		media_type  string                = fibCon.Query("type")
		name        string                = fibCon.Query("name")
		id          string                = fibCon.Query("id")
		seasonsList []models.ZooqleSeason = make([]models.ZooqleSeason, 0)
		zooqleItem  models.ZooqleItem     = models.ZooqleItem{}
		url         string                = helper.GetZooqleMediaUrl(media_type, name, id)
		c           *colly.Collector      = colly.NewCollector()
	)
	c.OnHTML("div.panel-group", func(e *colly.HTMLElement) {
		zooqleItem.Seasons = e.DOM.Find("div.panel.panel-default.eplist").Length()
		e.ForEach("div.panel.panel-default.eplist", func(i int, e *colly.HTMLElement) {
			var season models.ZooqleSeason = models.ZooqleSeason{}
			season.Season_No = zooqleItem.Seasons - i
			var episodeList []models.ZooqleEpisode = make([]models.ZooqleEpisode, 0)
			e.DOM.Find("li.list-group-item").Each(func(i int, sel *goquery.Selection) {
				var episode models.ZooqleEpisode = models.ZooqleEpisode{}
				episode.Episode_No = sel.Find("span.smaller.text-muted.epnum").Text()
				temp_url := sel.Children().Eq(4).AttrOr("data-href", "")
				if temp_url == "" {
					return
				}
				episode.Data_url = "https://zooqle.unblockit.id" + temp_url
				// episode.EpisodesData = temp(episode.Data_url)
				episodeList = append(episodeList, episode)
			})
			season.SeasonEpisode = episodeList
			seasonsList = append(seasonsList, season)
		})

	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnScraped(func(r *colly.Response) {
		data := models.ZooqleData{seasonsList}
		fibCon.JSON(data)
	})
	c.Visit(url)
}

func ZooqleMediaSearchController(fibCon *fiber.Ctx) {

	var searchQuery string = fibCon.Query("search")
	searchQuery = strings.ReplaceAll(searchQuery, " ", "%20")
	var url string = fmt.Sprintf("https://zooqle.unblockit.id/qss/%s?lc=en", searchQuery)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("Referer", "https://zooqle.com/")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var zooqleSearch []models.ZooqleSearch
	zooqleSearch, _ = models.UnmarshalZooqle(body)
	fibCon.JSON(zooqleSearch)
}

func ZooqleEpisodeController(fibCon *fiber.Ctx) {
	var (
		c               *colly.Collector           = colly.NewCollector()
		episodeDataList []models.ZooqleEpisodeData = make([]models.ZooqleEpisodeData, 0)
		url             string                     = fibCon.Query("url")
	)
	c.OnHTML("table", func(e *colly.HTMLElement) {

		e.ForEach("tr", func(i int, ei *colly.HTMLElement) {
			if ei.DOM.Children().Length() <= 5 {
				return
			}

			var episodeData models.ZooqleEpisodeData = models.ZooqleEpisodeData{}
			episodeData.Name = ei.ChildText("td:nth-child(2) a")
			var seedLeechString string = ei.ChildAttr("td:nth-child(6) div", "title")
			episodeData.Seeders = strings.Split(seedLeechString, " ")[1]
			episodeData.Leechers = strings.Split(seedLeechString, " ")[4]
			episodeData.Date = ei.ChildText("td:nth-child(5)")
			episodeData.Size = ei.ChildText("td:nth-child(4) div div")
			episodeData.Magnet = ei.ChildAttr("td:nth-child(3) ul li:nth-child(2) a", "href")
			episodeDataList = append(episodeDataList, episodeData)

		})

	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnScraped(func(r *colly.Response) {
		fibCon.JSON(episodeDataList)
	})
	c.Visit(url)
	c.Wait()
}

// func temp(url string) []models.ZooqleEpisodeData {
// 	var (
// 		c               *colly.Collector           = colly.NewCollector()
// 		episodeDataList []models.ZooqleEpisodeData = make([]models.ZooqleEpisodeData, 0)
// 	)
// 	c.OnHTML("table", func(e *colly.HTMLElement) {

// 		e.ForEach("tr", func(i int, ei *colly.HTMLElement) {
// 			if ei.DOM.Children().Length() <= 5 {
// 				return
// 			}

// 			var episodeData models.ZooqleEpisodeData = models.ZooqleEpisodeData{}
// 			episodeData.Name = ei.ChildText("td:nth-child(2) a")
// 			var seedLeechString string = ei.ChildAttr("td:nth-child(6) div", "title")
// 			episodeData.Seeders = strings.Split(seedLeechString, " ")[1]
// 			episodeData.Leechers = strings.Split(seedLeechString, " ")[4]
// 			episodeData.Date = ei.ChildText("td:nth-child(5)")
// 			episodeData.Size = ei.ChildText("td:nth-child(4) div div")
// 			episodeData.Magnet = ei.ChildAttr("td:nth-child(3) ul li:nth-child(2) a", "href")
// 			episodeDataList = append(episodeDataList, episodeData)

// 		})

// 	})
// 	c.OnError(func(r *colly.Response, err error) {
// 		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
// 	})
// 	c.OnScraped(func(r *colly.Response) {

// 	})
// 	c.Visit(url)
// 	c.Wait()
// 	return episodeDataList
// }
