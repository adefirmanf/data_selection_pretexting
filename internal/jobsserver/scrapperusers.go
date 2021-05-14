package jobsserver

import (
	scrapperusers "github.com/adefirmanf/data_selection_pretexting/internal/scrapper-users"
)

// ScrapperUsers .
type ScrapperUsers struct {
	jobserver             *JobServer
	scrapperServiceTweets *scrapperusers.ScrapperUsers
}

// NewScrapperUsers .
func NewScrapperUsers(job *JobServer, scrapper *scrapperusers.ScrapperUsers) *ScrapperUsers {
	return &ScrapperUsers{
		jobserver:             job,
		scrapperServiceTweets: scrapper,
	}
}

// Scrape .
func (s *ScrapperUsers) Scrape(UserID string) error {
	res, err := s.scrapperServiceTweets.FetchLookup(UserID)
	if err != nil {
		s.jobserver.queue.PushBack(UserID)
		return err
	}
	userResponse := res.Data

	s.jobserver.storage.InsertUser(userResponse.ID, userResponse.Username, userResponse.Name, userResponse.PublicMetrics.FollowingCount, userResponse.PublicMetrics.FollowersCount, userResponse.Verified, userResponse.CreatedAt)
	if err != nil {
		s.jobserver.queue.PushBack(UserID)
		return err
	}
	return nil
}
