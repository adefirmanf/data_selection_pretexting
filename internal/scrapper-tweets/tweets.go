package scrappertweets

// FetchTweets .
func (s *ScrapperTweets) FetchTweets(q *QueryURL) error {
	_, err := s.httpRequestTweets(*q)
	if err != nil {
		return err
	}

	return nil
}
