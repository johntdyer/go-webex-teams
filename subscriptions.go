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
	PersonID        string     `json:"personId,omitempty"`
	ApplicationID   string     `json:"applicationId,omitempty"`
	Applicationname string     `json:"applicationName,omitempty"`
	Created         *time.Time `json:"created,omitempty"`
}

// Subscriptions represents a collection of Subscriptions
type Subscriptions struct {
	Items []struct {
		Subscription
	} `json:"items"`
	PersonID string
	Type     string
	Links
}

// Get - GETs all subscriptions
func (subscriptions *Subscriptions) Get() (*Result, error) {
	body, _, err := get(SubscriptionsResource + subscriptions.buildQueryString())
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, subscriptions)
	if err != nil {
		return nil, err
	}
	return nil, nil

}

// Next - Moves to the next link from a link header for a large collection
func (subscriptions *Subscriptions) Next() (*Result, error) {
	if subscriptions.NextURL != "" {
		body, err := subscriptions.getCursor(subscriptions.NextURL)
		if err != nil {
			result := &Result{}
			json.Unmarshal(body, result)
			return result, err
		}
		return nil, nil
	}
	return nil, errors.New("next cursor not available")
}

// Last - Moves to the last link from a link header for a large collection
func (subscriptions *Subscriptions) Last() (*Result, error) {
	if subscriptions.LastURL != "" {
		body, err := subscriptions.getCursor(subscriptions.LastURL)
		if err != nil {
			result := &Result{}
			json.Unmarshal(body, result)
			return result, err
		}
		return nil, nil
	}
	return nil, errors.New("last cursor not available")
}

// First - Moves to the first link from a link header for a large collection
func (subscriptions *Subscriptions) First() (*Result, error) {
	if subscriptions.FirstURL != "" {
		body, err := subscriptions.getCursor(subscriptions.FirstURL)
		if err != nil {
			result := &Result{}
			json.Unmarshal(body, result)
			return result, err
		}
		return nil, nil
	}
	return nil, errors.New("first cursor not available")
}

// Previous - Moves to the previous link from a link header for a large collection
func (subscriptions *Subscriptions) Previous() (*Result, error) {
	if subscriptions.PreviousURL != "" {
		body, err := subscriptions.getCursor(subscriptions.PreviousURL)
		if err != nil {
			result := &Result{}
			json.Unmarshal(body, result)
			return result, err
		}
		return nil, nil
	}
	return nil, errors.New("previous cursor not available")
}

// getCursor - Gets the appropriate link associated to the link header
func (subscriptions *Subscriptions) getCursor(url string) ([]byte, error) {
	body, links, err := get(url)
	if err != nil {
		return body, err
	}
	if links != nil {
		subscriptions.Links = *links
	}
	err = json.Unmarshal(body, subscriptions)
	if err != nil {
		return body, err
	}
	return body, nil
}

// Get - GETs a subscription
func (subscription *Subscription) Get() (*Result, error) {
	body, _, err := get(SubscriptionsResource + "/" + subscription.ID)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, subscription)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Delete - DELETEs a subscription
func (subscription *Subscription) Delete() (*Result, error) {
	return delete(SubscriptionsResource + "/" + subscription.ID)
}

// buildQueryString - Builds the query string
func (subscriptions *Subscriptions) buildQueryString() string {
	query := ""
	if subscriptions.PersonID != "" {
		query = "?personId=" + subscriptions.PersonID
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
