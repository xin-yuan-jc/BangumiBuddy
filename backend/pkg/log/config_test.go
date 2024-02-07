package log

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestConfig_Build(t *testing.T) {
	temp, err := os.CreateTemp("", "test")
	require.NoError(t, err, "Failed to create temp file")
	config := Config{
		Level: zapcore.DebugLevel,
		Caller: struct {
			Enable bool `yaml:"enable"`
			Skip   int  `yaml:"skip" json:"skip"`
		}{
			Enable: true,
			Skip:   1,
		},
		Filename: temp.Name(),
	}
	defer func() {
		_ = os.Remove(temp.Name())
	}()

	log, err := config.Build()
	require.NoError(t, err, "Failed to create log")
	ctx := NewContextWithTraceID("traceid", zap.Field{
		Key: "a", Type: zapcore.StringType, String: "b",
	})
	log.Debug(ctx, "expect")

	byteContents, err := io.ReadAll(temp)
	require.NoError(t, err, "Couldn't read log contents from temp file")
	logs := string(byteContents)
	t.Logf(logs)
	assert.Regexp(t, "expect", logs, "Unexpected log outputx.")
}
