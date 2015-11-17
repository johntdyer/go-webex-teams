package spark

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
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
	//
	InactiveClientErr = errors.New("You must call InitalizeClient() before using this operation")
)

// Client is a new Spark client
type Client struct {
	Token string
	HTTP  *http.Client
}

// IntClient generates a new Spark client taking and setting the auth token
func InitClient(token string) {
	ActiveClient = &Client{
		Token: token,
		HTTP:  &http.Client{},
	}
}

// Calls an HTTP DELETE
func delete(resource string) error {
	req, _ := http.NewRequest("DELETE", BaseURL+resource, nil)
	_, err := processRequest(req)
	return err
}

// Calls an HTTP GET
func get(resource string) ([]byte, error) {
	req, _ := http.NewRequest("GET", BaseURL+resource, nil)
	return processRequest(req)
}

// Calls an HTTP POST
func post(resource string, body []byte) ([]byte, error) {
	req, _ := http.NewRequest("POST", BaseURL+resource, bytes.NewBuffer(body))
	return processRequest(req)
}

// Calls an HTTP PUT
func put(resource string, body []byte) ([]byte, error) {
	req, _ := http.NewRequest("PUT", BaseURL+resource, bytes.NewBuffer(body))
	return processRequest(req)
}

// Processes a HTTP POST/PUT request
func processRequest(req *http.Request) ([]byte, error) {
	if ActiveClient.Token == "" {
		return nil, InactiveClientErr
	}
	setHeaders(req)
	res, err := ActiveClient.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Set the headers for the HTTP requests
func setHeaders(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+ActiveClient.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
}
