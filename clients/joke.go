/*
Copyright © 2023 Rebecca [Becci] Govert <becci.govert@gmail.com>
*/
package clients

import (
	"encoding/json"
	"fmt"
	"log"
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

// GetRandomJoke takes in a first and last name and returns a Nerdy Chuck Norris Joke replacing `Chuck Norris` with the name provided
func GetRandomJoke(fName, lName string) (string, int, error) {
	url := fmt.Sprintf(jokeURL, fName, lName)
	resp, int, err := get(fmt.Sprintf(url))
	if err != nil {
		return "", int, err
	}
	var jokeResponse JokeResponse
	err = json.Unmarshal(resp, &jokeResponse)
	if err != nil {
		log.Printf("retying: failed to unmarshal joke response [%s] from [%s] for [%s %s] due to error [%s]", string(resp), url, fName, lName, err)
		// retry on unmarshall due to API not matching response structure as expected
		return GetRandomJoke(fName, lName)
	}
	return jokeResponse.JokeValue.Joke, int, nil
}
