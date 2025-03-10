package jupiter

import (
	"context"
	"fmt"
	"github.com/CathyOnMoon/jupiter-api/jupiter/model"
	"github.com/CathyOnMoon/jupiter-api/utils"
	"github.com/go-resty/resty/v2"
	"net/url"
)

func (c *Client) GetQuote(ctx context.Context, request model.QuoteRequest) (*model.QuoteResponse, *resty.Response, error) {
	params := url.Values{}
	for key, value := range utils.StructToValues(request) {
		params.Add(key, value)
	}
	quoteApi := "/swap/v1/quote"
	endpoint := fmt.Sprintf("%s%s?%s", c.client.BaseURL, quoteApi, params.Encode())
	var quoteResponse model.QuoteResponse
	resp, err := c.client.R().SetContext(ctx).SetResult(&quoteResponse).Get(endpoint)
	if err != nil {
		return nil, resp, err
	}
	if resp.IsError() {
		return nil, resp, fmt.Errorf("failed to get quote:%v", resp.StatusCode())
	}
	return &quoteResponse, resp, nil
}

func (c *Client) PostSwap(ctx context.Context, request model.SwapRequest) (*model.SwapResponse, *resty.Response, error) {
	swapApi := "/swap/v1/swap"
	endpoint := fmt.Sprintf("%s%s", c.client.BaseURL, swapApi)
	var swapResponse model.SwapResponse
	body, _ := request.ToJson()
	resp, err := c.client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&swapResponse).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}
	if resp.IsError() {
		return nil, resp, fmt.Errorf("failed to post swap:%v", resp.StatusCode())
	}
	return &swapResponse, resp, nil
}
