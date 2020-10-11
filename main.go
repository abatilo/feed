package main

import (
	"log"

	"github.com/munrocape/hn/hnclient"
)

type FeedItem struct {
	description string
	url         string
}

func main() {
	c := hnclient.NewClient()
	storyIDs, _ := c.GetTopStories(20)
	for _, storyID := range storyIDs {
		hnItem, _ := c.GetItem(storyID)
		item := &FeedItem{
			description: hnItem.Title,
			url:         hnItem.Url,
		}
		log.Printf("%#v\n", item)
	}
}
