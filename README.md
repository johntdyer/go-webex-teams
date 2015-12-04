# go-spark

A [Go](https://golang.org/) library for the [Spark REST APIs](https://developer.ciscospark.com).

## Installation

`go get sqbu-github.cisco.com/jgoecke/go-spark`

If this fails, you may also do:

```
cd $GOPATH/src
mkdir sqbu-github.cisco.com
cd sqbu-github.cisco.com
mkdir jgoecke
cd jgoecke
git clone https://sqbu-github.cisco.com/jgoecke/go-spark.git
```

## Usage

```go
package main

import (
  "fmt"
  "sqbu-github.cisco.com/jgoecke/go-spark"
)

func main() {
	spark.InitClient("<YOUR TOKEN>")
```

### Applications
```go
	// Applications (not implemented in Spark API yet)

	// Get all applications
	applications := spark.Applications{}
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
	application := spark.Application{ID: "123"}
	result, err := application.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(application)

	// Delete an application by ID
	application := spark.Application{ID: "123"}
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
	memberships := spark.Memberships{PersonEmail: "john@doe.com"}
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
	membership := spark.Membership{ID: "123"}
	result, err := membership.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(membership)

	// Delete an membership by ID
	membership := spark.Membership{ID: "123"}
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
	messages := spark.Messages{RoomID: "1234"}
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
	message := spark.Message{ID: "5678"}
	result, err := message.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(message)

	// Delete an message by ID
	message := spark.Message{ID: "5678"}
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
	people := spark.People{PersonEmail: "john@doe.com"}
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
	person := spark.Person{ID: "123"}
	result, err := person.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(person

	// Get the current authenticated user
	person := spark.Person{}
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
	rooms := spark.Rooms{PersonID: "abc123"}
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
	room := spark.Room{ID: "123"}
	result, err := room.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(room)

	// Delete a room by ID
	room := spark.Room{ID: "123"}
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
	// Subscriptions (not implemented in Spark API yet)
	
	// Get all subscriptions
	subscriptions := spark.Subscriptions{PersonID: "abc123"}
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
	subscription := spark.Subscription{ID: "123"}
	result, err := subscription.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(subscription)

	// Delete a subscription by ID
	subscription := spark.Subscription{ID: "123"}
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
	webhooks := spark.Webhooks{}
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
	webhook := spark.Webhook{ID: "123"}
	result, err := webhook.Get()
	if err != nil {
		fmt.Println(err)
		fmt.Println(result)
	}
	fmt.Println(webhook)

	// Delete a webhook by ID
	webhook := spark.Webhook{ID: "123"}
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

* To use the examples in the /examples folder, you must set an environment variable SPARK_TOKEN. You may obtain a token by logging in @ [http://developer.ciscospark.com](http://developer.ciscospark.com).
* This library also implements the TrackingID header used to trace requests in the Spark platform. If troubleshooting and working with our support, output the current values, seen in the example above, and send along with your support request.
