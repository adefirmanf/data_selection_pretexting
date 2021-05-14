package scrappertweets

import (
	"encoding/json"
	"io/ioutil"
)

// FetchTweets .
func (s *ScrapperTweets) FetchTweets(q *QueryURL) (*TweetResponse, error) {
	res, err := s.httpRequestTweets(*q)
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
	var data TweetResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func setUniqueTwitterAuthorID(TweetAuthorID string) error {
	return nil
}

func pushToQueue(TweetAuthorID string) {

}
