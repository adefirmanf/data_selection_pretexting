package scrapperusers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// FetchLookup .
func (s *ScrapperUsers) FetchLookup(UserID string) error {
	res, err := s.httpRequestUsers(UserID, "")
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var data userResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	if data.Errors != nil {
		return errors.New("Failed to retrieve data")
	}

	for _, v := range data.Data {
		fmt.Println(v.Username)
	}

	return nil
}
