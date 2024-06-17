package util

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"net/http"
)

var httpClient *resty.Client

func GetHttpClient() *resty.Client {
	if httpClient == nil {
		client := resty.New()
		client.SetTransport(&http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		})
		httpClient = client
	}
	return httpClient
}

func HttpGet(path string, result interface{}) error {
	_, err := GetHttpClient().R().SetResult(&result).Get(path)
	return err
}

func HttpPost(path string, body interface{}, response interface{}) error {
	_, err := GetHttpClient().R().SetResult(&response).SetBody(body).Post(path)
	return err
}
