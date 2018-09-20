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
	rooms := teams.Rooms{}
	_, err := rooms.Get()
	fmt.Println(teams.TrackingID())
	if err != nil {
		fmt.Println(err)
	}

	// Get all messages from a room
	messages := teams.Messages{RoomID: rooms.Items[1].ID}
	_, err = messages.Get()
	fmt.Println(teams.TrackingID())
	fmt.Println("Displaying messages for room -> " + rooms.Items[1].Title)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, message := range messages.Items {
			fmt.Println("*****")
			fmt.Println(message)
		}
		// Display an individual message
		message := teams.Message{ID: messages.Items[0].ID}
		_, err = message.Get()
		fmt.Println(teams.TrackingID())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("+++++")
			fmt.Println(message)
		}
	}

	// Create a message in the test room of WEBEX_TEAMS_TEST_ROOM
	message := teams.Message{
		RoomID: os.Getenv("WEBEX_TEAMS_TEST_ROOM"),
		Text:   "Hello for go-webex-teams!",
		Files:  []string{"http://49.media.tumblr.com/0cb3d95bf6263c0e27e2b141d0bd9409/tumblr_nnjbwx8fNo1rf78nfo1_500.gif"},
	}
	_, err = message.Post()
	fmt.Println(teams.TrackingID())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message)
	}

	// Delete the posted message
	// Not implemented
	// err = message.Delete()
	// fmt.Println(teams.TrackingID())
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(message)
	// }
}
