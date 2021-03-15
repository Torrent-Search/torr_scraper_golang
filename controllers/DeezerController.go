package controller

import (
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/gofiber/fiber"
	helper "github.com/scraper_v2/helper"
	"github.com/scraper_v2/models/music"
)

func DeezerController(fibCon *fiber.Ctx) {
	var (
		param  url.Values = url.Values{}
		url    string
		deezer music.Deezer
		err    error
	)
	param.Add("query", fibCon.Query("search"))
	url = fmt.Sprintf("https://thearq.tech/deezer?count=25&%s", param.Encode())
	res, err := helper.GetResponse(url)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	deezer, err = music.UnmarshalDeezer(body)

	if err != nil {
		fibCon.Status(204)
		return
	}
	fibCon.Status(200).JSON(deezer)
}
