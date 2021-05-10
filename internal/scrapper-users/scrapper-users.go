package scrapperusers

import (
	"net/http"

	"github.com/adefirmanf/data_selection_pretexting/internal/queue"
	"github.com/adefirmanf/data_selection_pretexting/internal/storage"
)

// ScrapperUsers .
type ScrapperUsers struct {
	HTTPClient *http.Client
	Storage    *storage.Storage
	Queue      *queue.Queue
}

// FetchFollowings .
func (s *ScrapperUsers) FetchFollowings(UserID string) {

}

// FetchFollowers .
func (s *ScrapperUsers) FetchFollowers(UserID string) {

}

// FetchLookup .
func (s *ScrapperUsers) FetchLookup(UserID string) {

}

// NewScrapperUser .
func NewScrapperUser(httpClient *http.Client, storage *storage.Storage, queue *queue.Queue) *ScrapperUsers {
	return &ScrapperUsers{
		HTTPClient: httpClient,
		Storage:    storage,
		Queue:      queue,
	}
}
