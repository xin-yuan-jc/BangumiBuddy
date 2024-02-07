package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config logger config
type Config struct {
	Level  zapcore.Level `yaml:"level" json:"level"` // log level，debug,info,warn,error,fatal
	Caller struct {
		Enable bool `yaml:"enable"`           // Whether to print function call information, such as file line number, function name, etc.
		Skip   int  `yaml:"skip" json:"skip"` // Tracking depth, default is 2
	} `yaml:"caller"`
	Filename string `yaml:"fileName" json:"fileName"` // log file path, absolute path
}

// Build new logger from Config
func (c *Config) Build() (*ZapLogger, error) {
	lv := zap.NewAtomicLevelAt(c.Level)
	var opts []zap.Option

	if c.Caller.Enable {
		opts = append(opts, zap.AddCaller())
		if c.Caller.Skip == 0 {
			c.Caller.Skip = 2
		}
		opts = append(opts, zap.AddCallerSkip(c.Caller.Skip))
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	file, err := os.Create(c.Filename)
	if err != nil {
		return nil, err
	}
	return NewZapLogger(
		zapcore.NewCore(encoder, zapcore.AddSync(file), lv),
		&lv,
		opts...), nil
}
