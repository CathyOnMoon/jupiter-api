package client

import (
	"context"
	"fmt"
	"jupiter-api/model"
	"net/url"
)

func (c *Client) CreateUltraOrder(ctx context.Context, request model.CreateOrderRequest) (*model.CreateOrderResponse, error) {
	params := url.Values{}
	params.Add("inputMint", request.InputMint)
	params.Add("outputMint", request.OutputMint)
	params.Add("amount", fmt.Sprintf("%d", request.Amount))
	params.Add("taker", request.Taker)
	endpoint := fmt.Sprintf("%s/ultra/v1/order?%s", c.config.BaseURL, params.Encode())
	var orderResponse model.CreateOrderResponse
	resp, err := c.client.R().SetContext(ctx).SetResult(&orderResponse).Get(endpoint)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to create order:%v", resp.StatusCode())
	}
	return &orderResponse, nil
}

func (c *Client) ExecuteUltraOrder(ctx context.Context, request model.ExecuteOrderRequest) (*model.ExecuteOrderResponse, error) {
	endpoint := fmt.Sprintf("%s/ultra/v1/execute", c.client.BaseURL)
	var executeOrderResponse model.ExecuteOrderResponse
	resp, err := c.client.R().
		SetContext(ctx).
		SetBody(map[string]string{
			"signedTransaction": request.SignedTransaction,
			"requestId":         request.RequestId,
		}).
		SetResult(&executeOrderResponse).
		Post(endpoint)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to execute order:%v", resp.StatusCode())
	}
	return &executeOrderResponse, nil
}
