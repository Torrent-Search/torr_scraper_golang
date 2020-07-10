package helper

import (
	"encoding/base64"
	"strings"

	"github.com/forgoer/openssl"
)

func GetJioSaavnUrl(enc_url string) string {
	var key []byte = []byte("38346591")
	var src []byte = []byte(enc_url)
	sDec, _ := base64.StdEncoding.DecodeString(string(src))
	dst, _ := openssl.DesECBDecrypt(sDec, key, openssl.PKCS5_PADDING)
	var finUrl string = string(dst)
	finUrl = strings.ReplaceAll(finUrl, "_96.mp4", "_320.mp4")
	return finUrl
}
