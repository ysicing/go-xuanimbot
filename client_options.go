package xuanim

type ClientOptionFunc func(*Client) error

// WithBaseURL sets the base URL for API requests to a custom endpoint.
func WithBaseURL(urlStr string) ClientOptionFunc {
	return func(c *Client) error {
		return c.setBaseURL(urlStr)
	}
}

func WithDevMode() ClientOptionFunc {
	return func(c *Client) error {
		return c.setDebug()
	}
}

func WithDumpAll() ClientOptionFunc {
	return func(c *Client) error {
		return c.setDumpAll()
	}
}

// WithoutProxy 禁用代理, 默认情况下会读取HTTP_PROXY/HTTPS_PROXY/http_proxy/https_proxy变量
func WithoutProxy() ClientOptionFunc {
	return func(c *Client) error {
		return c.setDisableProxy()
	}
}

func WithUserAgent(ua string) ClientOptionFunc {
	return func(c *Client) error {
		if ua == "" {
			ua = userAgent
		}
		return c.setReqUserAgent(ua)
	}
}
