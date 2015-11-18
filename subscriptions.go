package spark

import (
	"encoding/json"
	"time"
)

// Subscription represents a relationship between a person and an application.
type Subscription struct {
	ID              string     `json:"id,omitempty"`
	Personid        string     `json:"personId,omitempty"`
	Applicationid   string     `json:"applicationId,omitempty"`
	Applicationname string     `json:"applicationName,omitempty"`
	Created         *time.Time `json:"created,omitempty"`
}

// Subscriptions represent a slice of Subscription
type Subscriptions struct {
	Items []struct {
		Subscription
	} `json:"items"`
}

// Subscriptions fetches all subscriptions
func (subs *Subscriptions) Get() error {
	body, err := get(SubscriptionsResource)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, subs)
	if err != nil {
		return err
	}
	return nil

}

// Subscription fetches a subscription
func (sub *Subscription) Get() error {
	body, err := get(SubscriptionsResource + "/" + sub.ID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, sub)
	if err != nil {
		return err
	}
	return nil
}

// DeleteSubscription deletes a subscription
func (sub *Subscription) Delete() error {
	return delete(SubscriptionsResource + "/" + sub.ID)
}
