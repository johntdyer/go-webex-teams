/*
You must have these environment variables set to use this example:

WEBEX_TEAMS_TOKEN must be a valid developer token
WEBEX_TEAMS_TEST_ROOM is the Room ID of the room you want to POST test messages into
*/

package main

import (
	"fmt"
	"os"

	"github.com/johntdyer/go-webex-teams"
)

func main() {
	authorization := &teams.Authorization{AccessToken: os.Getenv("WEBEX_TEAMS_TOKEN")}
	teams.InitClient(authorization)

	// Get all rooms
	webhooks := teams.Webhooks{}
	_, err := webhooks.Get()
	fmt.Println(teams.TrackingID())
	if err != nil {
		fmt.Println(err)
	} else {
		for _, webhook := range webhooks.Items {
			fmt.Println(webhook)
			fmt.Println("*****")
		}
	}

	// {
	//            "resource" : "messages",
	//            "event" : "created",
	//            "filter" : "roomId=Y2lzY29zcGFyazovL3VzL1JPT00vYmJjZWIxYWQtNDNmMS0zYjU4LTkxNDctZjE0YmIwYzRkMTU0",
	//            "targetUrl" : "https://example.com/mywebhook",
	//            "name" : "My Awesome Webhook"
	//          }

	webhook := teams.Webhook{
		Resource:  "messages",
		Event:     "created",
		Filter:    "roomId=Y2lzY29zcGFyazovL3VzL1JPT00vYmJjZWIxYWQtNDNmMS0zYjU4LTkxNDctZjE0YmIwYzRkMTU0",
		TargetURL: "https://example.com/hook",
		Name:      "Awesomesauce",
	}
	_, err = webhook.Post()
	fmt.Println(teams.TrackingID())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(webhook)
	}
}
