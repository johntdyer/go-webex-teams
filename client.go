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
)

// Client is a new Spark client
type Client struct {
	Token string
	HTTP  *http.Client
}

// NewClient generates a new Spark client taking and setting the auth token
func NewClient(token string) Client {
	return Client{
		Token: token,
		HTTP:  &http.Client{},
	}
}

// Calls an HTTP DELETE
func (c Client) delete(resource string) error {
	req, _ := http.NewRequest("DELETE", BaseURL+resource, nil)
	c.setHeaders(req)
	_, err := c.HTTP.Do(req)
	return err
}

// Calls an HTTP GET
func (c Client) get(resource string) ([]byte, error) {
	req, _ := http.NewRequest("GET", BaseURL+resource, nil)
	return c.processRequest(req)
}

// Calls an HTTP POST
func (c Client) post(resource string, body []byte) ([]byte, error) {
	req, _ := http.NewRequest("POST", BaseURL+resource, bytes.NewBuffer(body))
	return c.processRequest(req)
}

// Calls an HTTP PUT
func (c Client) put(resource string, body []byte) ([]byte, error) {
	req, _ := http.NewRequest("PUT", BaseURL+resource, bytes.NewBuffer(body))
	return c.processRequest(req)
}

// Processes a HTTP POST/PUT request
func (c Client) processRequest(req *http.Request) ([]byte, error) {
	c.setHeaders(req)
	res, err := c.HTTP.Do(req)
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
func (c Client) setHeaders(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
}
