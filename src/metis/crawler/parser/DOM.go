package parser

import (
	"metis/crawler/client"
	"metis/crawler/sources"
	"metis/crawler/structs"
	"net/url"
	"strings"

	"golang.org/x/net/html"

	"github.com/PuerkitoBio/goquery"
)

var (
	souq   sources.Souq
	jumia  sources.Jumia
	amazon sources.Amazon
)

//DOM Represent HTML DOM parser
type DOM struct {
	siteName string
	// document DOM
	document *goquery.Document
	client   client.Client
}

//Load html
func (dom *DOM) Load(link string) (bool, error) {
	response, err := dom.client.Fetch(link)
	if err != nil {
		return false, err
	}
	root, err := html.Parse(strings.NewReader(response.String()))
	if err != nil {
		return false, err
	}
	dom.document = goquery.NewDocumentFromNode(root)
	return true, nil
}

// ExtractMetaData Should get all valid meta tags in the site.
func (dom *DOM) ExtractMetaData() []structs.Meta {
	metaTags := []structs.Meta{}
	dom.document.Find("meta").Each(func(index int, meta *goquery.Selection) {
		tag := structs.Meta{}
		if value, ok := meta.Attr("property"); ok {
			tag.Key = value
		}
		if value, ok := meta.Attr("name"); ok {
			tag.Key = value
		}
		if value, ok := meta.Attr("content"); ok {
			tag.Value = value
		}
		if tag.Value != "" {
			metaTags = append(metaTags, tag)
		}
	})
	return metaTags
}

// ExtractSiteName Should return site name in string.
func (dom *DOM) ExtractSiteName(link string) string {
	url, _ := url.Parse(link)
	dom.siteName = strings.Split(url.Host, ".")[1]
	return dom.siteName
}

// ExtractContent Should extract all content possible from site.
func (dom *DOM) ExtractContent() structs.Content {
	siteContent := structs.Content{}
	switch dom.siteName {
	case "souq":
		siteContent = souq.Content(dom.document)
	case "jumia":
		siteContent = jumia.Content(dom.document)
	case "amazon":
		siteContent = amazon.Content(dom.document)
	}
	return siteContent
}
