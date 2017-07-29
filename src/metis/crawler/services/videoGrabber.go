package services

import (
	"flag"
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"
)

//VideoGrabber Responsible for fetching videos from youtube.
type VideoGrabber struct {
}

//Videos Responsible for fetching videos from youtube.
func (videoGrabber VideoGrabber) Videos(query string) []youtube.SearchResult {
	videos := []youtube.SearchResult{}
	flag.Parse()
	client := &http.Client{
		Transport: &transport.APIKey{Key: "AIzaSyBCkyMrYh1uhOkhX8oFcGtJikrjkvuNjkA"},
	}
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}
	call := service.Search.List("id,snippet").
		Q(query).
		MaxResults(20)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}
	// Iterate through each item and add it to the correct list.
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos = append(videos, *item)
		}
	}
	return videos
}
