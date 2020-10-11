package source_test

import (
	"testing"

	"github.com/abatilo/feed/internal/source"
	gomock "github.com/golang/mock/gomock"
	"github.com/munrocape/hn/hnclient"
	"github.com/stretchr/testify/assert"
)

func Test_HNClient_FeedItems(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockHNItems := []hnclient.Item{
		{
			Title: "title 1",
			Url:   "url 1",
		},
		{
			Title: "title 2",
			Url:   "url 2",
		},
		{
			Title: "title 3",
			Url:   "url 3",
		},
	}

	mockClient := source.NewMockHNClient(ctrl)
	mockClient.EXPECT().GetTopStories(gomock.Any()).Return([]int{1, 2, 3}, nil)
	mockClient.EXPECT().GetItem(gomock.Eq(1)).Return(mockHNItems[0], nil)
	mockClient.EXPECT().GetItem(gomock.Eq(2)).Return(mockHNItems[1], nil)
	mockClient.EXPECT().GetItem(gomock.Eq(3)).Return(mockHNItems[2], nil)

	hnClient := source.NewHN(mockClient)
	feedItems := hnClient.FeedItems(3)

	expectedItems := []source.FeedItem{
		{
			Description: "title 1",
			URL:         "url 1",
		},
		{
			Description: "title 2",
			URL:         "url 2",
		},
		{
			Description: "title 3",
			URL:         "url 3",
		},
	}
	assert.Equal(expectedItems, feedItems)
}
