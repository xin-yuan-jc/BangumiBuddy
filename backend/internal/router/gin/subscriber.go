package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MangataL/BangumiBuddy/internal/bangumi"
)

// Subscriber is the gin handler for bangumi subscription
type Subscriber struct {
	subscriber bangumi.Subscriber
}

// NewSubscriber creates a new Subscriber
func NewSubscriber(subscriber bangumi.Subscriber) *Subscriber {
	return &Subscriber{
		subscriber: subscriber,
	}
}

// ParseRSS is the handler for parsing RSS link
func (s *Subscriber) ParseRSS(c *gin.Context) {
	link := c.Query("link")
	rsp, err := s.subscriber.ParseRSS(c.Request.Context(), link)
	if err != nil {
		writeError(c, err)
		return
	}
	type data struct {
		Name    string `json:"name"`
		Season  int    `json:"season"`
		Year    string `json:"year"`
		TMDBID  int    `json:"tmdb_id"`
		RSSLink string `json:"rss_link"`
	}
	c.JSON(http.StatusOK, data{
		Name:    rsp.Name,
		Season:  rsp.Season,
		Year:    rsp.Year,
		TMDBID:  rsp.TMDBID,
		RSSLink: rsp.RSSLink,
	})
}
