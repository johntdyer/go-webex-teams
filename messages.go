package spark

import (
	"encoding/json"
	"time"
)

// Message is how people communicate in rooms. Individual messages are timestamped and represented in the Spark app followed by a line break.
type Message struct {
	ID      string     `json:"id,omitempty"`
	Roomid  string     `json:"roomId,omitempty"`
	Text    string     `json:"text,omitempty"`
	Files   []string   `json:"files,omitempty"`
	Created *time.Time `json:"created,omitempty"`
}

type Messages struct {
	Items []struct {
		Message
	} `json:"items"`
	// Used as a URL query paramter
	Roomid string
}

// Messages fetches all messages based on the provided Roomid
func (msgs *Messages) Get() error {
	body, err := get(MessagesResource + "?roomId=" + msgs.Roomid)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, msgs)
	if err != nil {
		return err
	}
	return nil
}

// Message fetches a message based on the ID provided
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

// DeleteMessage deletes a message based on the ID provided
func (msg *Message) Delete() error {
	return delete(MessagesResource + "/" + msg.ID)
}

// Post creates a new message
func (msg *Message) Post() error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	body, err = post(MessagesResource, body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, msg)
	if err != nil {
		return err
	}
	return nil
}
