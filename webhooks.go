package spark

import (
	"encoding/json"
	"time"
)

type Webhook struct {
	ID        string     `json:"id,omitempty"`
	Resource  string     `json:"resource,omitempty"`
	Event     string     `json:"event,omitempty"`
	Filter    string     `json:"filter,omitempty"`
	Targeturl string     `json:"targetUrl,omitempty"`
	Name      string     `json:"name,omitempty"`
	Created   *time.Time `json:"created,omitempty"`
}

type Webhooks struct {
	Items []struct {
		Webhook
	} `json:"items"`
}

// Rooms fetches all rooms
func (webhooks *Webhooks) Get() error {
	body, err := get(WebhooksResource)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, webhooks)
	if err != nil {
		return err
	}
	return nil
}

// Room fetchs a room
func (webhook *Webhook) Get() error {
	body, err := get(WebhooksResource + "/" + webhook.ID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		return err
	}
	return nil
}

// DeleteRoom deletes a room
func (webhook *Webhook) Delete() error {
	return delete(WebhooksResource + "/" + webhook.ID)
}

// Post creates a new webhook
func (webhook *Webhook) Post() error {
	body, err := json.Marshal(webhook)
	if err != nil {
		return err
	}
	body, err = post(WebhooksResource, body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		return err
	}
	return nil
}

// Post updates a webhook
func (webhook *Webhook) Put() error {
	body, err := json.Marshal(webhook)
	if err != nil {
		return err
	}
	body, err = put(WebhooksResource+"/"+webhook.ID, body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		return err
	}
	return nil
}
