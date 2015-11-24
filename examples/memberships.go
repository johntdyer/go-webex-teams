/*
You must have these environment variables set to use this example:

SPARK_TOKEN must be a valid developer token
SPARK_TEST_ROOM is the Room ID of the room you want to POST test messages into
*/

package main

import (
	"fmt"
	"os"

	"../."
)

func main() {
	spark.InitClient(os.Getenv("SPARK_TOKEN"))

	// Get all memberships for a room
	memberships := spark.Memberships{Roomid: os.Getenv("SPARK_TEST_ROOM")}
	err := memberships.Get()
	fmt.Println(spark.TrackingID())
	if err != nil {
		fmt.Println(err)
	} else {
		for _, membership := range memberships.Items {
			fmt.Println(membership)
			fmt.Println("*****")
		}
	}
}
