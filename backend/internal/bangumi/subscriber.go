package bangumi

import (
	"context"
)

//go:generate mockgen -destination subscriber_mock.go -source $GOFILE -package $GOPACKAGE

// NewSubscriber 创建番剧订阅器
func NewSubscriber(dep Dependency) Subscriber {
	return &subscriber{
		rssParser:  dep.RSSParser,
		metaParser: dep.MetaParser,
	}
}

// Dependency subscriber初始化依赖
type Dependency struct {
	RSSParser
	MetaParser
}

// RSSParser RSS解析器
type RSSParser interface {
	Parse(ctx context.Context, link string) (RSS, error)
}

// MetaParser 番剧元数据解析器
type MetaParser interface {
	Search(ctx context.Context, name string) (MetaBase, error)
}

type subscriber struct {
	rssParser  RSSParser
	metaParser MetaParser
}

func (s *subscriber) ParseRSS(ctx context.Context, rssLink string) (ParseRSSRsp, error) {
	rss, err := s.rssParser.Parse(ctx, rssLink)
	if err != nil {
		return ParseRSSRsp{}, err
	}
	meta, err := s.metaParser.Search(ctx, rss.BangumiName)
	if err != nil {
		return ParseRSSRsp{}, err
	}
	return ParseRSSRsp{
		Name:    meta.ChineseName,
		Season:  rss.Season,
		Year:    meta.Year,
		TMDBID:  meta.TMDBID,
		RSSLink: rssLink,
	}, nil
}

func (s *subscriber) Subscribe(ctx context.Context, bangumi Bangumi) error {
	// TODO implement me
	panic("implement me")
}
