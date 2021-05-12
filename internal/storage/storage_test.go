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

	res, err := storage.GetUserByUserAuthorID("1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
	// s := storage.NewStorage()
}
