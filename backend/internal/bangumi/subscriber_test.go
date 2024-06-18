package bangumi

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSubscriber_Parse(t *testing.T) {
	testCases := []struct {
		name    string
		fake    func(t *testing.T) (SubscriberDep, func())
		link    string
		want    ParseRSSRsp
		wantErr bool
	}{
		{
			name: "success",
			fake: func(t *testing.T) (SubscriberDep, func()) {
				ctrl := gomock.NewController(t)
				rpm := NewMockRSSParser(ctrl)
				rpm.EXPECT().Parse(gomock.Any(), gomock.Any()).Return(RSS{
					BangumiName: "bangumi",
					Season:      2,
				}, nil).AnyTimes()
				mpm := NewMockMetaParser(ctrl)
				mpm.EXPECT().Search(gomock.Any(), gomock.Any()).Return(
					MetaBase{
						ChineseName: "chi",
						Year:        "2009",
						TMDBID:      11,
					}, nil).AnyTimes()
				return SubscriberDep{
					RSSParser:  rpm,
					MetaParser: mpm,
				}, ctrl.Finish
			},
			link: "test",
			want: ParseRSSRsp{
				Name:    "chi",
				Season:  2,
				Year:    "2009",
				TMDBID:  11,
				RSSLink: "test",
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dep, clo := tc.fake(t)
			defer clo()
			s := NewSubscriber(dep)

			got, err := s.ParseRSS(context.Background(), tc.link)
			t.Log(err)

			assert.Equal(t, tc.wantErr, err != nil)
			assert.Equal(t, tc.want, got)
		})
	}
}
