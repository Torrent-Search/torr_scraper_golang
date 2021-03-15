package controller

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os/exec"

	"github.com/gofiber/fiber"
	helper "github.com/scraper_v2/helper"
	"github.com/scraper_v2/models/music"
)

func YoutubeController(fibCon *fiber.Ctx) {
	var (
		param url.Values = url.Values{}
		url   string
		yt    music.Yt
		err   error
	)
	param.Add("query", fibCon.Query("search"))
	url = fmt.Sprintf("https://thearq.tech/youtube?count=15&%s", param.Encode())
	res, err := helper.GetResponse(url)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	yt, err = music.UnmarshalYt(body)

	if err != nil {
		fibCon.Status(204)
		return
	}
	fibCon.Status(200).JSON(yt)
}

func YtAudioUrl(fibCon *fiber.Ctx) {
	var (
		url       string
		query     string
		music_url string
	)

	query = fibCon.Query("search")
	url = fmt.Sprintf("https://youtube.com%s", query)

	cmd := exec.Command("youtube-dl", "-f", "bestaudio", "-g", url)

	stdout, err := cmd.Output()

	if err != nil {

		fibCon.Status(204)
		cmd.Process.Kill()
		return
	}

	music_url = string(stdout)
	cmd.Process.Kill()
	fibCon.Status(200).SendString(music_url)
}
