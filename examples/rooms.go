/*
You must have these environment variables set to use this example:

WEBEX_TEAMS_TOKEN must be a valid developer token
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

	// Rooms

	// // Get all rooms
	// rooms := teams.Rooms{}
	// err := rooms.Get()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for _, room := range rooms.Items {
	// 		fmt.Println(room)
	// 	}
	// 	fmt.Println(teams.TrackingID())
	// }

	// // Get a room by ID
	// room := rooms.Items[0]
	// err = room.Get()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(room)
	// 	fmt.Println(teams.TrackingID())
	// }

	// // Create a room
	newRoom := &teams.Room{Title: "IFTTT Notifications Room"}
	_, err := newRoom.Post()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(newRoom)
		fmt.Println(teams.TrackingID())
	}

	// time.Sleep(3 * time.Second)
	// // Update a room
	// newRoom.Title = "Project Unigorn Rocks!"
	// err = newRoom.Put()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(newRoom)
	// 	fmt.Println(teams.TrackingID())
	// }

	// time.Sleep(3 * time.Second)
	// // Delete a room by ID
	// err = newRoom.Delete()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Room " + newRoom.Title + " deleted!")
	// 	fmt.Println(teams.TrackingID())
	// }
}
