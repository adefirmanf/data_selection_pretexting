package scrappertweets_test

import (
	"net/http"
	"testing"

	scrappertweets "github.com/adefirmanf/data_selection_pretexting/internal/scrapper-tweets"
)

func TestScrapper(t *testing.T) {
	q := scrappertweets.NewQueryURL("JeniusConnect", "Whatsapp|Mohon Maaf", "")

	config := scrappertweets.NewConfig("https://api.twitter.com/2/tweets/search/recent", "")
	httpClient := http.Client{}
	scrapper := scrappertweets.NewScrapperTweets(config, &httpClient, nil, nil)
	err := scrapper.FetchTweets(q)
	if err != nil {
		t.Error(err)
	}

}
