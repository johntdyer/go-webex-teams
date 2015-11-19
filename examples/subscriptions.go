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
	memberships := spark.Memberships{PersonEmail: "jgoecke@cisco.com"}
	err := memberships.Get()
	displayTrackingID()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, membership := range memberships.Items {
			fmt.Println(membership)
			fmt.Println("*****")
		}
	}
}
