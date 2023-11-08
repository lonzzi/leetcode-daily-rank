package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lonzzi/leetcode-daily-rank/config"
	"github.com/lonzzi/leetcode-daily-rank/data"
	"github.com/lonzzi/leetcode-daily-rank/pkg/cron"
	"github.com/lonzzi/leetcode-daily-rank/routes"
	"github.com/lonzzi/leetcode-daily-rank/services/leetcode"
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

	// user, _ := leetcode.GetUsersByRank()
	// for _, u := range user {
	// 	fmt.Println(u)
	// }

	SaveUsers := func(conf *config.Config) func() {
		return func() {
			for _, user := range conf.LeetCode.UserSlug {
				leetcode.SaveUserProfile(user)
			}
		}
	}

	conf := config.GetConfig()
	SaveUsers(&conf)()
	cron.AddFunc("@every 10m", SaveUsers(&conf))

	cron.Start()
}
