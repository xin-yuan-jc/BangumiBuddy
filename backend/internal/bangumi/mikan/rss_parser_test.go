package mikan

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/MangataL/BangumiBuddy/internal/bangumi"
)

func TestParser_Parse(t *testing.T) {
	testCases := []struct {
		name    string
		stub    func(t *testing.T) (url string, clo func())
		wantErr bool
		want    bangumi.RSS
	}{
		{
			name: "success",
			stub: func(t *testing.T) (url string, clo func()) {
				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					_, _ = w.Write(rssTestContent(t))
				}))
				return ts.URL, ts.Close
			},
			wantErr: false,
			want: bangumi.RSS{
				BangumiName: "GIRLS BAND CRY",
				Items: []bangumi.RSSItem{
					{
						GUID:        "[喵萌Production&LoliHouse] GIRLS BAND CRY - 02 [WebRip 1080p HEVC-10bit AAC][简繁日内封字幕]",
						TorrentLink: "https://mikanime.tv/Download/20240414/3a2e456a689ead23ca8f49fdc74ba1872c6f0c12.torrent",
					},
					{
						GUID:        "[喵萌Production&LoliHouse] GIRLS BAND CRY - 01 [WebRip 1080p HEVC-10bit AAC][简繁日内封字幕]",
						TorrentLink: "https://mikanime.tv/Download/20240407/b13d145d95d9acdd5fc50784a6906007b540b468.torrent",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			url, clo := tc.stub(t)
			defer clo()
			p := NewParser()

			rss, err := p.Parse(context.Background(), url)
			t.Log(err)

			assert.Equal(t, tc.wantErr, err != nil)
			assert.Equal(t, tc.want, rss)
		})
	}
}

func rssTestContent(t *testing.T) []byte {
	data, err := os.ReadFile("./rss-test.xml")
	require.Nil(t, err)
	return data
}
