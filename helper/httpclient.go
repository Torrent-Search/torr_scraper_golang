package helper

import (
	"net"
	"net/http"
	"time"
)

func GetResponse(url string) (*http.Response, error) {
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var client = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	request, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(request)
	return res, err
}
