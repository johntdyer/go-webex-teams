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

	// Get all applications
	applications := teams.Applications{}
	_, err := applications.Get()
	fmt.Println(teams.TrackingID())
	if err != nil {
		fmt.Println(err)
	} else {
		for _, application := range applications.Items {
			fmt.Println("*****")
			fmt.Println(application)
		}
	}
}
