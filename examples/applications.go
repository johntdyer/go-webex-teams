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
	spark.InitClient(os.Getenv("SPARK_TOKEN"))

	// Get all applications
	applications := spark.Applications{}
	err := applications.Get()
	displayTrackingID()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, application := range applications.Items {
			fmt.Println("*****")
			fmt.Println(application)
		}
	}
}
