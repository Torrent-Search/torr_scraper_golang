package helper

import (
	"fmt"
	"net/url"
	"regexp"
)

func GenerateTorrentDownloadMagnet(str string) string {
	if str == "" {
		return ""
	}
	tracker := "&dn=&tr=udp%3A%2F%2Ftracker.torrent.eu.org%3A451%2Fannounce&tr=http%3A%2F%2Ftracker.ipv6tracker.ru%3A80%2Fannounce&tr=udp%3A%2F%2Fretracker.hotplug.ru%3A2710%2Fannounce&tr=https%3A%2F%2Ftracker.fastdownload.xyz%3A443%2Fannounce&tr=https%3A%2F%2Fopentracker.xyz%3A443%2Fannounce&tr=http%3A%2F%2Fopen.trackerlist.xyz%3A80%2Fannounce&tr=udp%3A%2F%2Ftracker.birkenwald.de%3A6969%2Fannounce&tr=https%3A%2F%2Ft.quic.ws%3A443%2Fannounce&tr=https%3A%2F%2Ftracker.parrotsec.org%3A443%2Fannounce&tr=udp%3A%2F%2Ftracker.supertracker.net%3A1337%2Fannounce&tr=http%3A%2F%2Fgwp2-v19.rinet.ru%3A80%2Fannounce&tr=udp%3A%2F%2Fbigfoot1942.sektori.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fcarapax.net%3A6969%2Fannounce&tr=udp%3A%2F%2Fretracker.akado-ural.ru%3A80%2Fannounce&tr=udp%3A%2F%2Fretracker.maxnet.ua%3A80%2Fannounce&tr=udp%3A%2F%2Fbt.dy20188.com%3A80%2Fannounce&tr=http%3A%2F%2F0d.kebhana.mx%3A443%2Fannounce&tr=http%3A%2F%2Ftracker.files.fm%3A6969%2Fannounce&tr=http%3A%2F%2Fretracker.joxnet.ru%3A80%2Fannounce&tr=http%3A%2F%2Ftracker.moxing.party%3A6969%2Fannounce"
	return "magnet:?xt=urn:btih:" + fmt.Sprintf("%s%s", str, tracker)
}

func GenerateYtsMagnet(info_hash string) string {
	trackers := [8]string{
		"udp://glotorrents.pw:6969/announce",
		"udp://tracker.opentrackr.org:1337/announce",
		"udp://torrent.gresille.org:80/announce",
		"udp://tracker.openbittorrent.com:80",
		"udp://tracker.coppersurfer.tk:6969",
		"udp://tracker.leechers-paradise.org:6969",
		"udp://p4p.arenabg.ch:1337",
		"udp://tracker.internetwarriors.net:1337",
	}
	params := url.Values{}
	for _, obj := range trackers {
		params.Add("tr", obj)
	}

	return fmt.Sprintf("magnet:?xt=urn:btih:%s&dn=&%s", info_hash, params.Encode())
}
func GetImgUrl(str string) string {
	if str == "" {
		return ""
	}
	// print(str)
	re := regexp.MustCompile(`https(.*?)jpg`)
	// match := re.FindString(str)
	var imgTags = re.FindAllStringSubmatch(str, -1)
	return imgTags[0][0]
}
