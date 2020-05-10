package routes

import (
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Rarbg(c *gin.Context) {
	search := strings.TrimSpace(c.Query("search"))
	api, err := New("cli")
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(500)
	}
	api.SearchString(search)
	api.Format("json_extended")
	api.Limit(20)
	results, err := api.Search()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(204)
		return
	}
	var infos []TorrentInfo
	for _, obj := range results {
		infos = append(infos, TorrentInfo{
			obj.Title,
			"",
			strconv.Itoa(obj.Seeders),
			strconv.Itoa(obj.Leechers),
			obj.PubDate,
			ByteCountDecimal(obj.Size),
			"--",
			obj.Download,
			"Rarbg",
		})
	}
	if len(infos) > 0 {
		c.JSON(200, TorrentRepo{infos})
	} else {
		c.AbortWithStatus(204)
	}
}
