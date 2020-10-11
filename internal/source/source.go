package source

import (
	"log"

	"github.com/munrocape/hn/hnclient"
)

type FeedItem struct {
	description string
	url         string
}

func HN() {
	c := hnclient.NewClient()
	storyIDs, _ := c.GetTopStories(20)
	for _, storyID := range storyIDs {
		hnItem, _ := c.GetItem(storyID)
		item := &FeedItem{
			description: hnItem.Title,
			url:         hnItem.Url,
		}
		log.Printf("Testing: %#v\n", item)
	}
}
