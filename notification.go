// AGPL License
// Copyright (c) 2023 ysicing <i@ysicing.me>

package xuanim

import "github.com/imroc/req/v3"

type NotificationService struct {
	client *Client
}

func (ns *NotificationService) SendUser(user SendUserNotification) (msgResp *MessageResp, resp *req.Response, err error) {
	if ns.client.Custom {
		resp, err = ns.client.client.R().
			SetHeader("token", ns.client.Token).
			SetHeader("caller", ns.client.Caller).
			SetHeader("Content-Type", "application/json").
			SetBody(&user).
			SetResult(&msgResp).
			Post(ns.client.RequestURL("/xuanxuan-notification.json"))
	} else {
		resp, err = ns.client.client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(&user).
			SetResult(msgResp).
			Post(ns.client.CustomRequestURL("/x.php", "sendNotification"))
	}
	return msgResp, resp, err
}

func (ns *NotificationService) SendChat(chat SendGroupNotification) (msgResp *MessageResp, resp *req.Response, err error) {
	if ns.client.Custom {
		resp, err = ns.client.client.R().
			SetHeader("token", ns.client.Token).
			SetHeader("caller", ns.client.Caller).
			SetHeader("Content-Type", "application/json").
			SetBody(&chat).
			SetResult(msgResp).
			Post(ns.client.RequestURL("/xuanxuan-chat.json"))
	} else {
		resp, err = ns.client.client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(&chat).
			SetResult(msgResp).
			Post(ns.client.CustomRequestURL("/x.php", "sendChatMessage"))
	}
	return msgResp, resp, err
}
