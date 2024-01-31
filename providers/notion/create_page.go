package notion

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func (c *Client) CreatePage(name string) string {
	inner := c.Inner()

	body := fmt.Sprintf(`{
	    "parent": { "database_id": "%s" },
	    "properties": {
	        "Name": { "title": [ { "text": { "content": "%s" } } ] },
		    "Short description": { "rich_text": [ { "text": { "content": "" } } ] }
	    }
	}`, c.Configuration.DatabaseId, name)

	resp, err := inner.R().SetBody(body).Post("pages")

	if err != nil {
		c.Logger.Error(err.Error())
	}

	id := gjson.Get(resp.String(), "id").String()

	return id
}
