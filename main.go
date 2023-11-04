package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lonzzi/leetcode-daily-rank/config"
	"github.com/lonzzi/leetcode-daily-rank/data"
	"github.com/lonzzi/leetcode-daily-rank/pkg/cron"
	"github.com/lonzzi/leetcode-daily-rank/routes"
)

func main() {
	initConfig()
	conf := config.GetConfig()

	r := gin.Default()
	routes.InitRoute(r)

	r.Run(fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port))
}

func initConfig() {
	data.Init()
	cron.Init()
	config.Init()
}
