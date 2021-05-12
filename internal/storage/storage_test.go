package storage_test

import (
	"fmt"
	"testing"

	"github.com/adefirmanf/data_selection_pretexting/internal/storage"
	"github.com/adefirmanf/data_selection_pretexting/internal/storage/postgresql"
	_ "github.com/lib/pq"
)

func TestStoragePostgresql(t *testing.T) {
	pg := postgresql.NewConfig("postgres://postgres:social_engineering@localhost:5432/social_engineering?sslmode=disable")
	storage := storage.NewStorage(pg.OpenConnection())

	tweets, err := storage.GetTweets()
	if err != nil {
		t.Error(err)
	}
	for _, v := range tweets {
		fmt.Println(&v.TweetText)
	}
	// err = storage.InsertTweets("12", "1212", "1212", "test", "bri", "", 0, time.Now(), false)
	if err != nil {
		t.Error(err)
	}
	// s := storage.NewStorage()
}
