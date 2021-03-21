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
		url      string
		query    string
		title    string
		filename string
	)

	query = fibCon.Query("search")
	title = fibCon.Query("title")
	url = fmt.Sprintf("https://youtube.com%s", query)

	// .%(ext)s gives unknown verb error when used with fmt.Sprintf
	filename = fmt.Sprintf("downloads/%s", title) + ".%(ext)s"
	cmd := exec.Command("youtube-dl", "-f", "bestaudio", "--extract-audio", "--audio-format", "mp3", url, "-o", filename)

	_, err := cmd.Output()

	if err != nil {

		fibCon.Status(204)
		cmd.Process.Kill()
		return
	}

	cmd.Process.Kill()
	fibCon.Status(200).SendString(fmt.Sprintf("downloads/%s.mp3", title))
}
