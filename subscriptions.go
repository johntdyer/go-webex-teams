package spark

import (
	"encoding/json"
	"errors"
	"net/url"
	"time"
)

// Represents a relationship between a person and an application.
type Subscription struct {
	ID              string     `json:"id,omitempty"`
	Personid        string     `json:"personId,omitempty"`
	Applicationid   string     `json:"applicationId,omitempty"`
	Applicationname string     `json:"applicationName,omitempty"`
	Created         *time.Time `json:"created,omitempty"`
}

// Represents a collection of Subscriptions
type Subscriptions struct {
	Items []struct {
		Subscription
	} `json:"items"`
	Personid string
	Type     string
	Links
}

// GETs all subscriptions
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

// Moves to the next link from a link header for a large collection
func (subs *Subscriptions) Next() error {
	if subs.NextURL != "" {
		err := subs.getCursor(subs.NextURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("next cursor not available")
	}
}

// Moves to the last link from a link header for a large collection
func (subs *Subscriptions) Last() error {
	if subs.LastURL != "" {
		err := subs.getCursor(subs.LastURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("last cursor not available")
	}
}

// Moves to the first link from a link header for a large collection
func (subs *Subscriptions) First() error {
	if subs.FirstURL != "" {
		err := subs.getCursor(subs.FirstURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("first cursor not available")
	}
}

// Moves to the previous link from a link header for a large collection
func (subs *Subscriptions) Previous() error {
	if subs.PreviousURL != "" {
		err := subs.getCursor(subs.PreviousURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("previous cursor not available")
	}
}

// Gets the appropriate link associated to the link header
func (subs *Subscriptions) getCursor(url string) error {
	body, links, err := get(url)
	if err != nil {
		return err
	}
	if links != nil {
		subs.Links = *links
	}
	err = json.Unmarshal(body, subs)
	if err != nil {
		return err
	}
	return nil
}

// GETs a subscription
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

// DELETEs a subscription
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
