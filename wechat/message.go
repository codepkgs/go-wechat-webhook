package wechat

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// RobotReturn 微信机器人返回的结果
type RobotReturn struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// ReturnResult 返回结果
func ReturnResult(bytes []byte) (*RobotReturn, error) {
	var r RobotReturn
	err := json.Unmarshal(bytes, &r)
	if err != nil {
		return nil, err
	} else {
		return &r, nil
	}
}

// Text 发送普通文本消息
func (c *Client) Text(content string, atMobiles []string, isAtAll bool) (*RobotReturn, error) {
	if isAtAll {
		atMobiles = append(atMobiles, "@all")
	}

	t := struct {
		Msgtype string `json:"msgtype"`
		Text    struct {
			Content             string   `json:"content"`
			MentionedMobileList []string `json:"mentioned_mobile_list"`
		} `json:"text"`
	}{
		Msgtype: "text",
		Text: struct {
			Content             string   `json:"content"`
			MentionedMobileList []string `json:"mentioned_mobile_list"`
		}{Content: content, MentionedMobileList: atMobiles},
	}

	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(c.WebhookAddress, body)
	if err != nil {
		return nil, err
	}

	return ReturnResult(resp)
}

// Markdown Markdown类型
func (c *Client) Markdown(content string, replaceAllTable bool) (*RobotReturn, error) {
	if replaceAllTable {
		content = strings.ReplaceAll(content, "\n\t", "\n")
	}

	t := struct {
		Msgtype  string `json:"msgtype"`
		Markdown struct {
			Content string `json:"content"`
		} `json:"markdown"`
	}{
		Msgtype: "markdown",
		Markdown: struct {
			Content string `json:"content"`
		}{Content: content},
	}

	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(c.WebhookAddress, body)
	if err != nil {
		return nil, err
	}

	return ReturnResult(resp)
}

// Image 图片类型的消息
func (c *Client) Image(image []byte) (*RobotReturn, error) {
	// image md5
	m := md5.New()
	if _, err := io.Copy(m, bytes.NewReader(image)); err != nil {
		return nil, err
	}
	mhash := fmt.Sprintf("%x", m.Sum(nil))

	// image hash
	b64 := make([]byte, base64.StdEncoding.EncodedLen(len(image)))
	base64.StdEncoding.Encode(b64, image)

	t := struct {
		Msgtype string `json:"msgtype"`
		Image   struct {
			Base64 string `json:"base64"`
			Md5    string `json:"md5"`
		} `json:"image"`
	}{
		Msgtype: "image",
		Image: struct {
			Base64 string `json:"base64"`
			Md5    string `json:"md5"`
		}{Base64: string(b64), Md5: mhash},
	}

	body, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(c.WebhookAddress, body)
	if err != nil {
		return nil, err
	}

	return ReturnResult(resp)
}
