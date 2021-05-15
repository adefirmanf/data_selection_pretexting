package app

import (
	"net/http"

	"github.com/adefirmanf/data_selection_pretexting/internal/queue"
	"github.com/adefirmanf/data_selection_pretexting/internal/queue/linkedlist"
	scrappertweets "github.com/adefirmanf/data_selection_pretexting/internal/scrapper-tweets"
	scrapperusers "github.com/adefirmanf/data_selection_pretexting/internal/scrapper-users"
	"github.com/adefirmanf/data_selection_pretexting/internal/storage"
	"github.com/adefirmanf/data_selection_pretexting/internal/storage/postgresql"
)

// App holds all services required by application.
// i.e Logger, Metrics, or domain business
type App struct {
	*Metrics
	ScrapperTweets *scrappertweets.ScrapperTweets
	ScrapperUsers  *scrapperusers.ScrapperUsers
	Storage        *storage.Storage
	Queue          *queue.Queue
}

// Metrics .
type Metrics struct {
}

// Logger .
type Logger struct{}

// MetricsApp .
func buildMetricsApp() *Metrics {
	return &Metrics{}
}

// New .
func New() *App {
	token := "AAAAAAAAAAAAAAAAAAAAAKGDOQEAAAAAlrZppByCtx%2FgMDwOxPtUq9%2B1rgk%3D59cCihRFivSTtY9RdUcAokS8yhwTWQJ93RvFIjre8mdJtuLIAl"
	pgconfig := postgresql.NewConfig("postgres://postgres:social_engineering@localhost:5432/social_engineering?sslmode=disable")
	ll := linkedlist.NewLinkedList()
	scrapperTweetsCfg := scrappertweets.NewConfig("https://api.twitter.com/2/tweets/search/recent", token)
	scrapperUsersCfg := scrapperusers.NewConfig("https://api.twitter.com", token)

	httpClient := http.Client{}

	return &App{
		Storage:        storage.NewStorage(pgconfig.OpenConnection()),
		ScrapperTweets: scrappertweets.NewScrapperTweets(scrapperTweetsCfg, &httpClient),
		ScrapperUsers:  scrapperusers.NewScrapperUser(scrapperUsersCfg, &httpClient),
		Queue:          queue.NewQueue(ll),
		Metrics:        buildMetricsApp(),
	}
}
