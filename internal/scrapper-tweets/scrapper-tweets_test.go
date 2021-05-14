package scrappertweets_test

import (
	"fmt"
	"net/http"
	"testing"

	scrappertweets "github.com/adefirmanf/data_selection_pretexting/internal/scrapper-tweets"
)

func TestScrapper(t *testing.T) {
	q := scrappertweets.NewQueryURL("JeniusConnect", "Whatsapp|Mohon Maaf", "")
	q.SetNextToken("b26v89c19zqg8o3foswqkkqydkzmbfk9ccajjpe0t5dh9")
	fmt.Println(q.Encode())
	config := scrappertweets.NewConfig("https://api.twitter.com/2/tweets/search/recent", "")
	httpClient := http.Client{}
	scrapper := scrappertweets.NewScrapperTweets(config, &httpClient)
	res, err := scrapper.FetchTweets(q)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
