package clients

import (
	"fmt"
	"io"

	"github.com/hashicorp/go-retryablehttp"
)

const (
	nameUrl = "https://names.mcquay.me/api/v0/"
	jokeURL = "http://joke.loc8u.com:8888/joke?limitTo=nerdy"
)

type GetResponse struct {
	Response   []byte
	StatusCode int
	Error      error
}

func get(url string) *GetResponse {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5
	resp, err := retryablehttp.Get(url)
	if err != nil {
		return &GetResponse{
			Response:   nil,
			StatusCode: resp.StatusCode,
			Error:      fmt.Errorf("failed to retrieve data from [%s] due to error [%s]", url, err),
		}
	}
	defer resp.Body.Close()
	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return &GetResponse{
			Response:   respData,
			StatusCode: resp.StatusCode,
			Error:      fmt.Errorf("failed to read response body [%+v] from [%s] due to error [%s]", respData, url, err),
		}
	}
	return &GetResponse{
		Response:   respData,
		StatusCode: resp.StatusCode,
		Error:      nil,
	}
}
