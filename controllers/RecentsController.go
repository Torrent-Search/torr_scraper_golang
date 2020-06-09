package controller

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber"
	helper "github.com/scraper_v2/helper"
	models "github.com/scraper_v2/models"
)

func RecentMoviesController(fibCon *fiber.Ctx) {
	url := "https://torrentgalaxy.to/"
	c := colly.NewCollector()
	var infos = make([]models.Recents, 0)
	var repo models.RecentRepo = models.RecentRepo{}
	var re models.Recents = models.Recents{}
	c.OnHTML(".panel-body.slidingDivf-6e422c70dd796e04eec79baaea3d169e3f1c5cd1", func(e *colly.HTMLElement) {
		// div:nth-child(3)
		e.ForEach("div:nth-child(3) .panel-body.slidingDivb-b6a23717a851a6fc9b4c2e09f0073f0857d7f4d8 div:nth-child(2) .tgxtable div.tgxtablerow", func(i int, a *colly.HTMLElement) {
			re.Name = a.ChildText("div:nth-child(1) a b")
			re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(2)", "href")
			if len(strings.Split(re.Url, "=")) == 1 {
				re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(3)", "href")
			}
			re.ImgFileUrl = helper.GetImgUrl(a.Attr("onmouseover"))
			re.Imdb_code = strings.Split(re.Url, "=")[1]
			infos = append(infos, re)
		})
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

func RecentShowsController(fibCon *fiber.Ctx) {
	url := "https://torrentgalaxy.to/"
	c := colly.NewCollector()
	var infos = make([]models.Recents, 0)
	var repo models.RecentRepo = models.RecentRepo{}
	var re models.Recents = models.Recents{}
	c.OnHTML(".panel-body.slidingDivf-6e422c70dd796e04eec79baaea3d169e3f1c5cd1", func(e *colly.HTMLElement) {
		// div:nth-child(3)
		e.ForEach("div:nth-child(4) .panel-body.slidingDivb-f4d4d7e21ce39705d6fca31c285a979a77742df9 div:nth-child(2) .tgxtable div.tgxtablerow", func(i int, a *colly.HTMLElement) {
			re.Name = a.ChildText("div:nth-child(1) a b")

			re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(2)", "href")
			if len(strings.Split(re.Url, "=")) == 1 {
				re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(3)", "href")
			}
			re.Imdb_code = strings.Split(re.Url, "=")[1]
			re.ImgFileUrl = helper.GetImgUrl(a.Attr("onmouseover"))
			infos = append(infos, re)
		})
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
