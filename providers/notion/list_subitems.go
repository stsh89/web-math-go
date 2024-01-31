package notion

import (
	"github.com/tidwall/gjson"
)

type ListResult struct {
	Id         string
	Properties ListResultProperties
}

type ListResultProperties struct {
	Name             string
	ShortDescription string
}

func (c *Client) ListSubitems() []ListResult {
	inner := c.Inner()

	resp, err := inner.R().Post("databases/" + c.Configuration.DatabaseId + "/query")

	if err != nil {
		c.Logger.Error(err.Error())
		return []ListResult{}
	}

	var listResults []ListResult

	results := gjson.Get(resp.String(), "results")

	for _, result := range results.Array() {
		resultString := result.String()

		listResults = append(listResults, ListResult{
			Id: gjson.Get(resultString, "id").String(),
			Properties: ListResultProperties{
				ShortDescription: gjson.Get(resultString, `properties.Short Description.rich_text.0.text.content`).String(),
				Name:             gjson.Get(resultString, `properties.Name.title.0.text.content`).String(),
			},
		})
	}

	return listResults
}
