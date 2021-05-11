package scrappertweets

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/adefirmanf/data_selection_pretexting/internal/queue"
	"github.com/adefirmanf/data_selection_pretexting/internal/storage"
)

type tweetResponse struct {
	Data []data `json:"data"`
	Meta meta   `json:"meta"`
}

type data struct {
	Text              string    `json:"text"`
	PossiblySensitive bool      `json:"possibly_sensitive"`
	AuthorID          string    `json:"author_id"`
	CreatedAt         time.Time `json:"created_at"`
	TweetID           string    `json:"id"`
}

type meta struct {
	ResultCount int    `json:"result_count"`
	NextToken   string `json:"next_token"`
}

// QueryURL .
type QueryURL struct {
	MentionedAccount string
	// You can set multiple keywords by using delimiter with comma
	SuspiciousKeywords string
	AdditionalParams   string
}

// NewQueryURL .
func NewQueryURL(mentionedAccount, suspiciousKeywords, additionalParams string) *QueryURL {
	return &QueryURL{
		MentionedAccount:   mentionedAccount,
		SuspiciousKeywords: suspiciousKeywords,
		AdditionalParams:   additionalParams,
	}
}

// Encode .
func (q *QueryURL) Encode() string {
	suspiciousKeywords := strings.Split(q.SuspiciousKeywords, "|")
	rawquery := fmt.Sprintf("@%s (%s) %s", q.MentionedAccount, strings.Join(suspiciousKeywords, " OR "), q.AdditionalParams)

	var s url.URL
	query := s.Query()
	query.Set("query", rawquery)
	return query.Encode()
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

// ScrapperTweets .
type ScrapperTweets struct {
	httpClient *http.Client
	storage    *storage.Storage
	queue      *queue.Queue
	config     *Config
}

// NewScrapperTweets .
func NewScrapperTweets(config *Config, httpClient *http.Client, storage *storage.Storage, queue *queue.Queue) *ScrapperTweets {
	return &ScrapperTweets{
		config:     config,
		httpClient: httpClient,
		storage:    storage,
		queue:      queue,
	}
}

// func (s *ScrapperTweets) httpRequestTweets(Query string, endpoint string, u url.URL) (*http.Response, error) {
// 	url := s.config.url + "/2/users/" + UserID + endpoint
// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Add("Authorization", "Bearer "+s.config.bearerToken)
// 	return s.httpClient.Do(req)
// }
