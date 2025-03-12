package jupiter

import (
	"context"
	"github.com/cathyonmoon/jupiter-api/jupiter/model"
	"log"
	"testing"
)

func TestClient_CreateUltraOrder(t *testing.T) {
	client, _ := NewClient(Config{
		ApiKey:  "",
		BaseURL: "https://api.jup.ag",
		Debug:   true,
		//Proxy:   "https://7Mh7Hyrdx3Hb:QbFAPbhgzjePQ9U@superproxy.zenrows.com:1338",
	})
	result, err, _ := client.CreateUltraOrder(context.Background(), model.CreateOrderRequest{
		InputMint:  "So11111111111111111111111111111111111111112",
		OutputMint: "7D9QsBnFcviLZbA8u7n6mF5k6NmN1pb3ZQwz9uiPpump",
		Amount:     0.1 * 10e8,
		Taker:      "7ECrW6Ecj7ucfNzzDsFj8PYpdgUH2ruyGVC2N8FAHMWt",
	})
	log.Println(result.ToJson())
	if err != nil {
		return
	}
}
