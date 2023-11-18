/*
Copyright © 2023 Rebecca [Becci] Govert <becci.govert@gmail.com>
*/
package clients

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// JokeResponse takes the json and formats to go struct for unmarshalling of API data
type JokeResponse struct {
	Response  string    `json:"type"`
	JokeValue JokeValue `json:"value"`
}

type JokeValue struct {
	Categories []string `json:"categories"`
	ID         int      `json:"id"`
	Joke       string   `json:"joke"`
}

type GetJokeResp struct {
	JokeResponse  *JokeResponse
	StatusCode    int
	ErrorResponse error
}

// GetRandomJoke takes in a first and last name and returns a Nerdy Chuck Norris Joke replacing `Chuck Norris` with the name provided
func GetRandomJoke(nr *NameResponse) *GetJokeResp {
	u, err := url.Parse(jokeURL)
	if err != nil {
		return &GetJokeResp{
			JokeResponse:  nil,
			StatusCode:    500,
			ErrorResponse: err,
		}
	}
	q := u.Query()
	q.Add("firstName", nr.FirstName)
	q.Add("lastName", nr.LastName)
	u.RawQuery = q.Encode()

	resp := get(u.String())
	if resp.Error != nil {
		return &GetJokeResp{
			JokeResponse:  nil,
			StatusCode:    resp.StatusCode,
			ErrorResponse: resp.Error,
		}
	}
	var jokeResponse *JokeResponse
	err = json.Unmarshal(resp.Response, &jokeResponse)
	if err != nil {
		return &GetJokeResp{
			JokeResponse:  nil,
			StatusCode:    500,
			ErrorResponse: fmt.Errorf("failed to unmarshal joke [%w]", err),
		}
	}
	return &GetJokeResp{
		JokeResponse:  jokeResponse,
		StatusCode:    200,
		ErrorResponse: nil,
	}
}
