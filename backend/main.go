package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	sqlitedriver "github.com/glebarez/sqlite"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"

	"github.com/MangataL/BangumiBuddy/internal/auth"
	"github.com/MangataL/BangumiBuddy/internal/auth/crypto/pbkdf2"
	"github.com/MangataL/BangumiBuddy/internal/auth/token/jwt"
	"github.com/MangataL/BangumiBuddy/internal/repository/viper"
	ginrouter "github.com/MangataL/BangumiBuddy/internal/router/gin"
	"github.com/MangataL/BangumiBuddy/pkg/log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r := gin.Default()
	// init logger
	if os.Getenv("DEV") != "true" {
		initLogger(ctx)
	}
	r.Use(log.GinLogger(), log.GinRecovery())
	// init web router
	r.LoadHTMLFiles("./web/index.html")
	r.Static("/static", "./web/static")
	r.Static("/favicon.ico", "./web/favicon.ico")

	_, err := gorm.Open(sqlitedriver.Open("data.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalf(ctx, "open db failed %s", err)
	}
	configPath := getConfigPath()
	initConfig(ctx, configPath)
	conf, err := viper.NewRepo(configPath)
	if err != nil {
		log.Fatalf(ctx, "init config failed %s", err)
	}
	authenticator := auth.New(auth.Dependencies{
		Config:        conf,
		Cipher:        pbkdf2.NewCipher(),
		TokenOperator: jwt.NewTokenOperator(),
	})

	authRouter := ginrouter.NewAuth(authenticator)
	r.POST("/apis/v1/token", authRouter.Token)
	apisRouter := r.Group("/apis/v1", authRouter.CheckToken)
	apisRouter.PUT("/user", authRouter.UpdateUser)
	r.NoRoute(authRouter.CheckToken, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	if err := r.Run("[::]:6937"); err != nil {
		log.Fatalf(ctx, "run server failed %s", err)
	}
}

const (
	defaultConfigPath = "/config/config.yaml"
)

func getConfigPath() string {
	if configPath := os.Getenv("CONFIG_FILE_PATH"); configPath != "" {
		return configPath
	}
	return defaultConfigPath
}

func initConfig(ctx context.Context, path string) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			configFile, err := os.Create(path)
			if err != nil {
				log.Fatalf(ctx, "create config file failed %s", err)
				return
			}
			_ = configFile.Close()
			return
		}
		log.Fatalf(ctx, "open config file failed %s", err)
	}
	_ = file.Close()
}

var logConfig = log.Config{
	Level: zapcore.DebugLevel,
	Caller: struct {
		Enable bool `yaml:"enable"`
		Skip   int  `yaml:"skip" json:"skip"`
	}{
		Enable: true,
	},
	Filename: "/app/data/log/log.log",
}

const (
	defaultLogPath = "/app/data/log/log.log"
)

func initLogger(ctx context.Context) {
	if logPath := os.Getenv("LOG_FILE_PATH"); logPath != "" {
		logConfig.Filename = logPath
	}
	logger, err := logConfig.Build()
	if err != nil {
		log.Fatal(ctx, "init logger failed %s", err)
	}
	log.SetLogger(logger)
}
