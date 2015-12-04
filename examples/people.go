/*
You must have these environment variables set to use this example:

SPARK_TOKEN must be a valid developer token
*/

package main

import (
	"fmt"
	"os"

	"sqbu-github.cisco.com/jgoecke/go-spark"
)

func main() {
	authorization := &Authorization{os.Getenv("SPARK_TOKEN")}
	spark.InitClient(authorization)

	// People

	// Get people
	people := spark.People{Email: "jgoecke@cisco.com"}
	err := people.Get()
	fmt.Println(spark.TrackingID())
	if err != nil {
		fmt.Println(err)
	} else {
		for _, person := range people.Items {
			fmt.Println(person)
		}
	}

	// Get a person
	person := spark.Person{ID: people.Items[0].ID}
	err = person.Get()
	fmt.Println(spark.TrackingID())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(person)
	}

	// Get the authenticated user
	person = spark.Person{}
	err = person.GetMe()
	fmt.Println(spark.TrackingID())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(person)
	}
}
