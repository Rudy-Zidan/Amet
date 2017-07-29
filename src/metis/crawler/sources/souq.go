package sources

import (
	"metis/crawler/structs"

	"github.com/PuerkitoBio/goquery"
)

// Souq Represent Souq.com Website
type Souq struct {
	document *goquery.Document
	content  structs.Content
}

// Content responsible for get product content.
func (souq Souq) Content(document *goquery.Document) structs.Content {
	souq.document = document
	souq.content.Title = souq.title()
	souq.content.Description = souq.description()
	souq.content.Details = souq.specs()
	return souq.content
}

// title responsible for get product title.
func (souq Souq) title() string {
	return souq.document.Find(".product-title > h1").Text()
}

// description responsible for get product description.
func (souq Souq) description() string {
	return souq.document.Find("#description-full > div > div > div > p").Text()
}

// specs responsible for get product specs.
func (souq Souq) specs() []structs.Meta {
	metaTags := []structs.Meta{}
	meta := structs.Meta{}
	souq.document.Find("dl.stats").Children().Each(func(index int, tag *goquery.Selection) {
		if tag.Nodes[0].Data == "dt" {
			meta.Key = tag.Text()
		} else if tag.Nodes[0].Data == "dd" {
			meta.Value = tag.Text()
			metaTags = append(metaTags, meta)
			meta = structs.Meta{}
		}
	})
	return metaTags
}
