package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var apiKey, channelID string

func init() {
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API Key is empty.")
	}

	channelID = os.Getenv("CHANNEL_ID")
	if channelID == "" {
		log.Fatal("Channel ID is empty.")
	}
}

const (
	videosListResultLimit = 50
)

func main() {
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	parts := "snippet,statistics"
	items, err := service.Videos.List(parts).Id(channelID).MaxResults(videosListResultLimit).Do()
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items.Items {
		log.Println("title: ", item.Snippet.Title)
		log.Println("view count: ", item.Statistics.ViewCount)
		// and so on
	}
}
