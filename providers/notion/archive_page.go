package notion

import (
	"github.com/tidwall/gjson"
)

func (c *Client) ArchivePage(id string) string {
	inner := c.Inner()

	body := `{"archived": true}`

	resp, err := inner.R().SetBody(body).Patch("pages/" + id)

	if err != nil {
		c.Logger.Error(err.Error())
	}

	return gjson.Get(resp.String(), "id").String()
}
