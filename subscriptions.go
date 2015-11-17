package spark

import (
	"encoding/json"
	"time"
)

// Subscription represents a relationship between a person and an application.
type Subscription struct {
	ID              string    `json:"id,omitempty"`
	Personid        string    `json:"personId,omitempty"`
	Applicationid   string    `json:"applicationId,omitempty"`
	Applicationname string    `json:"applicationName,omitempty"`
	Created         time.Time `json:"created,omitempty"`
}

// Subscriptions represent a slice of Subscription
type Subscriptions struct {
	Items []struct {
		Subscription
	} `json:"items"`
}

// Subscriptions fetches all subscriptions
func (c Client) Subscriptions() (*Subscriptions, error) {
	subscriptions := &Subscriptions{}
	body, err := c.get(SubscriptionsResource)
	if err != nil {
		return subscriptions, err
	}
	err = json.Unmarshal(body, &subscriptions)
	if err != nil {
		return subscriptions, err
	}
	return subscriptions, nil

}

// Subscription fetches a subscription
func (c Client) Subscription(id string) (*Subscription, error) {
	body, err := c.get(SubscriptionsResource + "/" + id)
	if err != nil {
		return nil, err
	}
	subscription := &Subscription{}
	err = json.Unmarshal(body, subscription)
	if err != nil {
		return nil, err
	}
	return subscription, nil
}

// DeleteSubscription deletes a subscription
func (c Client) DeleteSubscription(id string) error {
	return c.delete(SubscriptionsResource + "/" + id)
}
