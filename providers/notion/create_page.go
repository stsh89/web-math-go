package notion

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

func (c *Client) CreatePage(name string) string {
	client := resty.New()
	client.SetBaseURL("https://api.notion.com/v1/")
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Content-Type", "application/json")
	client.SetHeader("Notion-Version", "2022-06-28")
	client.SetAuthToken(c.Configuration.ApiKey)

	body := fmt.Sprintf(`{
	    "parent": { "database_id": "%s" },
	    "properties": {
	        "Name": { "title": [ { "text": { "content": "%s" } } ] },
		    "Short description": { "rich_text": [ { "text": { "content": "" } } ] }
	    }
	}`, c.Configuration.DatabaseId, name)

	c.Logger.Info(body)

	resp, err := client.R().SetBody(body).Post("pages")

	if err != nil {
		c.Logger.Error(err.Error())
	}

	id := gjson.Get(resp.String(), "id").String()

	return id
}
