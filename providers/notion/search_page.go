package notion

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func (c *Client) SearchPage(term string) string {
	inner := c.Inner()

	body := fmt.Sprintf(`{"filter": {"property": "Name", "title": {"equals": "%s"}}}`, term)

	resp, err := inner.R().SetBody(body).Post("databases/" + c.Configuration.DatabaseId + "/query")

	if err != nil {
		c.Logger.Error(err.Error())
	}

	return gjson.Get(resp.String(), "results.0.id").String()
}
