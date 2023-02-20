// AGPL License
// Copyright (c) 2023 ysicing <i@ysicing.me>

package xuanim

type ContentType string

const (
	ContentTypePlain ContentType = "plain" // 纯文本
	ContentTypeText  ContentType = "text"  // Markdown
)

type ActionType string

const (
	ActionTypePrimary   ActionType = "primary"
	ActionTypeSuccess   ActionType = "success"
	ActionTypeDanger    ActionType = "danger"
	ActionTypeWarning   ActionType = "warning"
	ActionTypeInfo      ActionType = "info"
	ActionTypeImportant ActionType = "important"
	ActionTypeSpecial   ActionType = "special"
)

type Actions struct {
	Label      string     `json:"label"`
	Icon       string     `json:"icon,omitempty"`
	URL        string     `json:"url,omitempty"`
	ActionType ActionType `json:"actionType,omitempty"`
}

type Sender struct {
	ID     any    `json:"id"` // 标识发送方唯一身份的字符串或数字
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

type SendBody struct {
	Title       string      `json:"title"`
	Subtitle    string      `json:"subtitle,omitempty"`
	Content     string      `json:"content,omitempty"`
	ContentType ContentType `json:"contentType"`
	URL         string      `json:"url,omitempty"`
	Actions     []Actions   `json:"actions,omitempty"`
	Sender      Sender      `json:"sender,omitempty"`
}

type SendUserNotification struct {
	SendBody
	Users any `json:"users"` // 使用一个用户 ID 数组指定通知发送给哪些用户 [1,4] 或者 ["ysicing"]
}

type SendGroupNotification struct {
	SendBody
	Gid string `json:"gid"` // 讨论组的全局唯一字符串
}

type MessageResp struct {
	Result  string `json:"result,omitempty"`
	Message string `json:"message,omitempty"`
}
