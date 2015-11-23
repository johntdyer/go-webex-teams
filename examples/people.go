/*
You must have these environment variables set to use this example:

SPARK_TOKEN must be a valid developer token
*/

package main

import (
	"fmt"
	"os"
	"strconv"

	"../."
)

// Output the TrackingID HTTP header value
func displayTrackingID() {
	fmt.Println("***Request TrackingID -> " + spark.ActiveClient.TrackingID + "_" + strconv.Itoa(spark.ActiveClient.Sequence))
}

func main() {
	token := os.Getenv("SPARK_TOKEN")
	spark.InitClient(token)

	// People

	// Get people
	people := spark.People{Email: "jgoecke@cisco.com"}
	err := people.Get()
	displayTrackingID()
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
	displayTrackingID()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(person)
	}

	// Get the authenticated user
	person = spark.Person{}
	err = person.GetMe()
	displayTrackingID()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(person)
	}
}
