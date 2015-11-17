package spark

import (
	"encoding/json"
	"time"
)

type Webhook struct {
	ID        string    `json:"id"`
	Resource  string    `json:"resource"`
	Event     string    `json:"event"`
	Filter    string    `json:"filter"`
	Targeturl string    `json:"targetUrl"`
	Name      string    `json:"name"`
	Created   time.Time `json:"created,omitempty"`
}

type Webhooks struct {
	Items []struct {
		Webhook
	} `json:"items"`
}

// Rooms fetches all rooms
func (c Client) Webhooks() (*Webhooks, error) {
	webhooks := &Webhooks{}
	body, err := c.get(WebhooksResource)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &webhooks)
	if err != nil {
		return nil, err
	}
	return webhooks, nil
}

// Room fetchs a room
func (c Client) Webhook(id string) (*Webhook, error) {
	body, err := c.get(WebhooksResource + "/" + id)
	if err != nil {
		return nil, err
	}
	webhook := &Webhook{}
	err = json.Unmarshal(body, webhook)
	if err != nil {
		return nil, err
	}
	return webhook, nil
}

// DeleteRoom deletes a room
func (c Client) DeleteWebhook(id string) error {
	return c.delete(WebhooksResource + "/" + id)
}
