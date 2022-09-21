package main

import (
	"log"

	xuanim "github.com/ysicing/go-xuanimbot"
)

func main() {
	im, err := xuanim.New("token", "ysicing",
		xuanim.WithBaseURL("http://127.0.0.1"),
		xuanim.WithDevMode(),
		xuanim.WithDumpAll(),
	)
	if err != nil {
		log.Fatal(err)
	}
	if _, _, err = im.Notification.SendUser(xuanim.UserMessage{
		Users: []string{"ysicing"},
		MessageBody: xuanim.MessageBody{
			Title: "66666",
		},
	}); err != nil {
		log.Println(err)
	}
	// if _, _, err = im.Notification.SendChat(xuanim.ChatMessage{
	// 	GID: "8c",
	// 	MessageBody: xuanim.MessageBody{
	// 		Title: "66666",
	// 	},
	// }); err != nil {
	// 	log.Println(err)
	// }
}
