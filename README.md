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
}
```

## TODO

* Add POST/PUT methods where appropriate
