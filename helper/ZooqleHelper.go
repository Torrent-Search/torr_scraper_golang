package helper

import (
	"fmt"
	"strings"
)

func GetZooqleMediaUrl(mediatype string, name string, id string) string {
	switch mediatype {
	case "t":
		mediatype = "tv"
	case "m":
		mediatype = "movie"
	}
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ReplaceAll(name, "'", "-")

	return fmt.Sprintf("https://zooqle.com/%s/%s-%s.html", mediatype, name, id)
}
