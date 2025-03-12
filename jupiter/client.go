package jupiter

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"time"
)

type Client struct {
	client *resty.Client
	config Config
}

type Config struct {
	BaseURL string
	ApiKey  string
	TimeOut time.Duration
	Debug   bool
	Proxy   string
}

func NewClient(config Config) (*Client, error) {
	client := resty.New()
	if config.BaseURL == "" {
		return nil, errors.New("baseURL is required")
	}
	client.BaseURL = config.BaseURL
	client.Debug = config.Debug
	client.SetHeader("x-api-key", config.ApiKey)
	client.SetTimeout(config.TimeOut)
	if config.Proxy != "" {
		client.SetProxy(config.Proxy)
	}
	return &Client{
		client: client,
		config: config,
	}, nil
}
