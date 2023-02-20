// AGPL License
// Copyright (c) 2023 ysicing <i@ysicing.me>

package main

import xuanim "github.com/ysicing/go-xuanimbot"

func main() {
	customClient, _ := xuanim.New("123", "xxx", xuanim.WithBaseURL("https://demo.xuanim"), xuanim.WithCustom(true))
	SendBody := xuanim.SendBody{
		Title:       "",
		Subtitle:    "",
		Content:     "",
		ContentType: "",
		URL:         "",
		Actions:     nil,
		Sender:      xuanim.Sender{},
	}
	customClient.Notification.SendUser(xuanim.SendUserNotification{
		Users: []string{"ysicing"},
		SendBody,
	})
}
