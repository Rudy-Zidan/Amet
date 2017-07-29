package services

import (
	"log"
	"metis/crawler/client"
	"metis/crawler/structs"
	"strings"

	"golang.org/x/net/html"

	"github.com/PuerkitoBio/goquery"
)

const (
	CNETURL = "https://www.cnet.com"
)

type CNET struct {
	reviews []structs.Review
	client  client.Client
}

func (cnet CNET) Reviews(query string) []structs.Review {
	query = cnet.optimizeQuery(query)
	log.Println("Search for", query)
	response, _ := cnet.client.Fetch(CNETURL + "/search/?query=" + query)
	root, err := html.Parse(strings.NewReader(response.String()))
	if err != nil {
		log.Fatal(err)
	}
	doc := goquery.NewDocumentFromNode(root)
	doc.Find(".searchItem.product").Each(func(index int, item *goquery.Selection) {
		itemInfo := item.Find(".itemInfo").First()
		itemType := itemInfo.Find(".type").First().Text()
		imageWrapper := item.Find(".imageLinkWrapper").First()
		itemFigure := imageWrapper.Find("figure").First()
		itemSpan := itemFigure.Find("span").First()
		if itemType == "Review" || itemType == "REVIEW" {
			itemLink := itemInfo.Find("a").First()
			itemReview := structs.Review{}
			if value, ok := itemSpan.Find("img").First().Attr("data-original"); ok {
				itemReview.ImageURL = value
			}
			itemReview.Title = itemLink.Find("h3").Text()
			if value, ok := itemLink.Attr("href"); ok {
				itemReview.Link = CNETURL + value
			}
			itemReview.Description = itemInfo.Find(".dek").First().Text()
			if value, ok := itemInfo.Find(".stars-rating").First().Attr("aria-label"); ok {
				itemReview.Rate = value
			}
			itemReview.Author = itemInfo.Find(".assetAuthor > a").First().Text()
			itemReview.Date = itemInfo.Find(".assetTime").First().Text()
			cnet.reviews = append(cnet.reviews, itemReview)
		}
	})
	return cnet.reviews
}

func (cnet *CNET) optimizeQuery(query string) string {
	query = strings.TrimSpace(query)
	query = strings.Replace(query, "\n", "", -1)
	query = strings.Replace(query, " ", "+", -1)
	return strings.Replace(query, "-", "", -1)
}
