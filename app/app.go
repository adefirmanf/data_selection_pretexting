package app

import (
	"net/http"

	"github.com/adefirmanf/data_selection_pretexting/config"
	"github.com/adefirmanf/data_selection_pretexting/config/env"
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
	config.Init(env.New())
	cfg := config.Load()

	token := cfg.BearerTokenTwitter()
	pgconfig := postgresql.NewConfig(cfg.PGConfigConnectionString())
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
