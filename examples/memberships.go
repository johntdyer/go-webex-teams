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

	// Get all memberships for a room
	memberships := teams.Memberships{RoomID: os.Getenv("WEBEX_TEAMS_TEST_ROOM")}
	_, err := memberships.Get()
	fmt.Println(teams.TrackingID())
	if err != nil {
		fmt.Println(err)
	} else {
		for _, membership := range memberships.Items {
			fmt.Println(membership)
			fmt.Println("*****")
		}
	}
}
