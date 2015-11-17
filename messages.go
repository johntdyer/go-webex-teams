package spark

import (
	"encoding/json"
	"time"
)

// Message is how people communicate in rooms. Individual messages are timestamped and represented in the Spark app followed by a line break.
type Message struct {
	ID      string    `json:"id,omitempty"`
	Roomid  string    `json:"roomId,omitempty"`
	Text    string    `json:"text,omitempty"`
	Files   []string  `json:"files,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

type Messages struct {
	Items []struct {
		Message
	} `json:"items"`
}

// Messages fetches all messages
func (c Client) Messages() (*Messages, error) {
	messages := &Messages{}
	body, err := c.get(MessagesResource)
	if err != nil {
		return messages, err
	}
	err = json.Unmarshal(body, &messages)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

// Message fetches a message
func (c Client) Message(id string) (*Message, error) {
	body, err := c.get(MessagesResource + "/" + id)
	if err != nil {
		return nil, err
	}
	message := &Message{}
	err = json.Unmarshal(body, message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

// DeleteMessage deletes a message
func (c Client) DeleteMessage(id string) error {
	return c.delete(MessagesResource + "/" + id)
}
