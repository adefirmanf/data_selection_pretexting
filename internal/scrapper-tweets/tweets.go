package scrappertweets

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// FetchTweets .
func (s *ScrapperTweets) FetchTweets(q *QueryURL) error {
	res, err := s.httpRequestTweets(*q)
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
	var data tweetResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	for _, v := range data.Data {
		fmt.Printf("ID : %s \n", v.TweetID)
		fmt.Printf("Tweet : %s", v.Text)
	}
	fmt.Println(res)

	return nil
}

func setUniqueTwitterAuthorID(TweetAuthorID string) error {
	return nil
}

func pushToQueue(TweetAuthorID string) {

}
