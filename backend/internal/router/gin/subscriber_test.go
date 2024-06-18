package gin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/MangataL/BangumiBuddy/internal/bangumi"
)

func TestSubscriber_ParseRSS(t *testing.T) {
	initArgs := func(link string) args {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request, _ = http.NewRequest("GET", "/bangumi/rss?link="+link, nil)
		return args{
			ctx:    c,
			writer: w,
		}
	}

	testCases := []struct {
		name       string
		args       args
		fake       func(t *testing.T) (bangumi.Subscriber, func())
		wantStatus int
		wantBody   string
	}{
		{
			name: "success",
			args: initArgs("link"),
			fake: func(t *testing.T) (bangumi.Subscriber, func()) {
				ctrl := gomock.NewController(t)
				sm := bangumi.NewMockSubscriber(ctrl)
				sm.EXPECT().ParseRSS(gomock.Any(), "link").Return(bangumi.ParseRSSRsp{
					Name:    "name",
					Season:  2,
					Year:    "2024",
					TMDBID:  1,
					RSSLink: "link",
				}, nil).AnyTimes()
				return sm, ctrl.Finish
			},
			wantStatus: http.StatusOK,
			wantBody:   `{"name":"name","season":2,"year":"2024","tmdb_id":1,"rss_link":"link"}`,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dep, clo := tc.fake(t)
			defer clo()
			a := NewSubscriber(dep)

			a.ParseRSS(tc.args.ctx)

			assert.Equal(t, tc.wantStatus, tc.args.writer.Code)
			assert.Equal(t, tc.wantBody, tc.args.writer.Body.String())
		})
	}
}
