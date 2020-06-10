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
	var listType string = fibCon.Query("list")
	var imdbCodes []string = make([]string, 0)
	c.OnHTML(".panel-body.slidingDivf-6e422c70dd796e04eec79baaea3d169e3f1c5cd1", func(e *colly.HTMLElement) {
		// div:nth-child(3)
		e.ForEach("div:nth-child(3) .panel-body.slidingDivb-b6a23717a851a6fc9b4c2e09f0073f0857d7f4d8 div:nth-child(2) .tgxtable div.tgxtablerow", func(i int, a *colly.HTMLElement) {
			re.Name = a.ChildText("div:nth-child(1) a b")
			re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(2)", "href")
			if len(strings.Split(re.Url, "=")) == 1 {
				re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(3)", "href")
				if len(strings.Split(re.Url, "=")) == 1 {
					return
				}
			}
			re.ImgFileUrl = helper.GetImgUrl(a.Attr("onmouseover"))
			re.Imdb_code = strings.Split(re.Url, "=")[1]
			if contains(&imdbCodes, re.Imdb_code) {
				return
			}
			imdbCodes = append(imdbCodes, re.Imdb_code)
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

	if listType == "short" || listType == "" {
		c.Visit(url)
	} else {
		allRecentMoview(fibCon)
	}
}

func RecentShowsController(fibCon *fiber.Ctx) {
	url := "https://torrentgalaxy.to/"
	c := colly.NewCollector()
	var infos = make([]models.Recents, 0)
	var repo models.RecentRepo = models.RecentRepo{}
	var re models.Recents = models.Recents{}
	var listType string = fibCon.Query("list")
	var imdbCodes []string = make([]string, 0)
	c.OnHTML(".panel-body.slidingDivf-6e422c70dd796e04eec79baaea3d169e3f1c5cd1", func(e *colly.HTMLElement) {
		// div:nth-child(3)
		e.ForEach("div:nth-child(4) .panel-body.slidingDivb-f4d4d7e21ce39705d6fca31c285a979a77742df9 div:nth-child(2) .tgxtable div.tgxtablerow", func(i int, a *colly.HTMLElement) {
			re.Name = a.ChildText("div:nth-child(1) a b")

			re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(2)", "href")
			if len(strings.Split(re.Url, "=")) == 1 {
				re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(3)", "href")
				if len(strings.Split(re.Url, "=")) == 1 {
					return
				}
			}
			re.Imdb_code = strings.Split(re.Url, "=")[1]
			re.ImgFileUrl = helper.GetImgUrl(a.Attr("onmouseover"))
			if contains(&imdbCodes, re.Imdb_code) {
				return
			}
			imdbCodes = append(imdbCodes, re.Imdb_code)
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
	if listType == "short" || listType == "" {
		c.Visit(url)
	} else {
		allRecentShows(fibCon)
	}
}

func allRecentMoview(fibCon *fiber.Ctx) {
	url := "https://torrentgalaxy.to/latest"
	c := colly.NewCollector()
	var infos = make([]models.Recents, 0)
	var repo models.RecentRepo = models.RecentRepo{}
	var re models.Recents = models.Recents{}
	var imdbCodes []string = make([]string, 0)
	c.OnHTML(".panel-body.slidingDivf-327914a1b3f6ce6845672274190f50c58c135c38", func(e *colly.HTMLElement) {

		e.ForEach(".panel-body.slidingDivb-58aff82def33b460cc7a5a1e8d41390b3b6d8ae7 .container-fluid .tgxtable div.tgxtablerow", func(i int, a *colly.HTMLElement) {
			re.Name = a.ChildText("div:nth-child(1) a b")

			re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(2)", "href")
			if len(strings.Split(re.Url, "=")) == 1 {
				re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(3)", "href")
				if len(strings.Split(re.Url, "=")) == 1 {
					return
				}
			}
			if len(strings.Split(re.Url, "=")) == 2 {
				re.Imdb_code = strings.Split(re.Url, "=")[1]
			} else {
				re.Imdb_code = ""
			}
			re.ImgFileUrl = helper.GetImgUrl(a.Attr("onmouseover"))
			if contains(&imdbCodes, re.Imdb_code) {
				return
			}
			imdbCodes = append(imdbCodes, re.Imdb_code)
			infos = append(infos, re)
		})
		e.ForEach(".panel-body.slidingDivb-46b37ed219a36b7a1e15853c1b1a61dcdd3228de .container-fluid .tgxtable div.tgxtablerow", func(i int, a *colly.HTMLElement) {
			re.Name = a.ChildText("div:nth-child(1) a b")

			re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(2)", "href")
			if len(strings.Split(re.Url, "=")) == 1 {
				re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(3)", "href")
				if len(strings.Split(re.Url, "=")) == 1 {
					return
				}
			}
			if len(strings.Split(re.Url, "=")) == 2 {
				re.Imdb_code = strings.Split(re.Url, "=")[1]
			} else {
				re.Imdb_code = ""
			}
			re.ImgFileUrl = helper.GetImgUrl(a.Attr("onmouseover"))
			if contains(&imdbCodes, re.Imdb_code) {
				return
			}
			imdbCodes = append(imdbCodes, re.Imdb_code)
			infos = append(infos, re)
		})
		e.ForEach(".panel-body.slidingDivb-4812a02e04964247810f869934ee89b5a4592dfe .container-fluid .tgxtable div.tgxtablerow", func(i int, a *colly.HTMLElement) {
			re.Name = a.ChildText("div:nth-child(1) a b")

			re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(2)", "href")
			if len(strings.Split(re.Url, "=")) == 1 {
				re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(3)", "href")
				if len(strings.Split(re.Url, "=")) == 1 {
					return
				}
			}
			if len(strings.Split(re.Url, "=")) == 2 {
				re.Imdb_code = strings.Split(re.Url, "=")[1]
			} else {
				re.Imdb_code = ""
			}
			re.ImgFileUrl = helper.GetImgUrl(a.Attr("onmouseover"))
			imdbCodes = append(imdbCodes, re.Imdb_code)

			if contains(&imdbCodes, re.Imdb_code) {
				return
			}
			imdbCodes = append(imdbCodes, re.Imdb_code)
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

func allRecentShows(fibCon *fiber.Ctx) {
	url := "https://torrentgalaxy.to/latest"
	c := colly.NewCollector()
	var infos = make([]models.Recents, 0)
	var repo models.RecentRepo = models.RecentRepo{}
	var re models.Recents = models.Recents{}
	var imdbCodes []string = make([]string, 0)
	c.OnHTML(".panel-body.slidingDivf-327914a1b3f6ce6845672274190f50c58c135c38", func(e *colly.HTMLElement) {

		e.ForEach(".panel-body.slidingDivb-27a8121fbd0b00182db87bc7d301e8bbb71e0367 .container-fluid .tgxtable div.tgxtablerow", func(i int, a *colly.HTMLElement) {
			re.Name = a.ChildText("div:nth-child(1) a b")

			re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(2)", "href")
			if len(strings.Split(re.Url, "=")) == 1 {
				re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(3)", "href")
				if len(strings.Split(re.Url, "=")) == 1 {
					return
				}
			}
			if len(strings.Split(re.Url, "=")) == 2 {
				re.Imdb_code = strings.Split(re.Url, "=")[1]
			} else {
				re.Imdb_code = ""
			}
			re.ImgFileUrl = helper.GetImgUrl(a.Attr("onmouseover"))
			if contains(&imdbCodes, re.Imdb_code) {
				return
			}
			imdbCodes = append(imdbCodes, re.Imdb_code)
			infos = append(infos, re)
		})
		e.ForEach(".panel-body.slidingDivb-f45052460401de73c44a8c8dbaa9ac3e31771172 .container-fluid .tgxtable div.tgxtablerow", func(i int, a *colly.HTMLElement) {
			re.Name = a.ChildText("div:nth-child(1) a b")

			re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(2)", "href")
			if len(strings.Split(re.Url, "=")) == 1 {
				re.Url = "https://torrentgalaxy.to" + a.ChildAttr("#click div a:nth-child(3)", "href")
				if len(strings.Split(re.Url, "=")) == 1 {
					return
				}
			}
			if len(strings.Split(re.Url, "=")) == 2 {
				re.Imdb_code = strings.Split(re.Url, "=")[1]
			} else {
				re.Imdb_code = ""
			}
			re.ImgFileUrl = helper.GetImgUrl(a.Attr("onmouseover"))
			if contains(&imdbCodes, re.Imdb_code) {
				return
			}
			imdbCodes = append(imdbCodes, re.Imdb_code)
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

func contains(s *[]string, e string) bool {
	for _, a := range *s {
		if a == e {
			return true
		}
	}
	return false
}
