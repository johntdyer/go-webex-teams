package spark

import (
	"encoding/json"
	"net/url"
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
	Personid string
	Type     string
}

// Subscriptions fetches all subscriptions
func (subs *Subscriptions) Get() error {
	body, _, err := get(SubscriptionsResource + subs.buildQueryString())
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
	body, _, err := get(SubscriptionsResource + "/" + sub.ID)
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

// Builds the query string
func (subscriptions *Subscriptions) buildQueryString() string {
	query := ""
	if subscriptions.Personid != "" {
		query = "?personId=" + subscriptions.Personid
		if subscriptions.Type != "" {
			query += "&type=" + url.QueryEscape(subscriptions.Type)
		}
	} else {
		if subscriptions.Type != "" {
			query = "?type=" + url.QueryEscape(subscriptions.Type)
		}
	}
	return query
}
