package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/scraper/routes"
)

func main() {

	listenPort := fmt.Sprintf(":%s", os.Getenv("PORT"))
	router := gin.Default()
	api := router.Group("api")
	{
		api.GET("/", routes.Ping)
		api.GET("/skytorrents", routes.Skytorrents)
		api.GET("/1337x", routes.Torr_1337x)
		api.GET("/1337x_mg", routes.Torr_1337x_getMagnet)
		api.GET("/horriblesubs", routes.Horriblesubs)
		api.GET("/nyaa", routes.Nyaa)
		api.GET("/kickass", routes.Kickass)
		api.GET("/kickass_mg", routes.Kickass_getMagnet)
		api.GET("/limetorrents", routes.Limetorrents)
		api.GET("/limetorrents_mg", routes.Limetorrents_getMagnet)
		api.GET("/piratebay", routes.PirateBay)
		api.GET("/torrentdownloads", routes.Torrentdownloads)
		api.GET("/torrentdownloads_mg", routes.Torrrentdownload_getMagnet)
		api.GET("/tgx", routes.Torrentgalaxy)
		// api.GET("/torrentdownloads_mg", routes.Torrrentdownload_getMagnet)

	}
	router.Run(listenPort)
}
