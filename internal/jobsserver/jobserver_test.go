package jobsserver_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/adefirmanf/data_selection_pretexting/internal/jobsserver"
	"github.com/adefirmanf/data_selection_pretexting/internal/queue"
	"github.com/adefirmanf/data_selection_pretexting/internal/queue/linkedlist"
	scrappertweets "github.com/adefirmanf/data_selection_pretexting/internal/scrapper-tweets"
	scrapperusers "github.com/adefirmanf/data_selection_pretexting/internal/scrapper-users"
	"github.com/adefirmanf/data_selection_pretexting/internal/storage"
	"github.com/adefirmanf/data_selection_pretexting/internal/storage/postgresql"

	_ "github.com/lib/pq"
)

func TestScrapperTweetsJob(t *testing.T) {
	config := jobsserver.NewConfig(1, "days")
	pgclient := postgresql.NewConfig("postgres://postgres:social_engineering@localhost:5432/social_engineering?sslmode=disable")
	storage := storage.NewStorage(pgclient.OpenConnection())
	queue := queue.NewQueue(linkedlist.NewLinkedList())
	job := jobsserver.NewJobServer(config, storage, queue)

	scrapperTw := scrappertweets.NewConfig("https://api.twitter.com/2/tweets/search/recent", "")
	httpClient := http.Client{}

	qs := jobsserver.NewQueryURLS()
	qs.Add(scrappertweets.NewQueryURL("JeniusConnect", "Whatsapp|Mohon Maaf", "is:reply"))
	qs.Add(scrappertweets.NewQueryURL("bankmandiri", "Whatsapp|Mohon Maaf", "is:reply"))

	start := jobsserver.NewScrapperTweets(job, scrappertweets.NewScrapperTweets(scrapperTw, &httpClient), qs)
	err := start.Scrape(10)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(queue.Size())
}

func TestScrapperUsersJob(t *testing.T) {
	config := jobsserver.NewConfig(1, "days")
	pgclient := postgresql.NewConfig("postgres://postgres:social_engineering@localhost:5432/social_engineering?sslmode=disable")
	storage := storage.NewStorage(pgclient.OpenConnection())
	queue := queue.NewQueue(linkedlist.NewLinkedList())
	job := jobsserver.NewJobServer(config, storage, queue)

	scrapperUsrCfg := scrapperusers.NewConfig("https://api.twitter.com", "")
	httpClient := http.Client{}

	queue.PushBack("115848008")
	queue.PushBack("115848008")

	start := jobsserver.NewScrapperUsers(job, scrapperusers.NewScrapperUser(scrapperUsrCfg, &httpClient))

	err := start.Scrape(queue.PullFront())
	if err != nil {
		t.Error(err)
	}
}
