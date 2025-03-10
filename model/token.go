package model

import "time"

type PriceAPIResponse struct {
	Data      map[string]TokenPrice `json:"data"`
	TimeTaken float64               `json:"timeTaken"`
}

type TokenPrice struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Price string `json:"price"`
}

type TokenInfo struct {
	Address           string     `json:"address"`
	Name              string     `json:"name"`
	Symbol            string     `json:"symbol"`
	Decimals          int        `json:"decimals"`
	LogoURI           string     `json:"logoURI"`
	Tags              []string   `json:"tags"`
	DailyVolume       *float64   `json:"daily_volume"` // 使用指针类型表示可能为 null
	CreatedAt         time.Time  `json:"created_at"`
	FreezeAuthority   *string    `json:"freeze_authority"`   // 使用指针类型表示可能为 null
	MintAuthority     *string    `json:"mint_authority"`     // 使用指针类型表示可能为 null
	PermanentDelegate *string    `json:"permanent_delegate"` // 使用指针类型表示可能为 null
	MintedAt          *time.Time `json:"minted_at"`          // 使用指针类型表示可能为 null
}
