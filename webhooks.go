package spark

import (
	"encoding/json"
	"errors"
	"time"
)

// Webhook represents a Webhook
type Webhook struct {
	ID        string     `json:"id,omitempty"`
	Resource  string     `json:"resource,omitempty"`
	Event     string     `json:"event,omitempty"`
	Filter    string     `json:"filter,omitempty"`
	Targeturl string     `json:"targetUrl,omitempty"`
	Name      string     `json:"name,omitempty"`
	Created   *time.Time `json:"created,omitempty"`
}

// Webhooks represents a collection of Webhooks
type Webhooks struct {
	Items []struct {
		Webhook
	} `json:"items"`
	Links
}

// Get - GETs all rooms
func (webhooks *Webhooks) Get() error {
	body, _, err := get(WebhooksResource)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, webhooks)
	if err != nil {
		return err
	}
	return nil
}

// Next - Moves to the next link from a link header for a large collection
func (webhooks *Webhooks) Next() error {
	if webhooks.NextURL != "" {
		err := webhooks.getCursor(webhooks.NextURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("next cursor not available")
}

// Last - Moves to the last link from a link header for a large collection
func (webhooks *Webhooks) Last() error {
	if webhooks.LastURL != "" {
		err := webhooks.getCursor(webhooks.LastURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("last cursor not available")
}

// First - Moves to the first link from a link header for a large collection
func (webhooks *Webhooks) First() error {
	if webhooks.FirstURL != "" {
		err := webhooks.getCursor(webhooks.FirstURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("first cursor not available")
}

// Previous - Moves to the previous link from a link header for a large collection
func (webhooks *Webhooks) Previous() error {
	if webhooks.PreviousURL != "" {
		err := webhooks.getCursor(webhooks.PreviousURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("previous cursor not available")
}

// getCursor - Gets the appropriate link associated to the link header
func (webhooks *Webhooks) getCursor(url string) error {
	body, links, err := get(url)
	if err != nil {
		return err
	}
	if links != nil {
		webhooks.Links = *links
	}
	err = json.Unmarshal(body, webhooks)
	if err != nil {
		return err
	}
	return nil
}

// Get - GETs a room by ID
func (webhook *Webhook) Get() error {
	body, _, err := get(WebhooksResource + "/" + webhook.ID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		return err
	}
	return nil
}

// Delete - DELETEs a room
func (webhook *Webhook) Delete() error {
	return delete(WebhooksResource + "/" + webhook.ID)
}

// Post - Creates (POSTs) a new webhook
func (webhook *Webhook) Post() error {
	body, err := json.Marshal(webhook)
	if err != nil {
		return err
	}
	body, _, err = post(WebhooksResource, body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		return err
	}
	return nil
}

// Put - Updates (PUTs) a webhook
func (webhook *Webhook) Put() error {
	body, err := json.Marshal(webhook)
	if err != nil {
		return err
	}
	body, _, err = put(WebhooksResource+"/"+webhook.ID, body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		return err
	}
	return nil
}
