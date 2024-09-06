package source

import (
	"context"
	"news-feed-bot/internal/model"
	"github.com/SlyMarbo/rss"
)

type RSSSource struct {
	URL        string
	SourceID   int64
	SourceName string
}

func NewRSSSourceFromModel(m model.Source) RSSSource {
	return RSSSource{
		URL:        m.FeedURL,
		SourceID:   m.ID,
		SourceName: m.Name,
	}
}
func (s RSSSource) loadFeed(ctx context.Context, url string) (*rss.Feed, error) {
	var (
		feedCh = make(chan *rss.Feed)
		errCh  = make(chan error)
	)
	go func() {
		feed, err := rss.Fetch(url
		)
	}()
}
