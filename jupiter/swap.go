package jupiter

import (
	"context"
	"fmt"
	"github.com/cathyonmoon/jupiter-api/jupiter/model"
	"github.com/cathyonmoon/jupiter-api/utils"
	"net/url"
)

func (c *Client) GetQuote(ctx context.Context, request model.QuoteRequest) (*model.QuoteResponse, error) {
	params := url.Values{}
	for key, value := range utils.StructToValues(request) {
		params.Add(key, value)
	}
	quoteApi := "/quote"
	endpoint := fmt.Sprintf("%s%s?%s", c.client.BaseURL, quoteApi, params.Encode())
	var quoteResponse model.QuoteResponse
	resp, err := c.client.R().SetContext(ctx).SetResult(&quoteResponse).Get(endpoint)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to get quote. StatusCode: %d. Resp: %s", resp.StatusCode(), resp.String())
	}
	return &quoteResponse, nil
}

func (c *Client) PostSwap(ctx context.Context, request model.SwapRequest) (*model.SwapResponse, error) {
	swapApi := "/swap"
	endpoint := fmt.Sprintf("%s%s", c.client.BaseURL, swapApi)
	var swapResponse model.SwapResponse
	body, err := request.ToJson()
	if err != nil {
		return nil, err
	}
	resp, err := c.client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetResult(&swapResponse).
		Post(endpoint)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("failed to post swap. StatusCode: %d. Resp: %s", resp.StatusCode(), resp.String())
	}
	return &swapResponse, nil
}
