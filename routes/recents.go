package routes

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func GetTrendingMovies(ginCon *gin.Context) {
	url := "https://torrentgalaxy.to/"
	c := colly.NewCollector()
	infos := make([]Recents, 0)
	c.OnHTML(".panel-body.slidingDivf-6e422c70dd796e04eec79baaea3d169e3f1c5cd1", func(e *colly.HTMLElement) {
		re := Recents{}
		// div:nth-child(3)
		e.ForEach("div:nth-child(3) .panel-body.slidingDivb-b6a23717a851a6fc9b4c2e09f0073f0857d7f4d8 div:nth-child(2) .tgxtable div.tgxtablerow", func(i int, a *colly.HTMLElement) {
			re.Name = a.ChildText("div:nth-child(1) a b")
			re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(2)", "href")
			if len(strings.Split(re.Url, "=")) == 1 {
				re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(3)", "href")
			}
			re.ImgFileUrl = GetImgUrl(a.Attr("onmouseover"))
			re.Imdb_code = strings.Split(re.Url, "=")[1]
			infos = append(infos, re)
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnScraped(func(r *colly.Response) {
		if len(infos) > 0 {
			ginCon.JSON(200, RecentsRepo{infos})
		} else {
			ginCon.AbortWithStatus(204)
		}
	})
	c.Visit(url)
}

func GetTrendingShows(ginCon *gin.Context) {
	url := "https://torrentgalaxy.to/"
	c := colly.NewCollector()
	infos := make([]Recents, 0)
	c.OnHTML(".panel-body.slidingDivf-6e422c70dd796e04eec79baaea3d169e3f1c5cd1", func(e *colly.HTMLElement) {
		re := Recents{}
		// div:nth-child(3)
		e.ForEach("div:nth-child(4) .panel-body.slidingDivb-f4d4d7e21ce39705d6fca31c285a979a77742df9 div:nth-child(2) .tgxtable div.tgxtablerow", func(i int, a *colly.HTMLElement) {
			re.Name = a.ChildText("div:nth-child(1) a b")

			re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(2)", "href")
			if len(strings.Split(re.Url, "=")) == 1 {
				re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(3)", "href")
			}
			re.Imdb_code = strings.Split(re.Url, "=")[1]
			re.ImgFileUrl = GetImgUrl(a.Attr("onmouseover"))
			infos = append(infos, re)
		})
	})
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.OnScraped(func(r *colly.Response) {
		if len(infos) > 0 {
			ginCon.JSON(200, RecentsRepo{infos})
		} else {
			ginCon.AbortWithStatus(204)
		}
	})
	c.Visit(url)
}
