package model

import "encoding/json"

// QuoteRequest represents the request parameters for the Jupiter Quote API
type QuoteRequest struct {
	InputMint                  string   `url:"inputMint"`
	OutputMint                 string   `url:"outputMint"`
	Amount                     int64    `url:"amount"`      // lamports
	SlippageBps                int      `url:"slippageBps"` //  slippageBps: 100 = 1%
	SwapMode                   string   `url:"swapMode"`    // "ExactIn" or "ExactOut" default: "ExactIn"
	Dexes                      []string `url:"dexes"`       // will format to "Raydium,Orca+V2,Meteora+DLMM"
	ExcludeDexes               []string `url:"excludeDexes"`
	RestrictIntermediateTokens bool     `url:"restrictIntermediateTokens"`
	OnlyDirectRoutes           bool     `url:"onlyDirectRoutes"`
	AsLegacyTransaction        bool     `url:"asLegacyTransaction"`
	PlatformFeeBps             int      `url:"platformFeeBps"`
	MaxAccounts                int      `url:"maxAccounts"`
}

type QuoteResponse struct {
	InputMint            string      `json:"inputMint"`
	InAmount             string      `json:"inAmount"`
	OutputMint           string      `json:"outputMint"`
	OutAmount            string      `json:"outAmount"`
	OtherAmountThreshold string      `json:"otherAmountThreshold"`
	SwapMode             string      `json:"swapMode"`
	SlippageBps          int         `json:"slippageBps"`
	PlatformFee          interface{} `json:"platformFee"`
	PriceImpactPct       string      `json:"priceImpactPct"`
	RoutePlan            []RoutePlan `json:"routePlan"`
	ScoreReport          interface{} `json:"scoreReport"`
	ContextSlot          int64       `json:"contextSlot"`
	TimeTaken            float64     `json:"timeTaken"`
	SwapUsdValue         string      `json:"swapUsdValue"`
	SimplerRouteUsed     bool        `json:"simplerRouteUsed"`
}

func (r QuoteResponse) ToJson() (string, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

type RoutePlan struct {
	SwapInfo SwapInfo `json:"swapInfo"`
	Percent  int      `json:"percent"`
}

type SwapInfo struct {
	AmmKey     string `json:"ammKey"`
	Label      string `json:"label"`
	InputMint  string `json:"inputMint"`
	OutputMint string `json:"outputMint"`
	InAmount   string `json:"inAmount"`
	OutAmount  string `json:"outAmount"`
	FeeAmount  string `json:"feeAmount"`
	FeeMint    string `json:"feeMint"`
}

type SwapRequest struct {
	UserPublicKey                 string        `json:"userPublicKey"`
	WrapAndUnwrapSol              bool          `json:"wrapAndUnwrapSol,omitempty"`
	UseSharedAccounts             bool          `json:"useSharedAccounts,omitempty"`
	FeeAccount                    string        `json:"feeAccount,omitempty"`
	TrackingAccount               string        `json:"trackingAccount,omitempty"`
	PrioritizationFeeLamports     interface{}   `json:"prioritizationFeeLamports,omitempty"`
	AsLegacyTransaction           bool          `json:"asLegacyTransaction,omitempty"`
	DestinationTokenAccount       string        `json:"destinationTokenAccount,omitempty"`
	DynamicComputeUnitLimit       bool          `json:"dynamicComputeUnitLimit,omitempty"`
	SkipUserAccountsRpcCalls      bool          `json:"skipUserAccountsRpcCalls,omitempty"`
	DynamicSlippage               bool          `json:"dynamicSlippage,omitempty"`
	ComputeUnitPriceMicroLamports int           `json:"computeUnitPriceMicroLamports,omitempty"`
	QuoteResponse                 QuoteResponse `json:"quoteResponse"`
}

func (r SwapRequest) ToJson() (string, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// PrioritizationFeeLamports 定义了优先级费用的结构
type PrioritizationFeeLamports struct {
	PriorityLevelWithMaxLamports *PriorityLevelWithMaxLamports `json:"priorityLevelWithMaxLamports,omitempty"`
	JitoTipLamports              int                           `json:"jitoTipLamports,omitempty"`
}

// PriorityLevelWithMaxLamports 定义了优先级级别和最大 Lamports 的结构
type PriorityLevelWithMaxLamports struct {
	PriorityLevel string `json:"priorityLevel"`
	MaxLamports   int    `json:"maxLamports"`
}

type SwapResponse struct {
	SwapTransaction           string `json:"swapTransaction"`
	LastValidBlockHeight      int    `json:"lastValidBlockHeight"`
	PrioritizationFeeLamports int    `json:"prioritizationFeeLamports,omitempty"`
}
