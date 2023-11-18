/*
Copyright © 2023 Rebecca [Becci] Govert <becci.govert@gmail.com>

Creates response to send to the user
*/
package server

import (
	"encoding/json"
	"log"

	"github.com/RebGov/playground/clients"
)

type ErrResponse struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	ErrorMsg   string `json:"error_message"`
}

func createResponse() ([]byte, error) {
	joke, errR := getJoke()
	if errR != nil {
		errResp, err := json.Marshal(errR)
		if err != nil {
			return nil, err
		}
		return errResp, nil
	}
	return []byte(joke), nil
}

func getJoke() (string, *ErrResponse) {
	nr := clients.GetRandomName()
	errResponse := &ErrResponse{}
	if nr.ErrorResponse != nil {
		log.Printf("failed to obtain name [%s]\n", nr.ErrorResponse)
		errResponse.ErrorMsg = nr.ErrorResponse.Error()
		errResponse.Status = "failure"
		errResponse.StatusCode = nr.StatusCode
		return "", errResponse
	}
	jr := clients.GetRandomJoke(nr.NameResponse)
	if jr.ErrorResponse != nil {
		log.Printf("failed with error [%s]\n", jr.ErrorResponse.Error())
		errResponse.ErrorMsg = jr.ErrorResponse.Error()
		errResponse.Status = "failure"
		errResponse.StatusCode = jr.StatusCode
		return "", errResponse
	}
	return jr.JokeResponse.JokeValue.Joke, nil
}
