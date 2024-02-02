package wechat

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

const (
	RequestTimeout = 5000 //请求超时时间
)

// 发送HTTP POST 请求
func (c *Client) do(url string, bytes []byte) ([]byte, error) {
	r := resty.New().
		SetTimeout(time.Duration(RequestTimeout)*time.Millisecond).
		R().
		SetHeader("Content-Type", "application/json")

	r.SetBody(bytes)
	fmt.Println(string(bytes))
	resp, err := r.Post(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", resp.StatusCode(), resp.Body())
	}

	return resp.Body(), nil
}
