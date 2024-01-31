package notion

import (
	"log/slog"

	"github.com/go-resty/resty/v2"
)

type Configuration struct {
	ApiKey     string
	DatabaseId string
}

type Client struct {
	Configuration Configuration
	Logger        *slog.Logger
}

func (c *Client) Inner() *resty.Client {
	inner := resty.New()

	inner.SetBaseURL("https://api.notion.com/v1/")
	inner.SetHeader("Accept", "application/json")
	inner.SetHeader("Content-Type", "application/json")
	inner.SetHeader("Notion-Version", "2022-06-28")
	inner.SetAuthToken(c.Configuration.ApiKey)

	return inner
}
