package xuanim

import "github.com/imroc/req/v3"

type NotificationService struct {
	client *Client
}

type MessageBody struct {
	Title    string `json:"title"`              // 推送标题
	SubTitle string `json:"subtitle,omitempty"` // 子标题
	Content  string `json:"content,omitempty"`  // 推送内容
	URL      string `json:"url,omitempty"`      // 跳转地址
}

type UserMessage struct {
	Users []string `json:"users"` // 推送用户
	MessageBody
}

type ChatMessage struct {
	GID string `json:"gid"` // 推送群主ID
	MessageBody
}

type MessageResp struct {
	Result  string `json:"result,omitempty"`
	Message string `json:"message,omitempty"`
}

func (ns *NotificationService) SendUser(user UserMessage) (*MessageResp, *req.Response, error) {
	var msgResp MessageResp
	resp, err := ns.client.client.R().
		SetHeader("token", ns.client.Token).
		SetHeader("caller", ns.client.Caller).
		SetHeader("Content-Type", "application/json").
		SetBody(&user).
		SetResult(&msgResp).
		Post(ns.client.RequestURL("/xuanxuan-notification.json"))
	return &msgResp, resp, err
}

func (ns *NotificationService) SendChat(chat ChatMessage) (*MessageResp, *req.Response, error) {
	var msgResp MessageResp
	resp, err := ns.client.client.R().
		SetHeader("token", ns.client.Token).
		SetHeader("caller", ns.client.Caller).
		SetHeader("Content-Type", "application/json").
		SetBody(&chat).
		SetResult(&msgResp).
		Post(ns.client.RequestURL("/xuanxuan-chat.json"))
	return &msgResp, resp, err
}
