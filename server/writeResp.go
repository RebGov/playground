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
	joke, int, err := getJoke()
	if err != nil {
		return errResponse(err.Error(), int)
	}
	return []byte(joke), nil
}

func errResponse(errMsg string, code int) ([]byte, error) {
	errResponse := &ErrResponse{
		StatusCode: code,
		Status:     "failure",
		ErrorMsg:   errMsg,
	}
	errResp, err := json.Marshal(errResponse)
	if err != nil {
		return nil, err
	}
	return errResp, nil
}

func getJoke() (string, int, error) {
	fName, lName, int, err := clients.GetRandomName()
	if err != nil {
		log.Printf("failed to obtain name [%s]\n", err)
		return "", int, err
	}
	joke, int, err := clients.GetRandomJoke(fName, lName)
	if err != nil {
		log.Printf("failed with error [%s]\n", err)
		return "", int, err
	}
	return joke, 200, nil
}
