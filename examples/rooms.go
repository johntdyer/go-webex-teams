/*
You must have these environment variables set to use this example:

SPARK_TOKEN must be a valid developer token
*/

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"../."
)

// Output the TrackingID HTTP header value
func displayTrackingID() {
	fmt.Println("***Request TrackingID -> " + spark.ActiveClient.TrackingID + "_" + strconv.Itoa(spark.ActiveClient.Sequence))
}

func main() {
	token := os.Getenv("SPARK_TOKEN")
	spark.InitClient(token)

	// Rooms

	// Get all rooms
	rooms := spark.Rooms{}
	err := rooms.Get()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, room := range rooms.Items {
			fmt.Println(room)
		}
		displayTrackingID()
	}

	// Get a room by ID
	room := rooms.Items[0]
	err = room.Get()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(room)
		displayTrackingID()
	}

	// // Create a room
	newRoom := &spark.Room{Title: "Project Unicorn"}
	err = newRoom.Post()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newRoom)
		displayTrackingID()
	}

	time.Sleep(3 * time.Second)
	// Update a room
	newRoom.Title = "Project Unigorn Rocks!"
	err = newRoom.Put()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newRoom)
		displayTrackingID()
	}

	time.Sleep(3 * time.Second)
	// Delete a room by ID
	err = newRoom.Delete()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Room " + newRoom.Title + " deleted!")
		displayTrackingID()
	}
}
