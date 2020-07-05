package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/gofiber/fiber"
	helper "github.com/scraper_v2/helper"
	"github.com/scraper_v2/models/music"
)

func JioSaavnRawController(fibCon *fiber.Ctx) {
	var (
		param         url.Values = url.Values{}
		url           string
		jioSaavnQuery music.JioSaavnQuery
	)
	param.Add("query", fibCon.Query("search"))
	url = fmt.Sprintf("https://www.saavn.com/api.php?__call=autocomplete.get&_marker=0&ctx=android&_format=json&_marker=0&%s", param.Encode())
	res, _ := helper.GetResponse(url)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	jioSaavnQuery, _ = music.UnmarshalJioSaavnQuery(body)
	fibCon.Status(200).JSON(jioSaavnQuery)
}
func JioSaavnSongController(fibCon *fiber.Ctx) {
	var (
		param    url.Values = url.Values{}
		url      string
		pid      string = fibCon.Query("search")
		result   map[string]music.SongsDataWithLink
		songData music.SongsDataWithLink = music.SongsDataWithLink{}
	)

	param.Add("pids", pid)
	url = "https://www.jiosaavn.com/api.php?cc=in&_marker=0%253F_marker%253D0&_format=json&model=Redmi_5A&__call=song.getDetails&" + param.Encode()
	res, _ := helper.GetResponse(url)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &result)
	songData = result[pid]
	songData.URL = helper.GetJioSaavnUrl(result[pid].URL)
	songData.Image = strings.ReplaceAll(songData.Image, "150x150", "500x500")
	fibCon.Status(200).JSON(songData)
}
func JioSaavnAlbumController(fibCon *fiber.Ctx) {
	var (
		param url.Values = url.Values{}
		url   string
		pid   string = fibCon.Query("search")
		album music.AlbumWithSongs
	)

	param.Add("albumid", pid)
	url = fmt.Sprintf("https://www.jiosaavn.com/api.php?_format=json&__call=content.getAlbumDetails&_format=json&%s", param.Encode())
	res, _ := helper.GetResponse(url)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	body = []byte(strings.Split(string(body), "-->")[1])
	album, _ = music.UnmarshalAlbumWithSongs(body)

	var temp string
	var tempList []music.SongsDataWithLink = make([]music.SongsDataWithLink, 0)
	for _, obj := range album.Songs {
		temp = obj.URL
		obj.URL = helper.GetJioSaavnUrl(temp)
		tempList = append(tempList, obj)
	}
	album.Songs = tempList

	fibCon.Status(200).JSON(album)
}
func JioSaavnPlaylistController(fibCon *fiber.Ctx) {
	var (
		param    url.Values = url.Values{}
		url      string
		pid      string = fibCon.Query("search")
		playlist music.Playlist
	)

	param.Add("listid", pid)
	url = fmt.Sprintf("https://www.jiosaavn.com/api.php?_format=json&__call=playlist.getDetails&%s", param.Encode())
	res, _ := helper.GetResponse(url)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	body = []byte(strings.Split(string(body), "-->")[1])
	playlist, _ = music.UnmarshalPlaylist(body)

	var temp string
	var tempList []music.SongsDataWithLink = make([]music.SongsDataWithLink, 0)
	for _, obj := range playlist.Songs {
		temp = obj.URL
		obj.URL = helper.GetJioSaavnUrl(temp)
		tempList = append(tempList, obj)
	}
	playlist.Songs = tempList

	fibCon.Status(200).JSON(playlist)
}

func JioSaavnHomeController(fibCon *fiber.Ctx) {
	var (
		url          string
		charts       music.Charts
		trending     music.Trending
		topPlaylists music.TopPlaylist
	)

	url = "https://www.jiosaavn.com/api.php?__call=webapi.getLaunchData&api_version=4&_format=json&_marker=0&ctx=wap6dot0&app_version=4"
	res, _ := helper.GetResponse(url)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	charts, _ = music.UnmarshalCharts(body)
	trending, _ = music.UnmarshalTrending(body)
	topPlaylists, _ = music.UnmarshalTopPlaylist(body)

	fibCon.Status(200).JSON(fiber.Map{"charts": charts.Items, "trending": trending.TrendingItem, "top_playlists": topPlaylists.TopPlaylistItem})
}
