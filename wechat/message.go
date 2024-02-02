package wechat

import (
	"encoding/json"
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
