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
		BaseURL: "https://public.jupiterapi.com",
		Debug:   true,
	})
	quote, _, err := client.GetQuote(context.Background(), model.QuoteRequest{
		InputMint:  "So11111111111111111111111111111111111111112",
		OutputMint: "4eU17A7HLhCTjgF3jjthyonYQXQe2ofnqD4RucDmpump",
		Amount:     0.1 * 10e8,
	})
	if err != nil {
		return
	}
	log.Println(quote.ToJson())
}
