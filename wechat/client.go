package wechat

import (
	"fmt"
	"strings"
)

type Client struct {
	WebhookAddress string
}

var ErrWebhookAddress = fmt.Errorf(`wechat webhook address must begin with "http://" or "https://"`)

func NewClient(webhookAddress, secret string) (*Client, error) {
	if !strings.HasPrefix(webhookAddress, "http://") && !strings.HasPrefix(webhookAddress, "https://") {
		return nil, ErrWebhookAddress
	}

	return &Client{
		WebhookAddress: webhookAddress,
	}, nil
}
