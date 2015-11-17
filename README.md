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
	sparkClient := spark.NewClient("<YOUR TOKEN>")

	// Interact with applications
	applications := sparkClient.Applications()
	fmt.println(applications)
	application := sparkClient.Application("123")
	fmt.println(application)
	sparkClient.DeleteApplication("123")

	// Interact with memberships
	memberships := sparkClient.Memberships()
	fmt.println(memberships)
	membership := sparkClient.Membership("456")
	fmt.println(membership)
	sparkClient.DeleteApplication("456")

	// Interact with messages
	messages := sparkClient.Messages()
	fmt.println(messages)
	message := sparkClient.Message("789")
	fmt.println(message)
	sparkClient.DeleteMessage("789")

	// Interact with people
	people := sparkClient.People()
	fmt.println(people)
	person := sparkClient.Person("789")
	fmt.println(person)

	// Interact with rooms
	rooms := sparkClient.Rooms()
	fmt.println(rooms)
	room := sparkClient.Room("901")
	fmt.println(room)
	sparkClient.DeleteRoom("901")

	// Interact with subscriptions
	subscriptions := sparkClient.Subscriptions()
	fmt.println(subscriptions)
	subscription := sparkClient.Subscription("901")
	fmt.println(subscription)
	sparkClient.DeleteSubscription("901")

	// Interact with webhooks
	webhooks := sparkClient.Webhooks()
	fmt.println(webhooks)
	webhook := sparkClient.Webhook("901")
	fmt.println(webhook)
	sparkClient.DeleteWebhook("901")
}
```

## TODO

* Add POST/PUT methods where appropriate
