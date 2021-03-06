package myhttp

import "time"
import "github.com/parnurzeal/gorequest"

type Auth struct {
	Username string
	Password string
}

type Client struct {
	Request    *gorequest.SuperAgent
	Proxy      string
	Auth       Auth
	Timeout    time.Duration
	RetryCount int
}

func NewClient(v ...interface{}) *Client {

	client := Client{}
	for k, v := range v {
		if k == 0 {
			client.Proxy = v.(string)
		}

		if k == 1 {
			client.Auth = v.(Auth)
		}

		if k == 2 {
			client.Timeout = v.(time.Duration)
		}

		if k == 3 {
			client.RetryCount = v.(int)
		}
	}

	request := gorequest.New()
	if client.Auth.Username != "" && client.Auth.Password != "" {
		request.SetBasicAuth(client.Auth.Username, client.Auth.Password)
	}

	if client.Proxy != "" {
		request.Proxy(client.Proxy)
	}

	if client.Timeout > 0 {
		request.Timeout(client.Timeout)
	}

	client.Request = request
	return &client
}

func (c *Client) Get(url string, headers map[string]string) (string, []error) {

	if len(headers) > 0 {
		for k, v := range headers {
			c.Request.Header.Add(k, v)
		}
	}

	_, body, errs := c.Request.Get(url).End()
	if errs != nil {
		return "", errs
	}
	return body, nil
}

func Post() {

}
