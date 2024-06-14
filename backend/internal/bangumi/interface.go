package bangumi

import (
	"context"
)

// Subscriber 番剧订阅器
type Subscriber interface {
	// Parse 解析RSS链接，获取番剧的基本信息
	Parse(ctx context.Context, rssLink string) (Bangumi, error)
	// Subscribe 订阅番剧
	Subscribe(ctx context.Context, bangumi Bangumi) error
}
