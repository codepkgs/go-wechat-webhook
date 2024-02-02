# 说明

> `Go` 版本的企业微信 `Webhook` 机器人 `SDK`

# 功能列表

## 支持的消息类型

[群机器人](https://developer.work.weixin.qq.com/document/path/91770)

- [x] 文本类型 `client.Text`
- [ ] Markdown类型 `client.Markdown`
- [ ] 图片类型 `client.Image`
- [ ] 图文类型 `client.News`
- [ ] 文件类型 `client.File`

# 示例

- 初始化 Client

  > WebhookAddress 为创建机器人时产生的 Webhook 地址。


  ```go
  client, err := wechat.NewClient("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxxxxxxxxxx")
  if err != nil {
	  fmt.Println(err)
  }
  ```

- 发送文本消息

  > `atMobiles`: at指定的用户。
  > `isAtAll`: at所有用户。
  > 如果 `atMobiles` 和 `isAtAll` 同时指定，则会同时生效。


  ```go
  ret, err := client.Text("hello wechat robot", []string{"18611111111"}, false)
  if err != nil {
      fmt.Println(err)
  } else {
      fmt.Printf("%#v", ret)
  }
  ```