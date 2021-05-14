package scrapperusers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// FetchLookup .
func (s *ScrapperUsers) FetchLookup(UserID string) (*UserResponse, error) {
	res, err := s.httpRequestUsers(UserID, "")
	if err != nil {
		return nil, err
	}

	res, err = httpErrorClientHandler(res)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data UserResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if data.Errors != nil {
		return nil, errors.New("Failed to retrieve data")
	}
	return &data, nil
}
