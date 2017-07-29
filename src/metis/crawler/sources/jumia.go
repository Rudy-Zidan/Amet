package sources

import (
	"metis/crawler/structs"

	"github.com/PuerkitoBio/goquery"
)

// Jumia Represent jumia.com Website
type Jumia struct {
	document *goquery.Document
	content  structs.Content
}

// Content responsible for get product content.
func (jumia Jumia) Content(document *goquery.Document) structs.Content {
	jumia.document = document
	jumia.content.Title = jumia.title()
	jumia.content.Description = jumia.description()
	jumia.content.Details = jumia.specs()
	return jumia.content
}

// title responsible for get product title.
func (jumia Jumia) title() string {
	return jumia.document.Find("h1.title").Text()
}

// description responsible for get product description.
func (jumia Jumia) description() string {
	return jumia.document.Find(".product-description > p").First().Text()
}

// specs responsible for get product specs.
func (jumia Jumia) specs() []structs.Meta {
	metaTags := []structs.Meta{}
	meta := structs.Meta{}
	first := true
	jumia.document.Find(".osh-row").Children().Each(func(index int, tag *goquery.Selection) {
		if first {
			meta.Key = tag.Text()
			first = false
		} else {
			meta.Value = tag.Text()
			first = true
			metaTags = append(metaTags, meta)
			meta = structs.Meta{}
		}
	})
	return metaTags
}
