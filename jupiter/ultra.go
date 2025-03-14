package jupiter

import (
	"context"
	"fmt"
	"github.com/cathyonmoon/jupiter-api/jupiter/model"
	"github.com/go-resty/resty/v2"
	"net/url"
)

func (c *Client) CreateUltraOrder(ctx context.Context, request model.CreateOrderRequest) (*model.CreateOrderResponse, *resty.Response, error) {
	params := url.Values{}
	params.Add("inputMint", request.InputMint)
	params.Add("outputMint", request.OutputMint)
	params.Add("amount", fmt.Sprintf("%d", request.Amount))
	params.Add("taker", request.Taker)
	endpoint := fmt.Sprintf("%s/ultra/v1/order?%s", c.config.BaseURL, params.Encode())
	var orderResponse model.CreateOrderResponse
	resp, err := c.client.R().SetContext(ctx).SetResult(&orderResponse).Get(endpoint)
	if err != nil {
		return nil, resp, err
	}
	return &orderResponse, resp, nil
}

func (c *Client) ExecuteUltraOrder(ctx context.Context, request model.ExecuteOrderRequest) (*model.ExecuteOrderResponse, *resty.Response, error) {
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
		return nil, resp, err
	}
	return &executeOrderResponse, resp, nil
}
