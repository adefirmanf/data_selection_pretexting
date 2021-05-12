package scrapperusers_test

import (
	"fmt"
	"net/http"
	"testing"

	scrapperusers "github.com/adefirmanf/data_selection_pretexting/internal/scrapper-users"
)

func TestScrapperUserLookup(t *testing.T) {
	type test struct {
		name     string
		url      string
		expected string
		is_error bool
	}

	tests := []test{
		{
			name:     "it should pass when success to get data",
			url:      "https://run.mocky.io/v3/b7bdb1c1-e2e0-4a52-9f98-bb8f0bc97e53",
			expected: "",
		},
		{
			name:     "it should error when failed to get data",
			url:      "https://run.mocky.io/v3/b7bdb1c1-e2e0-4a52-9f98-bb8f0bc97e53",
			expected: "",
			is_error: true,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			config := scrapperusers.NewConfig(v.url, "")
			httpClient := http.Client{}
			scrapper := scrapperusers.NewScrapperUser(config, &httpClient, nil, nil)
			err := scrapper.FetchLookup("1391596003908673539")

			if v.is_error {
				if err != nil {
					t.Fail()
				}
			}
		})
	}
}
func TestScrapperUserFollowers(t *testing.T) {
	// successURL := "https://run.mocky.io/v3/a758e8b8-a791-499e-99d6-e7a220e3f14a"
	errorURL := "https://run.mocky.io/v3/c5d2b804-3747-4422-b4da-ecb772ddaec7"

	url := errorURL
	config := scrapperusers.NewConfig(url, "")
	httpClient := http.Client{}
	scrapper := scrapperusers.NewScrapperUser(config, &httpClient, nil, nil)
	err := scrapper.FetchFollowers("1391596003908673539")
	fmt.Println(err)
}
