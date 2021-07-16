package helper

import (
	"log"
	"os"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
)

func GetProxy() colly.ProxyFunc {
	proxy1 := "socks5://" + os.Getenv("pr1")
	proxy2 := "socks5://" + os.Getenv("pr2")
	rp, err := proxy.RoundRobinProxySwitcher(proxy1, proxy2)
	if err != nil {
		log.Println(err)
	}
	return rp
}
