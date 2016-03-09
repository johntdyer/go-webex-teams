/*
You must have these environment variables set to use this example:

SPARK_TOKEN must be a valid developer token
SPARK_TEST_ROOM is the Room ID of the room you want to POST test messages into
*/

package main

import (
	"fmt"
	"os"

	"sqbu-github.cisco.com/jgoecke/go-spark"
)

func main() {
	authorization := &spark.Authorization{AccessToken: os.Getenv("SPARK_TOKEN")}
	spark.InitClient(authorization)

	// Get all rooms
	subscriptions := spark.Subscriptions{PersonID: "456"}
	_, err := subscriptions.Get()
	fmt.Println(spark.TrackingID())
	if err != nil {
		fmt.Println(err)
	} else {
		for _, subscription := range subscriptions.Items {
			fmt.Println(subscription)
			fmt.Println("*****")
		}
	}
}
