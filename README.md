# go-spark

A [Go](https://golang.org/) library for the [Spark REST APIs](https://developer.ciscospark.com).

## Usage

```go
package main

import (
  "fmt"
  "github.com/jsgoecke/go-spark"
)

func main() {
	spark.InitClient("<YOUR TOKEN>")

	// Applications

	// Get all applications
	applications := spark.Applications{}
	err := applications.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.println(applications)

	// Get an application by ID
	application := spark.Application{ID: "123"}
	err := application.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.println(application)

	// Delete an application by ID
	application := spark.Application{ID: "123"}
	err := application.Delete()
	if err != nil {
		fmt.Println(err)
	}

	// Create an application
	application := &Application{
		Name:          "Out of Office Assistant",
		Description:   "Does awesome things",
		Logo:          "logo.jpg",
		Keywords:      []string{"foo", "bar"},
		Contactemails: []string{"bob@foo.com", "alice@bar.org"},
		Redirecturls:  []string{"http://myapp.com/verify", "http://myapp.fr/verify"},
		Scopes:        []string{"foo", "bar"},
	}
	err := application.Post()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(application)

	// Update an application by ID
	application := &Application{
		ID: 			"1",
		Name:          "Out of Office Assistant",
		Description:   "Does awesome things",
		Logo:          "logo.jpg",
		Keywords:      []string{"foo", "bar"},
		Contactemails: []string{"bob@foo.com", "alice@bar.org"},
		Redirecturls:  []string{"http://myapp.com/verify", "http://myapp.fr/verify"},
		Scopes:        []string{"foo", "bar"},
	}
	err := application.Put()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(application)

	// Memberships
	
	// Get all memberships
	memberships := spark.Memberships{}
	err := memberships.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.println(memberhips)

	// Get an membership by ID
	membership := spark.Membership{ID: "123"}
	err := membership.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.println(membership)

	// Delete an membership by ID
	membership := spark.Membership{ID: "123"}
	err := membership.Delete()
	if err != nil {
		fmt.Println(err)
	}

	// Create a membership
	membership := &Membership{
		Roomid:      "123",
		Personid:    "456",
		PersonEmail: "john@doe.com",
		Ismoderator: true,
	}
	err := membership.Post()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(membership)

	// Update a membership
	membership := &Membership{
		ID: 		 "1",
		Roomid:      "123",
		Personid:    "456",
		PersonEmail: "john@doe.com",
		Ismoderator: true,
	}
	err := membership.Put()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(membership)

	// People
	
	// Get all people
	people := spark.People{}
	err := people.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.println(people)

	// Get a person by ID
	person := spark.Person{ID: "123"}
	err := person.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.println(person)

	// Rooms
	
	// Get all rooms
	rooms := spark.Rooms{}
	err := rooms.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.println(rooms)

	// Get a room by ID
	room := spark.Room{ID: "123"}
	err := room.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.println(room)

	// Delete a room by ID
	room := spark.Room{ID: "123"}
	err := room.Delete()
	if err != nil {
		fmt.Println(err)
	}

	// Create a room
	room := &Room{
		Title:   "Project Unicorn",
		Members: []string{"john@doe.com", "456"},
	}
	err := room.Post()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(room)

	// Update a room
	room := &Room{
		ID: 	"1",
		Title:   "Project Unicorn",
		Members: []string{"john@doe.com", "456"},
	}
	err := room.Put()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(room)

	// Subscriptions
	
	// Get all subscriptions
	subscriptions := spark.Subscriptions{}
	err := subscriptions.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.println(subscriptions)

	// Get a subscription by ID
	subscription := spark.Subscription{ID: "123"}
	err := subscription.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.println(subscription)

	// Delete a subscription by ID
	subscription := spark.Subscription{ID: "123"}
	err := room.Delete()
	if err != nil {
		fmt.Println(err)
	}

	// Webhooks
	
	// Get all webhooks
	webhooks := spark.Webhooks{}
	err := webhooks.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.println(webhooks)

	// Get a webhook by ID
	webhook := spark.Webhook{ID: "123"}
	err := webhook.Get()
	if err != nil {
		fmt.Println(err)
	}
	fmt.println(webhook)

	// Delete a webhook by ID
	webhook := spark.Webhook{ID: "123"}
	err := room.Delete()
	if err != nil {
		fmt.Println(err)
	}

	// Create a webhook
	webhook := &Webhook{
		Resource:  "messages",
		Event:     "created",
		Filter:    "room=123",
		Targeturl: "http://foo.com/bar",
		Name:      "My Awesome webhook",
	}
	err := webhook.Put()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(webhook)
}
```

## Examples

To use the examples in the /examples folder, you must set an environment variable SPARK_TOKEN. You may obtain a token by logging in @ [http://developer.ciscospark.com](http://developer.ciscospark.com).

