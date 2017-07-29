package sources

import (
	"metis/crawler/structs"

	"github.com/PuerkitoBio/goquery"
)

// Amazon Represent amazon.com Website
type Amazon struct {
	document *goquery.Document
	content  structs.Content
}

// Content responsible for get product content.
func (amazon Amazon) Content(document *goquery.Document) structs.Content {
	amazon.document = document
	amazon.content.Title = amazon.title()
	amazon.content.Description = amazon.description()
	amazon.content.Details = amazon.specs()
	return amazon.content
}

// title responsible for get product title.
func (amazon Amazon) title() string {
	return amazon.document.Find("#productTitle").Text()
}

// description responsible for get product description.
func (amazon Amazon) description() string {
	return amazon.document.Find("#productDescription > p").First().Text()
}

// specs responsible for get product specs.
func (amazon Amazon) specs() []structs.Meta {
	metaTags := []structs.Meta{}
	detailsTable := amazon.document.Find("table.prodDetTable").First()
	detailsTable.Find("tr").Each(func(index int, tag *goquery.Selection) {
		meta := structs.Meta{}
		meta.Key = tag.Find("th").First().Text()
		valueTD := tag.Find("td").First()
		valueTD.Find("script").ReplaceWith("")
		valueTD.Find("style").ReplaceWith("")
		valueTD.Find("a").ReplaceWith("")
		meta.Value = valueTD.Text()
		metaTags = append(metaTags, meta)
	})
	return metaTags
}
