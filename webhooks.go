package webexTeams

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
	TargetURL string     `json:"targetUrl,omitempty"`
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
func (webhooks *Webhooks) Get() (*Result, error) {
	body, _, err := get(WebhooksResource)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, webhooks)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Next - Moves to the next link from a link header for a large collection
func (webhooks *Webhooks) Next() (*Result, error) {
	if webhooks.NextURL != "" {
		body, err := webhooks.getCursor(webhooks.NextURL)
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
func (webhooks *Webhooks) Last() (*Result, error) {
	if webhooks.LastURL != "" {
		body, err := webhooks.getCursor(webhooks.LastURL)
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
func (webhooks *Webhooks) First() (*Result, error) {
	if webhooks.FirstURL != "" {
		body, err := webhooks.getCursor(webhooks.FirstURL)
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
func (webhooks *Webhooks) Previous() (*Result, error) {
	if webhooks.PreviousURL != "" {
		body, err := webhooks.getCursor(webhooks.PreviousURL)
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
func (webhooks *Webhooks) getCursor(url string) ([]byte, error) {
	body, links, err := get(url)
	if err != nil {
		return body, err
	}
	if links != nil {
		webhooks.Links = *links
	}
	err = json.Unmarshal(body, webhooks)
	if err != nil {
		return body, err
	}
	return body, nil
}

// Get - GETs a room by ID
func (webhook *Webhook) Get() (*Result, error) {
	body, _, err := get(WebhooksResource + "/" + webhook.ID)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Delete - DELETEs a room
func (webhook *Webhook) Delete() (*Result, error) {
	return delete(WebhooksResource + "/" + webhook.ID)
}

// Post - Creates (POSTs) a new webhook
func (webhook *Webhook) Post() (*Result, error) {
	body, err := json.Marshal(webhook)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	body, _, err = post(WebhooksResource, body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Put - Updates (PUTs) a webhook
func (webhook *Webhook) Put() (*Result, error) {
	body, err := json.Marshal(webhook)
	if err != nil {
		return nil, err
	}
	body, _, err = put(WebhooksResource+"/"+webhook.ID, body)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
