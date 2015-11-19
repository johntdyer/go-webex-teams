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
	subscriptions := spark.Subscriptions{Personid: "456"}
	err := subscriptions.Get()
	displayTrackingID()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, subscription := range subscriptions.Items {
			fmt.Println(subscription)
			fmt.Println("*****")
		}
	}
}
