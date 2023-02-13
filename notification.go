package xuanim

import "github.com/imroc/req/v3"

type NotificationService struct {
	client *Client
}

type ContentType string

var (
	ContentPlain    ContentType = "plain" // 纯文本
	ContentMarkdown ContentType = "text"  // markdown
)

type MessageBody struct {
	Title       string      `json:"title"`              // 推送标题
	SubTitle    string      `json:"subtitle,omitempty"` // 子标题
	Content     string      `json:"content,omitempty"`  // 推送内容
	ContentType ContentType `json:"contentType"`        // 类型
	URL         string      `json:"url,omitempty"`      // 跳转地址
	Actions     []Actions   `json:"actions,omitempty"`  // 使用 操作对象数组表示该通知支持的操作
	Sender      Sender      `json:"sender,omitempty"`   // 通知的 发送方信息对象
}

type ActionType string

var (
	ActionTypePrimary   ActionType = "primary"
	ActionTypeSuccess   ActionType = "success"
	ActionTypeDanger    ActionType = "danger"
	ActionTypeWarning   ActionType = "warning"
	ActionTypeInfo      ActionType = "info"
	ActionTypeImportant ActionType = "important"
	ActionTypeSpecial   ActionType = "special"
)

type Actions struct {
	Label string     `json:"label"`
	Icon  string     `json:"icon,omitempty"`
	URL   string     `json:"url"`
	Type  ActionType `json:"type,omitempty"`
}

type Sender struct {
	ID     interface{} `json:"id"` // 数字或者字符串
	Name   string      `json:"name,omitempty"`
	Avatar string      `json:"avatar"`
}

type UserMessage struct {
	Users []string `json:"users"` // 推送用户, 使用一个用户 ID 数组指定通知发送给哪些用户，例如[1, 4]，也可以指定一个用户账号组成的数组，例如['admin', 'zhangsan', 'lisi']
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
		SetSuccessResult(&msgResp).
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
		SetSuccessResult(&msgResp).
		Post(ns.client.RequestURL("/xuanxuan-chat.json"))
	return &msgResp, resp, err
}
