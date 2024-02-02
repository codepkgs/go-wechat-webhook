# 说明

> `Go` 版本的企业微信 `Webhook` 机器人 `SDK`

# 功能列表

## 支持的消息类型

[群机器人](https://developer.work.weixin.qq.com/document/path/91770)

- [x] 文本类型 `client.Text`
- [x] Markdown类型 `client.Markdown`
- [x] 图片类型 `client.Image`
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

- 发送Markdown消息

  注意：如果使用多行字符串的话，需要设置 `replaceAllTable` 为 `true`，将所有的 `\n\t` 替换为 `\n`，否则发出来的可能不是Markdown格式的消息。

  ```go
  ret, err := client.Markdown(fmt.Sprintf(`
  # %s
  **%s**
  [这是一个链接](%s)
  > 这是一个引用文本
  * 列表1
  * 列表2
  <font color="red">红色字体</font>`, "一级标题", "加粗", "http://work.weixin.qq.com/api/doc"), true,
  )
  if err != nil {
      fmt.Println(err)
  } else {
      fmt.Printf("%#v", ret)
  }
  ```

- 发送图片类型的消息

  注意：图片类型支持png和jpg格式，不能超过2M。

  ```go
  f, err := os.Open("/Users/hezhang/Desktop/test.png")
  if err != nil {
      fmt.Println(err)
  }
  defer func() { _ = f.Close() }()

  ibs, err := io.ReadAll(f)
  if err != nil {
      fmt.Println(err)
  }

  ret, err := client.Image(ibs)
  if err != nil {
      fmt.Println(err)
  } else {
      fmt.Printf("%#v", ret)
  }
  ```