package model

import "encoding/json"

type CreateOrderRequest struct {
	InputMint  string
	OutputMint string
	Amount     int64
	Taker      string
}

type CreateOrderResponse struct {
	SwapType                  string                 `json:"swapType"`
	Environment               string                 `json:"environment"`
	RequestID                 string                 `json:"requestId"`
	InAmount                  string                 `json:"inAmount"`
	OutAmount                 string                 `json:"outAmount"`
	OtherAmountThreshold      string                 `json:"otherAmountThreshold"`
	SwapMode                  string                 `json:"swapMode"`
	SlippageBps               int                    `json:"slippageBps"`
	PriceImpactPct            string                 `json:"priceImpactPct"`
	RoutePlan                 []RoutePlan            `json:"routePlan"`
	InputMint                 string                 `json:"inputMint"`
	OutputMint                string                 `json:"outputMint"`
	FeeBps                    int                    `json:"feeBps,omitempty"`
	PlatformFee               *PlatformFee           `json:"platformFee"`
	Taker                     string                 `json:"taker,omitempty"`
	Gasless                   bool                   `json:"gasless,omitempty"`
	Transaction               string                 `json:"transaction,omitempty"`
	PrioritizationType        string                 `json:"prioritizationType,omitempty"`
	PrioritizationFeeLamports int                    `json:"prioritizationFeeLamports,omitempty"`
	LastValidBlockHeight      int                    `json:"lastValidBlockHeight,omitempty"`
	DynamicSlippageReport     *DynamicSlippageReport `json:"dynamicSlippageReport,omitempty"`
	TotalTime                 int                    `json:"totalTime"`
}

func (r CreateOrderResponse) ToJson() (string, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

type DynamicSlippageReport struct {
	SlippageBps                  int     `json:"slippageBps,omitempty"`
	OtherAmount                  *int    `json:"otherAmount,omitempty"`
	SimulatedIncurredSlippageBps *int    `json:"simulatedIncurredSlippageBps,omitempty"`
	AmplificationRatio           *string `json:"amplificationRatio,omitempty"`
	CategoryName                 string  `json:"categoryName,omitempty"`
	HeuristicMaxSlippageBps      int     `json:"heuristicMaxSlippageBps,omitempty"`
	RtseSlippageBps              int     `json:"rtseSlippageBps,omitempty"`
	FailedTxnEstSlippage         int     `json:"failedTxnEstSlippage,omitempty"`
	PriceMovementEstSlippage     int     `json:"priceMovementEstSlippage,omitempty"`
	EmaEstSlippage               int     `json:"emaEstSlippage,omitempty"`
}

type PlatformFee struct {
	Amount string `json:"amount,omitempty"`
	FeeBps int    `json:"feeBps,omitempty"`
}

type ExecuteOrderRequest struct {
	SignedTransaction string
	RequestId         string
}

type ExecuteOrderResponse struct {
	Status             string      `json:"status"`
	Signature          string      `json:"signature"`
	Slot               string      `json:"slot"`
	Code               int         `json:"code"`
	Error              string      `json:"error,omitempty"`              // 失败时存在
	InputAmountResult  string      `json:"inputAmountResult,omitempty"`  // 成功时存在
	OutputAmountResult string      `json:"outputAmountResult,omitempty"` // 成功时存在
	SwapEvents         []SwapEvent `json:"swapEvents,omitempty"`         // 成功时存在
}

type SwapEvent struct {
	InputMint    string `json:"inputMint"`
	InputAmount  string `json:"inputAmount"`
	OutputMint   string `json:"outputMint"`
	OutputAmount string `json:"outputAmount"`
}

func (r ExecuteOrderResponse) ToJson() (string, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
