package scrapperusers

import (
	"encoding/json"
	"io/ioutil"
)

// FetchFollowings .
func (s *ScrapperUsers) FetchFollowings(UserID string) (*FriendshipResponse, error) {
	res, err := s.httpRequestUsers(UserID, "/following")
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
	var data FriendshipResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	// fmt.Println(data.Data)
	return &data, nil
}
