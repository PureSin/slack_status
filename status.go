package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"os"
)

func main() {
	api := slack.New(os.Getenv("SLACK_API_TOKEN"))
	// api.SetDebug(true)
	if len(os.Args) < 2 {
		fmt.Println("Need status arg: auto|away")
		return
	}
	status := os.Args[1]
	err := api.SetUserPresence(status)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}
