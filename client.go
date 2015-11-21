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
	// MessagesResource is the resource for managing messages
	PeopleResource = "/people"
	// PeopleResource is the resource for managing people
	MessagesResource = "/messages"
	// RoomsResource is the resource for managing rooms
	RoomsResource = "/rooms"
	// SubscriptionsResource is the resource for managing subscriptions
	SubscriptionsResource = "/subscriptions"
	// WebhooksResource is the resource for managing webhooks
	WebhooksResource = "/webhooks"
	// ActiveClient
	ActiveClient = &Client{}
	// InactiveClientErr raises if you have not done an InitClient()
	InactiveClientErr = errors.New("You must call InitalizeClient() before using this operation")
)

// Client is a new Spark client
type Client struct {
	Token      string
	TrackingID string
	Sequence   int
	Increment  chan int
	Finished   chan bool
	HTTP       *http.Client
}

// Links for pagination
type Links struct {
	NextURL     string
	LastURL     string
	FirstURL    string
	PreviousURL string
}

// IntClient generates a new Spark client taking and setting the auth token
func InitClient(token string) {
	ActiveClient = &Client{
		Token:      token,
		HTTP:       &http.Client{},
		TrackingID: "go-spark_" + uuid(),
		Sequence:   0,
		Increment:  make(chan int),
		Finished:   make(chan bool),
	}
	go incrementer()
}

// incrementer updates the Sequence of the request
func incrementer() {
	for {
		<-ActiveClient.Increment
		ActiveClient.Sequence++
		ActiveClient.Finished <- true
	}
}

// Calls an HTTP DELETE
func delete(resource string) error {
	req, _ := http.NewRequest("DELETE", BaseURL+resource, nil)
	_, _, err := processRequest(req)
	return err
}

// Calls an HTTP GET
func get(resource string) ([]byte, *Links, error) {
	req, _ := http.NewRequest("GET", BaseURL+resource, nil)
	return processRequest(req)
}

// Calls an HTTP POST
func post(resource string, body []byte) ([]byte, *Links, error) {
	req, _ := http.NewRequest("POST", BaseURL+resource, bytes.NewBuffer(body))
	return processRequest(req)
}

// Calls an HTTP PUT
func put(resource string, body []byte) ([]byte, *Links, error) {
	req, _ := http.NewRequest("PUT", BaseURL+resource, bytes.NewBuffer(body))
	return processRequest(req)
}

// Processes a HTTP POST/PUT request
func processRequest(req *http.Request) ([]byte, *Links, error) {
	if ActiveClient.Token == "" {
		return nil, nil, InactiveClientErr
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
	link := res.Header.Get("Link")
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	if link != "" {
		links := parseLink(link)
		return body, links, nil
	}
	return body, nil, nil
}

// Set the headers for the HTTP requests
func setHeaders(req *http.Request) {
	ActiveClient.Increment <- 1
	<-ActiveClient.Finished
	req.Header.Set("Authorization", "Bearer "+ActiveClient.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("TrackingID", ActiveClient.TrackingID+"_"+strconv.Itoa(ActiveClient.Sequence))
}

// Parses a link header
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

// Parses a URL from within a link header
func parseURL(url string) string {
	url = strings.TrimRight(url, ">")
	url = strings.TrimLeft(url, " <")
	url = strings.TrimLeft(url, "<")
	return url
}

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
