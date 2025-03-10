package solana

import (
	"context"
	"log"
	"testing"
)

func TestClient_GetWalletAccount(t *testing.T) {
	client, _ := NewClient("https://virulent-wild-sky.solana-mainnet.quiknode.pro/f5c66c5f51866f6a1e049e744292bf21b3894b7d/")
	account, err := client.GetWalletAccount(context.Background(), "7ECrW6Ecj7ucfNzzDsFj8PYpdgUH2ruyGVC2N8FAHMWt")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(account.Lamports)
}

func TestClient_GetWalletTokenAccount(t *testing.T) {
	ctx := context.Background()
	client, _ := NewClient("https://virulent-wild-sky.solana-mainnet.quiknode.pro/f5c66c5f51866f6a1e049e744292bf21b3894b7d/")
	account, err := client.GetWalletTokenAccount(ctx, "7ECrW6Ecj7ucfNzzDsFj8PYpdgUH2ruyGVC2N8FAHMWt", "3hnUeTyU22B8KQyeRfremEfYsabY4YDXNcPuwZgCpump")
	if err != nil {
		t.Error(err)
		return
	}
	log.Println(account.Amount)

	//result, err := client.rpcClient.GetTokenAccountBalance(ctx, account.Owner, rpc.CommitmentFinalized)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//t.Log(result.Value.Amount)
}
