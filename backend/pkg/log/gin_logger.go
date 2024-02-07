package log

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// GinLogger returns a gin.HandlerFunc (middleware) that logs requests using Logger
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		fields := []zapcore.Field{
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.Duration("latency", time.Since(start)),
		}
		ctx := WithFields(c.Request.Context(), fields...)
		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				Error(ctx, e)
			}
		}
		Info(ctx, path)
	}
}

// GinRecovery returns a middleware for a given logger that recovers from any panics and logs the errors.
func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se := (&os.SyscallError{}); errors.As(ne.Err, &se) {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") ||
							strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					ctx := WithFields(c.Request.Context(),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					Error(ctx, c.Request.URL.Path)
					// If the connection is dead, we can't write a status to it.
					_ = c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				ctx := WithFields(c.Request.Context(),
					zap.Time("time", time.Now()),
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.String("stack", string(debug.Stack())),
				)
				Error(ctx, "[Recovery from panic]")
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
