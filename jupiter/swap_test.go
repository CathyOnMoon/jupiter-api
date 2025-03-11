package jupiter

import (
	"context"
	"github.com/cathyonmoon/jupiter-api/jupiter/model"
	"log"
	"testing"
)

func TestClient_GetQuote(t *testing.T) {
	client, _ := NewClient(Config{
		ApiKey:  "",
		BaseURL: "https://api.jup.ag",
		Debug:   true,
		//Proxy:   "https://7Mh7Hyrdx3Hb:QbFAPbhgzjePQ9U@superproxy.zenrows.com:1338",
	})
	quote, _, err := client.GetQuote(context.Background(), model.QuoteRequest{
		InputMint:  "So11111111111111111111111111111111111111112",
		OutputMint: "7HDrQfWp7gpyWPbnJKYmKAqsAkevZAQayB5wF6zppump",
		Amount:     0.1 * 10e8,
	})
	if err != nil {
		return
	}
	log.Println(quote.ToJson())
}
