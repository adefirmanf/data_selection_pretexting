package scrappertweets_test

import (
	"fmt"
	"testing"

	scrappertweets "github.com/adefirmanf/data_selection_pretexting/internal/scrapper-tweets"
)

func TestScrapper(t *testing.T) {
	q := scrappertweets.NewQueryURL("kontakBRI", "Whatsapp|Mohon Maaf", "")
	fmt.Println(q.Encode())

}
