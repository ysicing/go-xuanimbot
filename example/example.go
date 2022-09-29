package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	xuanim "github.com/ysicing/go-xuanimbot"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	app := cli.NewApp()
	app.Name = "小喧喧大助手"
	app.Usage = "小喧喧大助手"
	app.Action = run
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug",
			EnvVar: "XIM_DEBUG",
		},
		cli.BoolFlag{
			Name:   "all",
			Usage:  "group all",
			EnvVar: "XIM_GROUPALL",
		},
		cli.StringFlag{
			Name:   "token",
			Usage:  "token",
			EnvVar: "XIM_TOKEN",
		},
		cli.StringFlag{
			Name:   "caller",
			Usage:  "caller",
			EnvVar: "XIM_CALLER",
		},
		cli.StringFlag{
			Name:   "api",
			Usage:  "api",
			EnvVar: "XIM_API",
		},
		cli.StringSliceFlag{
			Name:   "users",
			Usage:  "users",
			EnvVar: "XIM_USERS",
		},
		cli.StringFlag{
			Name:   "group",
			Usage:  "group id",
			EnvVar: "XIM_GID",
		},
		cli.StringFlag{
			Name:   "title",
			Usage:  "title",
			EnvVar: "XIM_TITLE",
		},
		cli.StringFlag{
			Name:   "message",
			Usage:  "message",
			EnvVar: "XIM_MESSAGE",
		},
		cli.StringFlag{
			Name:   "url",
			Usage:  "url",
			EnvVar: "XIM_URL",
		},
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	var op []xuanim.ClientOptionFunc
	op = append(op, xuanim.WithBaseURL(c.String("api")))
	if c.Bool("debug") {
		op = append(op, xuanim.WithDevMode(), xuanim.WithDumpAll())
	}
	im, err := xuanim.New(c.String("token"), c.String("caller"), op...)
	if err != nil {
		return err
	}
	users := c.StringSlice("users")
	gid := c.String("group")
	title := c.String("title")
	if len(title) == 0 {
		title = "收到一条消息"
	}
	msg := c.String("message")
	if len(msg) == 0 {
		msg = "测试嘿嘿"
	}
	if len(users) == 0 && len(gid) == 0 {
		return fmt.Errorf("群组或者个人")
	}
	if len(users) > 0 {
		if _, _, err = im.Notification.SendUser(xuanim.UserMessage{
			Users: users,
			MessageBody: xuanim.MessageBody{
				Title:   title,
				Content: msg,
				URL:     c.String("url"),
			},
		}); err != nil {
			return err
		}
	}
	if c.Bool("all") {
		msg = fmt.Sprintf("%s @all", msg)
	}
	if _, _, err = im.Notification.SendChat(xuanim.ChatMessage{
		GID: gid,
		MessageBody: xuanim.MessageBody{
			Title:   title,
			Content: msg,
			URL:     c.String("url"),
		},
	}); err != nil {
		return err
	}
	return nil
}
