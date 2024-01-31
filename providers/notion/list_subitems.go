package notion

import (
	"github.com/go-resty/resty/v2"
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
	client := resty.New()
	client.SetBaseURL("https://api.notion.com/v1/databases/")
	client.SetHeader("Accept", "application/json")
	client.SetHeader("Notion-Version", "2022-06-28")
	client.SetAuthToken(c.Configuration.ApiKey)

	resp, err := client.R().Post(c.Configuration.DatabaseId + "/query")

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
