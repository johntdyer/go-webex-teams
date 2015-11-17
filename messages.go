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
func (msgs *Messages) Get() error {
	body, err := get(MessagesResource)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, msgs)
	if err != nil {
		return err
	}
	return nil
}

// Message fetches a message
func (msg *Message) Get() error {
	body, err := get(MessagesResource + "/" + msg.ID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, msg)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMessage deletes a message
func (msg *Message) Delete() error {
	return delete(MessagesResource + "/" + msg.ID)
}
