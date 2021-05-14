package jobsserver_test

import (
	"net/http"
	"testing"

	"github.com/adefirmanf/data_selection_pretexting/internal/jobsserver"
	"github.com/adefirmanf/data_selection_pretexting/internal/queue"
	"github.com/adefirmanf/data_selection_pretexting/internal/queue/linkedlist"
	scrappertweets "github.com/adefirmanf/data_selection_pretexting/internal/scrapper-tweets"
	"github.com/adefirmanf/data_selection_pretexting/internal/storage"
	"github.com/adefirmanf/data_selection_pretexting/internal/storage/postgresql"

	_ "github.com/lib/pq"
)

func TestScrapperJob(t *testing.T) {
	config := jobsserver.NewConfig(1, "days")
	pgclient := postgresql.NewConfig("postgres://postgres:social_engineering@localhost:5432/social_engineering?sslmode=disable")
	storage := storage.NewStorage(pgclient.OpenConnection())
	queue := queue.NewQueue(linkedlist.NewLinkedList())
	job := jobsserver.NewJobServer(config, storage, queue)

	scrapperTw := scrappertweets.NewConfig("https://api.twitter.com/2/tweets/search/recent", "")
	httpClient := http.Client{}

	qs := jobsserver.NewQueryURLS()
	qs.Add(scrappertweets.NewQueryURL("JeniusConnect", "Whatsapp|Mohon Maaf", ""))
	qs.Add(scrappertweets.NewQueryURL("bankmandiri", "Whatsapp|Mohon Maaf", ""))

	start := jobsserver.NewScrapperTweets(scrappertweets.NewScrapperTweets(scrapperTw, &httpClient), job, qs)
	err := start.Scrape(10)
	if err != nil {
		t.Error(err)
	}
}
