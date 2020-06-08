package routes

import (
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Rarbg(c *gin.Context) {

	search := strings.ReplaceAll(strings.TrimSpace(c.Query("search")), " ", "%20")
	api, err := New(search)
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
			strings.Split(obj.PubDate, " ")[0],
			ByteCountDecimal(obj.Size),
			"N/A",
			obj.Download,
			"Rarbg",
			getRarbg_fileurl(obj.Download),
		})
	}
	if len(infos) > 0 {
		c.JSON(200, TorrentRepo{infos})
	} else {
		c.AbortWithStatus(204)
	}
}
