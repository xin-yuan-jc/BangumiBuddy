package main

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
	sqlitedriver "gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/MangataL/BangumiBuddy/internal/auth"
	"github.com/MangataL/BangumiBuddy/internal/repository/sqlite"
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

	db, err := gorm.Open(sqlitedriver.Open("data.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalf(ctx, "open db failed %s", err)
	}

	authRouter := ginrouter.NewAuth(auth.NewAuth(sqlite.NewUserRepo(db)))
	apisRouter := r.Group("/apis", authRouter.CheckCookie)
	apisRouter.POST("/login", authRouter.Login)
	apisRouter.POST("/logout", authRouter.Logout)
	apisRouter.POST("/user/update", authRouter.UpdateUser)
	r.NoRoute(authRouter.CheckCookie, func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	if err := r.Run("[::]:6937"); err != nil {
		log.Fatalf(ctx, "run server failed %s", err)
	}
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

func initLogger(ctx context.Context) {
	logger, err := logConfig.Build()
	if err != nil {
		log.Fatal(ctx, "init logger failed %s", err)
	}
	log.SetLogger(logger)
}
