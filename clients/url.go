package clients

import (
	"fmt"
	"io"
	"net/http"
)

const (
	nameUrl = "https://names.mcquay.me/api/v0/"
	jokeURL = "http://joke.loc8u.com:8888/joke?limitTo=nerdy&firstName=%s&lastName=%s"
)

func get(url string) ([]byte, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("failed to retrieve data from [%s] due to error [%s]", url, err)
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return respData, resp.StatusCode, fmt.Errorf("failed to read response body [%+v] from [%s] due to error [%s]", respData, url, err)
	}
	return respData, resp.StatusCode, nil
}
