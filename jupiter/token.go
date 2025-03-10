package jupiter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cathyonmoon/jupiter-api/jupiter/model"
)

func (c *Client) GetTokenPrice(ctx context.Context, addresses string) (*model.PriceAPIResponse, error) {
	resp, err := c.client.R().SetContext(ctx).SetQueryParams(map[string]string{
		"ids": addresses,
	}).Get("/price/v2")
	if err != nil {
		return nil, err
	}
	var priceResp model.PriceAPIResponse
	err = json.Unmarshal(resp.Body(), &priceResp)
	if err != nil {
		return nil, err
	}
	return &priceResp, nil
}

func (c *Client) GetTokenInfo(ctx context.Context, address string) (*model.TokenInfo, error) {
	var tokenInfo model.TokenInfo
	endpoint := fmt.Sprintf("/tokens/v1/token/%s", address)
	resp, err := c.client.R().SetContext(ctx).SetResult(&tokenInfo).Get(endpoint)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to get token info:%v", resp.StatusCode())
	}
	return &tokenInfo, nil
}
