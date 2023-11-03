package main

import (
	"github.com/lonzzi/leetcode-daily-rank/data"
	"github.com/lonzzi/leetcode-daily-rank/services/leetcode"
)

func main() {
	data.Init()

	err := leetcode.SaveUserProfile("lonzzi")
	if err != nil {
		panic(err)
	}
}
