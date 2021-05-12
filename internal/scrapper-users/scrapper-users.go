package scrapperusers

import (
	"errors"
	"net/http"
	"time"

	"github.com/adefirmanf/data_selection_pretexting/internal/queue"
	"github.com/adefirmanf/data_selection_pretexting/internal/storage"
)

var (
	errorNotFound     = "HTTP Error is Not Found"
	rateLimitExceeded = "Rate Limit Exceeded"
	forbidden         = "Invalid Token / Forbidden access usage"
)

// UserResponse .
type userResponse struct {
	Errors interface{} `json:"errors"`
	Data   []data      `json:"data"`
}

type data struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	Verified  bool      `json:"verified"`
}

type friendshipResponse struct {
	Meta resultCount `json:"meta"`
}

type resultCount struct {
	ResultCount int `json:"result_count"`
}

// Config .
type Config struct {
	url         string
	bearerToken string
}

// NewConfig .
func NewConfig(url string, bearerToken string) *Config {
	return &Config{
		url:         url,
		bearerToken: bearerToken,
	}
}

// ScrapperUsers .
type ScrapperUsers struct {
	httpClient *http.Client
	storage    *storage.Storage
	queue      *queue.Queue
	config     *Config
}

// NewScrapperUser .
func NewScrapperUser(config *Config, httpClient *http.Client, storage *storage.Storage, queue *queue.Queue) *ScrapperUsers {
	return &ScrapperUsers{
		config:     config,
		httpClient: httpClient,
		storage:    storage,
		queue:      queue,
	}
}
func (s *ScrapperUsers) httpRequestUsers(UserID string, endpoint string) (*http.Response, error) {
	url := s.config.url + "/2/users/" + UserID + endpoint
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+s.config.bearerToken)
	return s.httpClient.Do(req)
}

func httpErrorClientHandler(response *http.Response) (*http.Response, error) {
	if response.StatusCode == http.StatusNotFound {
		return response, errors.New(errorNotFound)
	}
	if response.StatusCode == http.StatusTooManyRequests {
		return response, errors.New(rateLimitExceeded)
	}
	if response.StatusCode == http.StatusForbidden {
		return response, errors.New(forbidden)
	}
	return response, nil
}
