/*
You must have these environment variables set to use this example:

SPARK_TOKEN must be a valid developer token
SPARK_TEST_ROOM is the Room ID of the room you want to POST test messages into
*/

package main

import (
	"fmt"
	"os"
	"strconv"

	"../."
)

// Output the TrackingID HTTP header value
func displayTrackingID() {
	fmt.Println("***Request TrackingID -> " + spark.ActiveClient.TrackingID + "_" + strconv.Itoa(spark.ActiveClient.Sequence))
}

func main() {
	spark.InitClient(os.Getenv("SPARK_TOKEN"))

	// Get all rooms
	webhooks := spark.Webhooks{}
	err := webhooks.Get()
	displayTrackingID()
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

	webhook := spark.Webhook{
		Resource:  "messages",
		Event:     "created",
		Filter:    "roomId=Y2lzY29zcGFyazovL3VzL1JPT00vYmJjZWIxYWQtNDNmMS0zYjU4LTkxNDctZjE0YmIwYzRkMTU0",
		Targeturl: "https://example.com/hook",
		Name:      "Awesomesauce",
	}
	err = webhook.Post()
	displayTrackingID()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(webhook)
	}
}
