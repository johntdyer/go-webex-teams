/*
You must have these environment variables set to use this example:

SPARK_TOKEN must be a valid developer token
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

	// Rooms

	// // Get all rooms
	// rooms := spark.Rooms{}
	// err := rooms.Get()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for _, room := range rooms.Items {
	// 		fmt.Println(room)
	// 	}
	// 	fmt.Println(spark.TrackingID())
	// }

	// // Get a room by ID
	// room := rooms.Items[0]
	// err = room.Get()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(room)
	// 	fmt.Println(spark.TrackingID())
	// }

	// // Create a room
	newRoom := &spark.Room{Title: "IFTTT Notifications Room"}
	_, err := newRoom.Post()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newRoom)
		fmt.Println(spark.TrackingID())
	}

	// time.Sleep(3 * time.Second)
	// // Update a room
	// newRoom.Title = "Project Unigorn Rocks!"
	// err = newRoom.Put()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(newRoom)
	// 	fmt.Println(spark.TrackingID())
	// }

	// time.Sleep(3 * time.Second)
	// // Delete a room by ID
	// err = newRoom.Delete()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Room " + newRoom.Title + " deleted!")
	// 	fmt.Println(spark.TrackingID())
	// }
}
