package scrappertweets

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	errorNotFound     = "HTTP Error is Not Found"
	rateLimitExceeded = "Rate Limit Exceeded"
	forbidden         = "Forbidden access usage"
	unauthorized      = "Unauthorized / Invalid token "
)

// TweetResponse .
type TweetResponse struct {
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
	NextToken          string
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

	if q.NextToken != "" {
		query.Set("next_token", q.NextToken)

	}
	query.Set("query", rawquery)

	return query.Encode()
}

// SetNextToken .
func (q *QueryURL) SetNextToken(nextToken string) {
	q.NextToken = nextToken
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
	config     *Config
}

// NewScrapperTweets .
func NewScrapperTweets(config *Config, httpClient *http.Client) *ScrapperTweets {
	return &ScrapperTweets{
		config:     config,
		httpClient: httpClient,
	}
}

func (s *ScrapperTweets) httpRequestTweets(q QueryURL) (*http.Response, error) {
	url := s.config.url + "?tweet.fields=author_id,id,created_at,possibly_sensitive&" + q.Encode()

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
	if response.StatusCode == http.StatusUnauthorized {
		return response, errors.New(unauthorized)
	}
	return response, nil
}
