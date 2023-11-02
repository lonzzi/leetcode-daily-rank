package main

import (
	"fmt"
	"github.com/lonzzi/leetcode-daily-rank/pkg/leetcode"
)

func main() {
	response, err := leetcode.GetUserProfilePublicProfile("lonzzi")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}
