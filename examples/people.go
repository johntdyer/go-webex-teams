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

	// People

	// Get people
	people := teams.People{Email: "johndye@cisco.com"}
	_, err := people.Get()
	fmt.Println(teams.TrackingID())
	if err != nil {
		fmt.Println(err)
	} else {
		for _, person := range people.Items {
			fmt.Println(person)
		}
	}

	// Get a person
	person := teams.Person{ID: people.Items[0].ID}
	_, err = person.Get()
	fmt.Println(teams.TrackingID())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(person)
	}

	// Get the authenticated user
	person = teams.Person{}
	_, err = person.GetMe()
	fmt.Println(teams.TrackingID())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(person)
	}
}
