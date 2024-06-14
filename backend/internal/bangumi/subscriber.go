package bangumi

import (
	"context"
)

// NewSubscriber 创建番剧订阅器
func NewSubscriber(dep Dependency) Subscriber {
	return &subscriber{
		rssParser: dep.RSSParser,
	}
}

// Dependency subscriber初始化依赖
type Dependency struct {
	RSSParser
}

// RSSParser RSS解析器
type RSSParser interface {
	Parse(ctx context.Context, link string) (RSS, error)
}

type subscriber struct {
	rssParser RSSParser
}

func (s *subscriber) Parse(ctx context.Context, rssLink string) (Bangumi, error) {
	// TODO implement me
	panic("implement me")
}

func (s *subscriber) Subscribe(ctx context.Context, bangumi Bangumi) error {
	// TODO implement me
	panic("implement me")
}
