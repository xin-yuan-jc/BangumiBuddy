package mikan

import (
	"context"
	"strings"

	"github.com/mmcdole/gofeed"

	"github.com/MangataL/BangumiBuddy/internal/bangumi"
	"github.com/MangataL/BangumiBuddy/pkg/log"
)

func NewParser() bangumi.RSSParser {
	return &parser{
		fp: gofeed.NewParser(),
	}
}

type parser struct {
	fp *gofeed.Parser
}

func (p *parser) Parse(ctx context.Context, link string) (bangumi.RSS, error) {
	feed, err := p.fp.ParseURLWithContext(link, ctx)
	if err != nil {
		return bangumi.RSS{}, err
	}
	return bangumi.RSS{
		BangumiName: getBangumiName(feed.Title),
		Items:       getItems(ctx, feed.Items),
	}, nil
}

func getBangumiName(title string) string {
	const mikanPrefix = "Mikan Project - "
	if strings.HasPrefix(title, mikanPrefix) {
		return title[len(mikanPrefix):]
	}
	return title
}

func getItems(ctx context.Context, items []*gofeed.Item) []bangumi.RSSItem {
	rssItems := make([]bangumi.RSSItem, 0, len(items))
	for _, item := range items {
		if len(item.Enclosures) == 0 {
			log.Warnf(ctx, "item %s has no enclosures", item.GUID)
			continue
		}
		rssItems = append(rssItems, bangumi.RSSItem{
			GUID:        item.GUID,
			TorrentLink: item.Enclosures[0].URL,
		})
	}
	return rssItems
}
