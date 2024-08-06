package bangumi

// SubscriptionStatus 番剧状态
type SubscriptionStatus string

const (
	// StatusSubscripting 订阅中
	StatusSubscripting SubscriptionStatus = "subscripting"
	// StatusStopped 停止订阅
	StatusStopped SubscriptionStatus = "stopped"
)

// CoverageStrategy 覆盖策略
type CoverageStrategy string

const (
	// CoverageStrategyNone 不覆盖
	CoverageStrategyNone CoverageStrategy = "none"
	// CoverageStrategyHighQuality 高质量覆盖策略，文件大小大的覆盖小的
	CoverageStrategyHighQuality CoverageStrategy = "highQuality"
	// CoverageStrategyLatest 最新覆盖策略，文件时间新的覆盖旧的
	CoverageStrategyLatest CoverageStrategy = "latest"
)

// Bangumi 番剧信息
type Bangumi struct {
	Name       string             // 番剧名称
	RSSLink    string             // 番剧RSS链接
	Status     SubscriptionStatus // 订阅状态
	IncludeReg string             // 包含匹配，正则表达式，作用于RSS标题
	ExcludeReg string             // 排除匹配，正则表达式，作用于RSS标题
	Meta
}

// ParseRSSRsp 解析RSS返回的番剧信息
type ParseRSSRsp struct {
	Name    string
	Season  int
	Year    string
	TMDBID  int
	RSSLink string
}

// FileStatus 文件状态
type FileStatus string

const (
	// FileStatusNotProcessed 未处理
	FileStatusNotProcessed FileStatus = "notProcessed"
	// FileStatusIgnore 忽略
	FileStatusIgnore FileStatus = "ignore"
	// FileStatusDownloading 下载中
	FileStatusDownloading FileStatus = "downloading"
	// FileStatusDownloaded 已下载
	FileStatusDownloaded FileStatus = "downloaded"
	// FileStatusLinked 已链接
	FileStatusLinked FileStatus = "linked"
	// FileStatusLinkedError 链接错误
	FileStatusLinkedError FileStatus = "linkError"
)

// File 番剧文件
type File struct {
	Path         string     // 文件路径
	FileName     string     // 种子文件名
	Status       FileStatus // 文件状态
	StatusDetail string     // 文件状态详情，一般用于存储错误信息
	VideoPath    string     // 链接路径
	RSSTitle     string     // RSS订阅文件中的标题名
	Episode      string     // 集数
}

// RSSItem RSS节点信息
type RSSItem struct {
	GUID        string
	TorrentLink string
}

// RSS RSS信息
type RSS struct {
	BangumiName string
	Season      int
	Items       []RSSItem
}

// Meta 番剧元数据信息
type Meta struct {
	MetaBase
	Season      int
	OriginName  string
	EnglishName string
}

// MetaBase 番剧基础元数据信息
type MetaBase struct {
	ChineseName string
	Year        string
	TMDBID      int
}

// SubscribeReq 订阅请求
type SubscribeReq struct {
	Name       string // 番剧名
	RSSLink    string // RSS链接
	Season     int    // 季数
	IncludeReg string // 包含匹配，正则表达式，作用于RSS标题
	ExcludeReg string // 排除匹配，正则表达式，作用于RSS标题
	TMDBID     int    // TMDB ID
}

// ListBangumiReq 查询番剧请求
type ListBangumiReq struct {
}

// DownloadReq 下载请求
type DownloadReq struct {
	TorrentLink string
	SavePath    string
}
