package tmdb

import (
	"context"

	"github.com/cyruzin/golang-tmdb"

	"github.com/MangataL/BangumiBuddy/internal/bangumi"
	"github.com/MangataL/BangumiBuddy/pkg/errs"
	"github.com/MangataL/BangumiBuddy/pkg/log"
)

func NewParser(c *tmdb.Client) bangumi.MetaParser {
	return &client{
		client: c,
	}
}

type client struct {
	client *tmdb.Client
}

func (t *client) Search(ctx context.Context, name string) (bangumi.MetaBase, error) {
	tvs, err := t.client.GetSearchTVShow(name, map[string]string{
		"language": "zh",
		"page":     "1",
	})
	if err != nil {
		return bangumi.MetaBase{}, err
	}
	log.Debugf(ctx, "search %s got tvs: %+v", name, tvs)
	if len(tvs.Results) == 0 {
		return bangumi.MetaBase{}, errs.NewNotFound("未搜索到番剧")
	}
	tv := tvs.Results[0]
	return bangumi.MetaBase{
		ChineseName: tv.Name,
		Year:        getYear(ctx, tv.FirstAirDate),
		TMDBID:      int(tv.ID),
	}, nil
}

func getYear(ctx context.Context, date string) string {
	if len(date) < 4 {
		log.Warnf(ctx, "invalid date: %s", date)
		return ""
	}
	return date[:4]
}
