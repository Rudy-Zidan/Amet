package crawler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"metis/crawler/parser"
	"metis/crawler/services"
	"metis/crawler/structs"

	youtube "google.golang.org/api/youtube/v3"
)

// Response represent Crawler output
type Response struct {
	SiteName string
	Meta     []structs.Meta
	Content  structs.Content
	Videos   []youtube.SearchResult
	Reviews  []structs.Review
	Error    error
}

var (
	response Response
	//DOM represent DOM parser of html
	DOM          parser.DOM
	videoGrabber services.VideoGrabber
	cnet         services.CNET
)

// Retrive Link information
func Retrive(link string) []byte {
	if state, err := DOM.Load(link); state {
		response.SiteName = DOM.ExtractSiteName(link)
		response.Meta = DOM.ExtractMetaData()
		response.Content = DOM.ExtractContent()
		response.Videos = videoGrabber.Videos(response.Content.Title)
		response.Reviews = cnet.Reviews(response.Content.Title)
		response.Error = nil
	} else {
		fmt.Println(err.Error())
		response.Error = err
	}
	err := ioutil.WriteFile("result.json", inJSON(&response), 0644)
	if err != nil {
		fmt.Println("failed")
	}
	log.Println("===============================================================")
	return inJSON(&response)
}

func inJSON(response *Response) []byte {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error")
	}
	return jsonResponse
}
