/*
Copyright © 2023 Rebecca [Becci] Govert <becci.govert@gmail.com>
*/
package clients

import (
	"encoding/json"
	"fmt"
)

type NameResponse struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}
type GetNameResp struct {
	NameResponse  *NameResponse
	StatusCode    int
	ErrorResponse error
}

// GetRandomName retrieves a random name from APIa
func GetRandomName() *GetNameResp {
	resp := get(nameUrl)
	if resp.Error != nil {
		return &GetNameResp{
			NameResponse:  nil,
			StatusCode:    resp.StatusCode,
			ErrorResponse: resp.Error,
		}
	}
	var nameResp *NameResponse
	err := json.Unmarshal(resp.Response, &nameResp)
	if err != nil {
		return &GetNameResp{
			NameResponse:  nil,
			StatusCode:    500,
			ErrorResponse: fmt.Errorf("failed to unmarshal name %w", err),
		}

	}
	return &GetNameResp{
		NameResponse:  nameResp,
		StatusCode:    200,
		ErrorResponse: nil,
	}
}
