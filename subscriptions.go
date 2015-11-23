package spark

import (
	"encoding/json"
	"errors"
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

// Subscriptions represents a collection of Subscriptions
type Subscriptions struct {
	Items []struct {
		Subscription
	} `json:"items"`
	Personid string
	Type     string
	Links
}

// Get - GETs all subscriptions
func (subscriptions *Subscriptions) Get() error {
	body, _, err := get(SubscriptionsResource + subscriptions.buildQueryString())
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, subscriptions)
	if err != nil {
		return err
	}
	return nil

}

// Next - Moves to the next link from a link header for a large collection
func (subscriptions *Subscriptions) Next() error {
	if subscriptions.NextURL != "" {
		err := subscriptions.getCursor(subscriptions.NextURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("next cursor not available")
}

// Last - Moves to the last link from a link header for a large collection
func (subscriptions *Subscriptions) Last() error {
	if subscriptions.LastURL != "" {
		err := subscriptions.getCursor(subscriptions.LastURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("last cursor not available")
}

// First - Moves to the first link from a link header for a large collection
func (subscriptions *Subscriptions) First() error {
	if subscriptions.FirstURL != "" {
		err := subscriptions.getCursor(subscriptions.FirstURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("first cursor not available")
}

// Previous - Moves to the previous link from a link header for a large collection
func (subscriptions *Subscriptions) Previous() error {
	if subscriptions.PreviousURL != "" {
		err := subscriptions.getCursor(subscriptions.PreviousURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("previous cursor not available")
}

// getCursor - Gets the appropriate link associated to the link header
func (subscriptions *Subscriptions) getCursor(url string) error {
	body, links, err := get(url)
	if err != nil {
		return err
	}
	if links != nil {
		subscriptions.Links = *links
	}
	err = json.Unmarshal(body, subscriptions)
	if err != nil {
		return err
	}
	return nil
}

// Get - GETs a subscription
func (subscription *Subscription) Get() error {
	body, _, err := get(SubscriptionsResource + "/" + subscription.ID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, subscription)
	if err != nil {
		return err
	}
	return nil
}

// Delete - DELETEs a subscription
func (subscription *Subscription) Delete() error {
	return delete(SubscriptionsResource + "/" + subscription.ID)
}

// buildQueryString - Builds the query string
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
