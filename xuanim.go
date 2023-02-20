// AGPL License
// Copyright (c) 2023 ysicing <i@ysicing.me>

package xuanim

import (
	"fmt"
	"net/url"

	"github.com/imroc/req/v3"
)

const (
	userAgent = "go-xuanim"
)

type Client struct {
	client  *req.Client
	baseURL *url.URL
	Token   string // token
	Caller  string // caller 机器人
	Custom  bool   // 是否定制

	Notification *NotificationService
}

func New(token, caller string, options ...ClientOptionFunc) (*Client, error) {
	client := &Client{}
	client.client = req.C().SetLogger(nil)
	client.setReqUserAgent(userAgent)
	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(client); err != nil {
			return nil, err
		}
	}
	client.Caller = caller
	client.Token = token
	client.Notification = &NotificationService{client: client}
	return client, nil
}

func (c *Client) setBaseURL(urlStr string) error {
	baseURL, err := url.Parse(urlStr)
	if err != nil {
		return err
	}

	// Update the base URL of the client.
	c.baseURL = baseURL

	return nil
}

func (c *Client) setDebug() error {
	c.client.EnableDebugLog()
	return nil
}

func (c *Client) setDumpAll() error {
	c.client.EnableDumpAll()
	return nil
}

func (c *Client) setDisableProxy() error {
	c.client.SetProxy(nil)
	return nil
}

func (c *Client) setReqUserAgent(ua string) error {
	c.client.SetUserAgent(ua)
	return nil
}

func (c *Client) setCustom(custom bool) error {
	c.Custom = custom
	return nil
}

func (c *Client) CustomRequestURL(path, t string) string {
	u := *c.baseURL
	u.Path = c.baseURL.Path + path
	return fmt.Sprintf("%s?m=im&f=%s&code=%s&token=%s", u.String(), t, c.Caller, c.Token)
}

func (c *Client) RequestURL(path string) string {
	u := *c.baseURL
	u.Path = c.baseURL.Path + path
	return u.String()
}
