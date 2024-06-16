package mikan

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/mmcdole/gofeed"
	"github.com/nssteinbrenner/anitogo"

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
	items := getItems(ctx, feed.Items)
	season := parseBangumiSeason(ctx, items)
	return bangumi.RSS{
		BangumiName: getBangumiName(feed.Title),
		Season:      season,
		Items:       getItems(ctx, feed.Items),
	}, nil
}

const (
	defaultSeason = 1
)

func parseBangumiSeason(ctx context.Context, items []bangumi.RSSItem) int {
	if len(items) == 0 {
		return defaultSeason
	}
	episode := anitogo.Parse(items[0].GUID, anitogo.DefaultOptions)
	log.Debugf(ctx, "episode: %+v", episode)
	if len(episode.AnimeSeason) == 0 {
		// 默认为第一季
		return defaultSeason
	}
	season, _ := strconv.Atoi(episode.AnimeSeason[0])
	return season
}

var (
	mikanTitleRegex = regexp.MustCompile(`Mikan Project\s*-\s*(.*?)(\s*第.*季)?$`)
)

func getBangumiName(title string) string {
	if matches := mikanTitleRegex.FindStringSubmatch(title); len(matches) > 1 {
		return matches[1]
	}

	// 兜底措施
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
