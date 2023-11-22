package connect

import (
	"net/http"
	"time"
)

var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: 5 * time.Second,
}

func Get(url string) bool {

	rsp, err := client.Get(url)
	if err != nil {
		return false
	}
	rsp.Body.Close()
	return rsp.StatusCode == http.StatusOK
}
