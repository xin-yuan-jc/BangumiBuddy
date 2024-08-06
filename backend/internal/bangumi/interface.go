package bangumi

import (
	"context"
)

//go:generate mockgen -destination interface_mock.go -source $GOFILE -package $GOPACKAGE

// Subscriber 服务层
type Subscriber interface {
	// ParseRSS 解析RSS链接，获取番剧的基本信息，供用户确认
	ParseRSS(ctx context.Context, rssLink string) (ParseRSSRsp, error)
	// Subscribe 订阅番剧
	Subscribe(ctx context.Context, req SubscribeReq) error
}
