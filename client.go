package spark

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	// BaseURL              = "https://api.ciscospark.com/v1"
	BaseURL = "https://api.ciscospark.com/hydra/api/v1"
	// ApplicationsResource is the resource for managing applications
	ApplicationsResource = "/applications"
	// MembershipsResource is the resource for managing memberships
	MembershipsResource = "/memberships"
	// PeopleResource is the resource for managing people
	PeopleResource = "/people"
	// MessagesResource is the resource for managing messages
	MessagesResource = "/messages"
	// RoomsResource is the resource for managing rooms
	RoomsResource = "/rooms"
	// SubscriptionsResource is the resource for managing subscriptions
	SubscriptionsResource = "/subscriptions"
	// WebhooksResource is the resource for managing webhooks
	WebhooksResource = "/webhooks"
	// ActiveClient represents the client used to connect to the Spark API
	ActiveClient = &Client{}
	// ErrInactiveClient raises if you have not done an InitClient()
	ErrInactiveClient = errors.New("You must call InitalizeClient() before using this operation")
)

// Client represents a new Spark client
type Client struct {
	Token            string
	TrackingIDPrefix string
	Sequence         int
	Increment        chan int
	Finished         chan bool
	HTTP             *http.Client
}

// Links struct to store pagination details from a link header
type Links struct {
	NextURL     string
	LastURL     string
	FirstURL    string
	PreviousURL string
}

// InitClient - Generates a new Spark client taking and setting the auth token
func InitClient(token string) {
	ActiveClient = &Client{
		Token:            token,
		HTTP:             &http.Client{},
		TrackingIDPrefix: "go-spark_" + uuid(),
		Sequence:         0,
		Increment:        make(chan int),
		Finished:         make(chan bool),
	}
	go incrementer()
}

// Increment - Updates the sequence of the request to ensure a unique identifier
// for each API request via the TrackingID header
func incrementer() {
	for {
		<-ActiveClient.Increment
		ActiveClient.Sequence++
		ActiveClient.Finished <- true
	}
}

// delete - Calls an HTTP DELETE
func delete(resource string) error {
	req, _ := http.NewRequest("DELETE", BaseURL+resource, nil)
	_, _, err := processRequest(req)
	return err
}

// get - Calls an HTTP GET
func get(resource string) ([]byte, *Links, error) {
	req, _ := http.NewRequest("GET", BaseURL+resource, nil)
	return processRequest(req)
}

// post - Calls an HTTP POST
func post(resource string, body []byte) ([]byte, *Links, error) {
	req, _ := http.NewRequest("POST", BaseURL+resource, bytes.NewBuffer(body))
	return processRequest(req)
}

// put - Calls an HTTP PUT
func put(resource string, body []byte) ([]byte, *Links, error) {
	req, _ := http.NewRequest("PUT", BaseURL+resource, bytes.NewBuffer(body))
	return processRequest(req)
}

// processRequest - Processes an HTTP request
func processRequest(req *http.Request) ([]byte, *Links, error) {
	if ActiveClient.Token == "" {
		return nil, nil, ErrInactiveClient
	}
	setHeaders(req)
	res, err := ActiveClient.HTTP.Do(req)
	if err != nil {
		return nil, nil, err
	}
	if res.StatusCode != 200 {
		return nil, nil, errors.New(res.Status)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	linkHeader := res.Header.Get("Link")
	if linkHeader != "" {
		links := parseLink(linkHeader)
		return body, links, nil
	}
	return body, nil, nil
}

// setHeaders - Set the headers for the HTTP requests
func setHeaders(req *http.Request) {
	ActiveClient.Increment <- 1
	<-ActiveClient.Finished
	req.Header.Set("Authorization", "Bearer "+ActiveClient.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("TrackingID", TrackingID())
}

// parseLink - Parses a link header
func parseLink(link string) *Links {
	links := &Links{}
	items := strings.Split(link, ",")
	for _, item := range items {
		ele := strings.Split(item, `; rel="`)
		switch strings.TrimRight(ele[1], `"`) {
		case "next":
			links.NextURL = parseURL(ele[0])
		case "last":
			links.LastURL = parseURL(ele[0])
		case "first":
			links.FirstURL = parseURL(ele[0])
		case "prev":
			links.PreviousURL = parseURL(ele[0])
		}
	}
	return links
}

// parseURL - Parses a URL from within a link header
func parseURL(url string) string {
	url = strings.TrimRight(url, ">")
	url = strings.TrimLeft(url, " <")
	url = strings.TrimLeft(url, "<")
	return url
}

// TrackingID returns the current value used to set the
// unique TrackingID HTTP header with each request to the
// Spark API
func TrackingID() string {
	return ActiveClient.TrackingIDPrefix + "_" + strconv.Itoa(ActiveClient.Sequence)
}

// uuid - Creates a UUID to be used for a unique identifier
// for the tracing TrackingID header
func uuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(-1)
		return ""
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
