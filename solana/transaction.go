package solana

import (
	"encoding/base64"
	"fmt"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

// DecodeTransactionFromBase64 反序列化交易
func DecodeTransactionFromBase64(txStr string) (solana.Transaction, error) {
	txBytes, err := base64.StdEncoding.DecodeString(txStr)
	if err != nil {
		return solana.Transaction{}, fmt.Errorf("could not decode transaction: %w", err)
	}

	tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(txBytes))
	if err != nil {
		return solana.Transaction{}, fmt.Errorf("could not deserialize transaction: %w", err)
	}

	return *tx, nil
}

// SignTransaction 签名交易
func SignTransaction(wallet *solana.Wallet, tx solana.Transaction) (solana.Transaction, error) {
	txMessageBytes, err := tx.Message.MarshalBinary()
	if err != nil {
		return solana.Transaction{}, fmt.Errorf("could not serialize transaction: %w", err)
	}

	signature, err := wallet.PrivateKey.Sign(txMessageBytes)
	if err != nil {
		return solana.Transaction{}, fmt.Errorf("could not sign transaction: %w", err)
	}

	tx.Signatures = []solana.Signature{signature}

	return tx, nil
}
