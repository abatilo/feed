//go:generate mockgen -destination hn_mock.go -package source . HNClient

package source

import (
	"github.com/munrocape/hn/hnclient"
)

type HNClient interface {
	GetTopStories(int) ([]int, error)
	GetItem(id int) (hnclient.Item, error)
}

type HackerNews struct {
	client HNClient
}

func NewHN(client HNClient) *HackerNews {
	hn := &HackerNews{client}
	return hn
}

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
