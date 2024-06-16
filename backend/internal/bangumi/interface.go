package bangumi

import (
	"context"
)

// Subscriber 番剧订阅器
type Subscriber interface {
	// ParseRSS 解析RSS链接，获取番剧的基本信息，供用户确认
	ParseRSS(ctx context.Context, rssLink string) (ParseRSSRsp, error)
	// Subscribe 订阅番剧
	Subscribe(ctx context.Context, bangumi Bangumi) error
}
