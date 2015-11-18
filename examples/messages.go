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
	rooms := spark.Rooms{}
	err := rooms.Get()
	displayTrackingID()
	if err != nil {
		fmt.Println(err)
	}

	// Get all messages from a room
	messages := spark.Messages{Roomid: rooms.Items[1].ID}
	err = messages.Get()
	displayTrackingID()
	fmt.Println("Displaying messages for room -> " + rooms.Items[1].Title)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, message := range messages.Items {
			fmt.Println("*****")
			fmt.Println(message)
		}
		// Display an individual message
		message := spark.Message{ID: messages.Items[0].ID}
		err = message.Get()
		displayTrackingID()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("+++++")
			fmt.Println(message)
		}
	}

	// Create a message in the test room of SPARK_TEST_ROOM
	message := spark.Message{
		Roomid: os.Getenv("SPARK_TEST_ROOM"),
		Text:   "Hello for go-spark!",
		Files:  []string{"http://49.media.tumblr.com/0cb3d95bf6263c0e27e2b141d0bd9409/tumblr_nnjbwx8fNo1rf78nfo1_500.gif"},
	}
	err = message.Post()
	displayTrackingID()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message)
	}

	// Delete the posted message
	// Not implemented
	// err = message.Delete()
	// displayTrackingID()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(message)
	// }
}
