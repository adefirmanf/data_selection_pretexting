package jobsserver

import (
	"log"

	scrappertweets "github.com/adefirmanf/data_selection_pretexting/internal/scrapper-tweets"
)

// QueryURLS .
type QueryURLS struct {
	list []*scrappertweets.QueryURL
}

// NewQueryURLS .
func NewQueryURLS() *QueryURLS {
	return &QueryURLS{}
}

// Add .
func (q *QueryURLS) Add(query *scrappertweets.QueryURL) {
	q.list = append(q.list, query)
}

// ScrapperTweets .
type ScrapperTweets struct {
	jobserver             *JobServer
	queryURLS             *QueryURLS
	scrapperServiceTweets *scrappertweets.ScrapperTweets
}

// NewScrapperTweets .
func NewScrapperTweets(s *scrappertweets.ScrapperTweets, job *JobServer, qs *QueryURLS) *ScrapperTweets {
	return &ScrapperTweets{
		jobserver:             job,
		queryURLS:             qs,
		scrapperServiceTweets: s,
	}
}

// Scrape .
func (s *ScrapperTweets) Scrape(maxBatch int) error {
	for _, queryURL := range s.queryURLS.list {
		nextToken := ""
		i := 0
		// queryURL = queryURL.
		for i < maxBatch {

			queryURL.SetNextToken(nextToken)
			res, err := s.scrapperServiceTweets.FetchTweets(queryURL)
			if err != nil {
				return err
			}
			if res.Meta.NextToken == "" {
				nextToken = ""
				break
			}
			for _, data := range res.Data {
				s.jobserver.storage.InsertTweets(data.TweetID, data.AuthorID, data.Text, queryURL.SuspiciousKeywords, queryURL.MentionedAccount, queryURL.AdditionalParams, 0, data.CreatedAt, data.PossiblySensitive)
			}
			log.Println(res)
			nextToken = res.Meta.NextToken
			i = i + 1
		}
	}
	return nil
}
