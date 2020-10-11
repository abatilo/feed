//go:generate mockgen -destination hn_mock.go -package source . HNClient

package source

import (
	"github.com/munrocape/hn/hnclient"
)

// HNClient has the functions used from the munrocape/hn library that we use.
// It's wrapped in an interface like this so that we can mock and test it.
type HNClient interface {
	GetTopStories(int) ([]int, error)
	GetItem(id int) (hnclient.Item, error)
}

// HackerNews is our used Hacker News client
type HackerNews struct {
	client HNClient
}

// NewHN will create a new HackerNews client to use as a source
func NewHN(client HNClient) *HackerNews {
	hn := &HackerNews{client}
	return hn
}

// FeedItems will grab as many stop stories as specified
func (h *HackerNews) FeedItems(count int) []FeedItem {
	feedItems := make([]FeedItem, 0)
	storyIDs, _ := h.client.GetTopStories(count)
	for _, storyID := range storyIDs {
		hnItem, _ := h.client.GetItem(storyID)
		item := FeedItem{
			Description: hnItem.Title,
			URL:         hnItem.Url,
		}
		feedItems = append(feedItems, item)
	}
	return feedItems
}
