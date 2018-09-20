# go-webex-teams

A [Go](https://golang.org/) library for the [WebEx Teams REST APIs](https://developer.webex.com).

## Installation

`go get github.com/johntdyer/go-webex-teams`



## Usage

```go
package main

import (
  "fmt"
  "github.com/johntdyer/go-webex-teams"
)

func main() {
	/*
	You may authorize with an AccessToken from the Developer Portal
	as follows:

	authorization := &Authorization{AccessToken: "123"}

	or you may authorize as part of an Oauth flow as follows:
	*/
	authorization := &Authorization{
		ClientID: "123",
		ClientSecret: "secret",
		Code: "567",
		RedirectURL: "http://your-server.com/auth?code=567",
	}
	teams.InitClient(authorization)
```

### Applications
```go
	// Applications (not implemented in Teams API yet)

	// Get all applications
	applications := teams.Applications{}
	result, err := applications.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(applications)

	// Get applications based on returned link header when available
	result, err := applications.Next()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(applications)
	result, err := applications.Previous()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(applications)
	result, err := applications.Last()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(applications)
	result, err := applications.First()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(applications)

	// Get an application by ID
	application := teams.Application{ID: "123"}
	result, err := application.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(application)

	// Delete an application by ID
	application := teams.Application{ID: "123"}
	result, err := application.Delete()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}

	// Create an application
	application := &Application{
		Name:          "Out of Office Assistant",
		Description:   "Does awesome things",
		Logo:          "logo.jpg",
		Keywords:      []string{"foo", "bar"},
		Contactemails: []string{"bob@foo.com", "alice@bar.org"},
		RedirectURLs:  []string{"http://myapp.com/verify", "http://myapp.fr/verify"},
		Scopes:        []string{"foo", "bar"},
	}
	result, err := application.Post()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
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
		RedirectURLs:  []string{"http://myapp.com/verify", "http://myapp.fr/verify"},
		Scopes:        []string{"foo", "bar"},
	}
	result, err := application.Put()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(application)
```

### Memberships
```go
	// Memberships

	// Get all memberships
	memberships := teams.Memberships{PersonEmail: "john@doe.com"}
	result, err := memberships.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(memberhips)

	// Get memberships based on returned link header when available
	result, err := memberships.Next()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(memberships)
	result, err := memberships.Previous()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(memberships)
	result, err := memberships.Last()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(memberships)
	result, err := memberships.First()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(memberships)

	// Get an membership by ID
	membership := teams.Membership{ID: "123"}
	result, err := membership.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(membership)

	// Delete an membership by ID
	membership := teams.Membership{ID: "123"}
	result, err := membership.Delete()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}

	// Create a membership
	membership := &Membership{
		RoomID:      "123",
		PersonID:    "456",
		PersonEmail: "john@doe.com",
		Ismoderator: true,
	}
	result, err := membership.Post()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(membership)

	// Update a membership
	membership := &Membership{
		ID: 		 "1",
		RoomID:      "123",
		PersonID:    "456",
		PersonEmail: "john@doe.com",
		Ismoderator: true,
	}
	result, err := membership.Put()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(membership)
```

### Messages
```go
	// Messages

	// Get all messages for a room
	messages := teams.Messages{RoomID: "1234"}
	result, err := messages.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(messages)

	// Get messages based on returned link header when available
	result, err := messages.Next()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(messages)
	result, err := messages.Previous()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(messages)
	result, err := messages.Last()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(messages)
	result, err := messages.First()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(messages)

	// Get an message by ID
	message := teams.Message{ID: "5678"}
	result, err := message.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(message)

	// Delete an message by ID
	message := teams.Message{ID: "5678"}
	result, err := message.Delete()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}

	// Create a message
	message := &Message{
		RoomID:	"4567",
		Text:	"This is my awesome message!",
		Files: 	[]string{"http://foo.com/image1.jpg", "http://foo.comimage2.jpg"},
	}
	result, err := message.Post()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(message)
```

### People
```go
	// People

	// Get people
	people := teams.People{PersonEmail: "john@doe.com"}
	result, err := people.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(people)

	// Get people based on returned link header when available
	result, err := people.Next()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(people)
	result, err := people.Previous()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(people)
	result, err := people.Last()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(people)
	result, err := people.First()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(people)

	// Get a person by ID
	person := teams.Person{ID: "123"}
	result, err := person.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(person

	// Get the current authenticated user
	person := teams.Person{}
	result, err := person.GetMe()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(person)```
```

### Rooms
```go
	// Rooms

	// Get all rooms
	rooms := teams.Rooms{}
	result, err := rooms.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(rooms)

	// Get rooms based on returned link header when available
	result, err := rooms.Next()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(rooms)
	result, err := rooms.Previous()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(rooms)
	result, err := rooms.Last()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(rooms)
	result, err := rooms.First()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(rooms)

	// Get a room by ID
	room := teams.Room{ID: "123"}
	result, err := room.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(room)

	// Delete a room by ID
	room := teams.Room{ID: "123"}
	result, err := room.Delete()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}

	// Create a room
	room := &Room{
		Title:   "Project Unicorn",
		Members: []string{"john@doe.com", "456"},
	}
	result, err := room.Post()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(room)

	// Update a room
	room := &Room{
		ID: 	"1",
		Title:   "Project Unicorn",
		Members: []string{"john@doe.com", "456"},
	}
	result, err := room.Put()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(room)
```

### Subscriptions
```go
	// Subscriptions (not implemented in Teams API yet)

	// Get all subscriptions
	subscriptions := teams.Subscriptions{PersonID: "abc123"}
	result, err := subscriptions.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(subscriptions)

	// Get subscriptions based on returned link header when available
	result, err := subscriptions.Next()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(subscriptions)
	result, err := subscriptions.Previous()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(subscriptions)
	result, err := subscriptions.Last()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(subscriptions)
	result, err := subscriptions.First()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(subscriptions)

	// Get a subscription by ID
	subscription := teams.Subscription{ID: "123"}
	result, err := subscription.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(subscription)

	// Delete a subscription by ID
	subscription := teams.Subscription{ID: "123"}
	result, err := room.Delete()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
```

### Webhooks
```go
	// Webhooks

	// Get all webhooks
	webhooks := teams.Webhooks{}
	result, err := webhooks.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(webhooks)

	// Get webhooks based on returned link header when available
	result, err := webhooks.Next()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(webhooks)
	result, err := webhooks.Previous()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(webhooks)
	result, err := webhooks.Last()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(webhooks)
	result, err := webhooks.First()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(webhooks)

	// Get a webhook by ID
	webhook := teams.Webhook{ID: "123"}
	result, err := webhook.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(webhook)

	// Delete a webhook by ID
	webhook := teams.Webhook{ID: "123"}
	result, err := room.Delete()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}

	// Create a webhook
	webhook := &Webhook{
		Resource:  "messages",
		Event:     "created",
		Filter:    "room=123",
		Targeturl: "http://foo.com/bar",
		Name:      "My Awesome webhook",
	}
	result, err := webhook.Put()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(webhook)
}
```

## Notes

* To use the examples in the /examples folder, you must set an environment variable WEBEX_TEAMS_TOKEN. You may obtain a token by logging in @ [http://developer.webex.com](http://developer.webex.com).
* This library also implements the TrackingID header used to trace requests in the Teams platform. If troubleshooting and working with our support, output the current values, seen in the example above, and send along with your support request.
