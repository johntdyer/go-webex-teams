/*
You must have these environment variables set to use this example:

SPARK_TOKEN must be a valid developer token
*/

package main

import (
	"fmt"
	"os"

	"../."
)

func main() {
	spark.InitClient(os.Getenv("SPARK_TOKEN"))

	// Get all applications
	applications := spark.Applications{}
	err := applications.Get()
	fmt.Println(spark.TrackingID())
	if err != nil {
		fmt.Println(err)
	} else {
		for _, application := range applications.Items {
			fmt.Println("*****")
			fmt.Println(application)
		}
	}
}
