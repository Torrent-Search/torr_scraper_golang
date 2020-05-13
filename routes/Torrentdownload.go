package routes

import (
	"fmt"
	"strings"

	"github.com/mmcdole/gofeed"

	"github.com/gin-gonic/gin"
)

func Torrentdownloads(c *gin.Context) {
	search := c.Query("search")
	url := fmt.Sprint("https://www.torrentdownload.info/feed?q=", strings.TrimSpace(search))
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)
	items := feed.Items
	if len(items) > 0 {
		infos := make([]TorrentInfo, 0)

		for _, obj := range items {
			desc := strings.Split(obj.Description, " ")

			tr := TorrentInfo{}
			tr.Name = obj.Title
			tr.Url = obj.Link
			tr.Date = strings.Trim(fmt.Sprint(strings.Split(obj.Published, " ")[0:3]), "[]")
			tr.Size = strings.Trim(fmt.Sprint(desc[1:3]), "[]")
			tr.Seeders = desc[4]
			tr.Leechers = desc[7]
			tr.Uploader = "--"
			tr.Magnet = gn_TorrDwnd_mg(desc[9])
			tr.Website = "Torrent Download"
			infos = append(infos, tr)
		}
		c.JSON(200, infos)
	} else {
		c.AbortWithStatus(204)
	}
}
