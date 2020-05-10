package routes

import (
	"fmt"
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
func ByteCountDecimal(b uint64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}
