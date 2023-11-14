/*
Copyright © 2023 Rebecca [Becci] Govert <becci.govert@gmail.com>
*/
package clients

import (
	"encoding/json"
	"log"
)

type NameResponse struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// GetRandomName retrieves a random name from APIa
func GetRandomName() (string, string, int, error) {
	resp, int, err := get(nameUrl)
	var nameResp NameResponse
	err = json.Unmarshal(resp, &nameResp)
	if err != nil {
		log.Printf("retrying: failed to unmarshal random name response [%s] from [%s] due to error [%s]", string(resp), nameUrl, err)
		// retry on unmarshall due to API not matching response structure as expected
		return GetRandomName()
	}
	return nameResp.FirstName, nameResp.LastName, int, nil
}
