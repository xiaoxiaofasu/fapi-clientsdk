package client

type Client struct {
	AccessKey   string
	SecretKey   string
	GatewayHost string
	UseGateway  bool
}

type Option func(*Client)

func SetAkSk(ak, sk string) Option {
	return func(c *Client) {
		c.AccessKey = ak
		c.SecretKey = sk
	}
}

func UseGateway(url string) Option {
	return func(c *Client) {
		c.GatewayHost = url
		c.UseGateway = true
	}
}

func NewClient(opts ...Option) *Client {
	client := &Client{}
	for _, opt := range opts {
		opt(client)
	}
	return client
}
