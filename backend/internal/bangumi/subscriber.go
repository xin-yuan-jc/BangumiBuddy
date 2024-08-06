package bangumi

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/MangataL/BangumiBuddy/pkg/log"
)

//go:generate mockgen -destination subscriber_mock.go -source $GOFILE -package $GOPACKAGE

// MustNewSubscriber 创建订阅器
func MustNewSubscriber(dep SubscriberDep) Subscriber {
	ctx, cancel := context.WithCancel(context.Background())
	s := &subscriber{
		rssParser:  dep.RSSParser,
		metaParser: dep.MetaParser,
		repo:       dep.Repository,
		stop:       cancel,
	}
	go s.runDownloading(ctx)
	go s.runLinking(ctx)
	return s
}

// SubscriberDep subscriber初始化依赖
type SubscriberDep struct {
	RSSParser
	MetaParser
	Repository
}

// RSSParser RSS解析器
type RSSParser interface {
	Parse(ctx context.Context, link string) (RSS, error)
}

// MetaParser 番剧元数据解析器
type MetaParser interface {
	Search(ctx context.Context, name string) (MetaBase, error)
	Parse(ctx context.Context, id int) (Meta, error)
}

// Repository 存储层
type Repository interface {
	Save(ctx context.Context, bangumi Bangumi) error
	List(ctx context.Context, req ListBangumiReq) ([]Bangumi, int, error)
	ListFiles(ctx context.Context, bangumiName string) ([]File, error)
	SaveFile(ctx context.Context, file File) error
}

// Downloader 下载器
type Downloader interface {
	Download(ctx context.Context, req DownloadReq) error
	Downloaded() chan<- Torrent
}

// Config 配置项
type Config interface {
	RSSCheckInterval() time.Duration
	SavePath() string
}

type subscriber struct {
	rssParser  RSSParser
	metaParser MetaParser
	repo       Repository
	config     Config
	downloader Downloader

	stop func()

	rssTicker *time.Ticker
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

func (s *subscriber) Subscribe(ctx context.Context, req SubscribeReq) error {
	meta, err := s.metaParser.Parse(ctx, req.TMDBID)
	if err != nil {
		return err
	}
	meta.Season = req.Season
	bangumi := Bangumi{
		Name:       req.Name,
		RSSLink:    req.RSSLink,
		Status:     StatusSubscripting,
		IncludeReg: req.IncludeReg,
		ExcludeReg: req.ExcludeReg,
		Meta:       meta,
	}
	return errors.WithMessage(s.repo.Save(ctx, bangumi), "保存失败")
}

func (s *subscriber) runDownloading(ctx context.Context) {
	ticker := time.NewTicker(s.config.RSSCheckInterval())
	s.rssTicker = ticker
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			return
		case <-ticker.C:
			s.downloadBangumi(ctx)
		}
	}
}

func (s *subscriber) downloadBangumi(ctx context.Context) {
	bangumis, _, err := s.repo.List(ctx, ListBangumiReq{})
	if err != nil {
		log.Errorf(ctx, "获取番剧列表失败：%s", err)
		return
	}
	for _, bangumi := range bangumis {
		rss, err := s.rssParser.Parse(ctx, bangumi.RSSLink)
		if err != nil {
			log.Errorf(ctx, "解析RSS失败：%s", err)
			continue
		}
		files, err := s.repo.ListFiles(ctx, bangumi.Name)
		if err != nil {
			log.Errorf(ctx, "获取番剧文件列表失败：%s", err)
			continue
		}
		for _, item := range rss.Items {
			if fileProcessed(files, item) {
				continue
			}
			if ignore(item, bangumi) {
				if err := s.repo.SaveFile(ctx, File{
					RSSTitle: item.GUID,
					Status:   FileStatusIgnore,
				}); err != nil {
					log.Errorf(ctx, "保存文件忽略状态失败：%s", err)
				}
				continue
			}
			path := s.config.SavePath() + "/" + bangumi.Name
			if err := s.downloader.Download(ctx, DownloadReq{
				TorrentLink: item.TorrentLink,
				SavePath:    path,
			}); err != nil {
				log.Errorf(ctx, "下载文件失败：%s", err)
				continue
			}
			if err := s.repo.SaveFile(ctx, File{
				Path:     path,
				FileName: "",
				Status:   FileStatusDownloading,
				RSSTitle: item.GUID,
				Episode:  "",
			}); err != nil {
				log.Errorf(ctx, "保存文件下载状态失败：%s", err)
				continue
			}
		}
	}
}

func fileProcessed(files []File, item RSSItem) bool {
	for _, file := range files {
		if file.RSSTitle == item.GUID {
			if file.Status != FileStatusNotProcessed {
				return true
			}
		}
	}
	return false
}

func ignore(item RSSItem, bangumi Bangumi) bool {

}

func (s *subscriber) runLinking(ctx context.Context) {

}
