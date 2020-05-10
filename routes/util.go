package routes

import (
	"log"
	"regexp"
	"strings"
)

func isMagnet(str string) bool {

	// log.Println(str)
	match, err := regexp.MatchString("magnet:\\?xt=urn:[a-z0-9]+:[a-z0-9]{32}", str)

	if err != nil {
		log.Println(err)
	}
	// println(str)
	// println(match)
	return match
}

func replace(str string, new string, old string) string {
	return strings.ReplaceAll(str, new, old)
}
