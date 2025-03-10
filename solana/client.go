package solana

import (
	"context"
	"fmt"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"log"
)

const defaultMaxRetries = uint(20)

type Client struct {
	rpcClient *rpc.Client
}

func NewClient(rpcEndpoint string) (*Client, error) {
	if rpcEndpoint == "" {
		return nil, fmt.Errorf("rpcEndpoint is required when no RPC service is provided")
	}
	rpcClient := rpc.New(rpcEndpoint)
	c := &Client{
		rpcClient: rpcClient,
	}
	return c, nil
}

func (c *Client) SendTransaction(ctx context.Context, wallet *solana.Wallet, txBase64 string) (*solana.Signature, error) {
	// 反序列化交易
	tx, err := DecodeTransactionFromBase64(txBase64)
	if err != nil {
		return nil, fmt.Errorf("could not deserialize swap transaction: %w", err)
	}
	// 获取最近的区块哈希
	latestBlockhash, err := c.rpcClient.GetLatestBlockhash(ctx, "")
	if err != nil {
		return nil, fmt.Errorf("could not get latest blockhash: %w", err)
	}
	tx.Message.RecentBlockhash = latestBlockhash.Value.Blockhash
	// 对交易进行签名
	tx, err = SignTransaction(wallet, tx)
	if err != nil {
		return nil, fmt.Errorf("could not sign swap transaction: %w", err)
	}
	// 发送交易到rpc
	maxRetries := defaultMaxRetries
	sig, err := c.rpcClient.SendTransactionWithOpts(ctx, &tx, rpc.TransactionOpts{
		MaxRetries:          &maxRetries,
		MinContextSlot:      &latestBlockhash.Context.Slot,
		PreflightCommitment: rpc.CommitmentProcessed,
	})
	return &sig, err
}

type TxStatus struct {
	Status string
	Error  error
}

func (c *Client) CheckSignature(ctx context.Context, signatureStr string) (*TxStatus, error) {
	sig, err := solana.SignatureFromBase58(signatureStr)
	if err != nil {
		return nil, fmt.Errorf("could not convert signature from base58: %w", err)
	}

	status, err := c.rpcClient.GetSignatureStatuses(ctx, false, sig)
	if err != nil {
		return nil, fmt.Errorf("could not get signature status: %w", err)
	}

	if len(status.Value) == 0 {
		return nil, fmt.Errorf("could not confirm transaction: no valid status")
	}

	if status.Value[0] == nil || status.Value[0].ConfirmationStatus != rpc.ConfirmationStatusFinalized {
		// 交易正在处理中
		return &TxStatus{
			Status: string(rpc.ConfirmationStatusProcessed),
			Error:  nil,
		}, nil
	}

	if status.Value[0].Err != nil {
		// 交易已完成，但是未被确认
		return &TxStatus{
			Status: string(rpc.ConfirmationStatusFinalized),
			Error:  fmt.Errorf("transaction confirmed with error: %s", status.Value[0].Err),
		}, nil
	}

	// 交易已完成，并且被确认
	return &TxStatus{
		Status: string(rpc.ConfirmationStatusConfirmed),
		Error:  nil,
	}, nil
}

func (c *Client) GetWalletAccount(ctx context.Context, walletAddress string) (*rpc.Account, error) {
	account := solana.MustPublicKeyFromBase58(walletAddress)
	result, err := c.rpcClient.GetAccountInfo(ctx, account)
	if err != nil {
		return nil, fmt.Errorf("获取钱包账户失败: %w", err)
	}
	return result.Value, nil
}

func (c *Client) GetWalletTokenAccount(ctx context.Context, walletAddress string, tokenAddress string) (*token.Account, error) {
	var walletAccount = solana.MustPublicKeyFromBase58(walletAddress)
	var tokenPubkey = solana.MustPublicKeyFromBase58(tokenAddress)
	ata, _, err := solana.FindAssociatedTokenAddress(walletAccount, tokenPubkey)
	if err != nil {
		log.Printf("failed to get associated token address: %v", err)
		return nil, err
	}
	result, err := c.rpcClient.GetAccountInfo(ctx, ata)
	if err != nil {
		return nil, fmt.Errorf("获取token账户失败: %w", err)
	}
	var tokenAccount token.Account
	err = tokenAccount.UnmarshalWithDecoder(bin.NewBinDecoder(result.GetBinary()))
	if err != nil {
		return nil, fmt.Errorf("failed to decode token account data: %v", err)
	}

	return &tokenAccount, nil
}
