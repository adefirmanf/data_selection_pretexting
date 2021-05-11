package scrapperusers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// FetchFollowers .
func (s *ScrapperUsers) FetchFollowers(UserID string) error {
	res, err := s.httpRequestUsers(UserID, "/followers")
	if err != nil {
		return err
	}
	res, err = httpErrorClientHandler(res)
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
	// fmt.Println(data.Data)
	return nil
}
