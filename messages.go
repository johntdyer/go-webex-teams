package spark

import (
	"encoding/json"
	"errors"
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
	Links
}

// Messages fetches all messages based on the provided Roomid
func (msgs *Messages) Get() error {
	body, links, err := get(MessagesResource + "?roomId=" + msgs.Roomid)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, msgs)
	if links != nil {
		msgs.Links = *links
	}
	if err != nil {
		return err
	}
	return nil
}

func (msgs *Messages) Next() error {
	if msgs.NextURL != "" {
		err := msgs.getCursor(msgs.NextURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("next cursor not available")
	}
}

func (msgs *Messages) Last() error {
	if msgs.LastURL != "" {
		err := msgs.getCursor(msgs.LastURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("last cursor not available")
	}
}

func (msgs *Messages) First() error {
	if msgs.FirstURL != "" {
		err := msgs.getCursor(msgs.FirstURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("first cursor not available")
	}
}

func (msgs *Messages) Previous() error {
	if msgs.PreviousURL != "" {
		err := msgs.getCursor(msgs.PreviousURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("previous cursor not available")
	}
}

func (msgs *Messages) getCursor(url string) error {
	body, links, err := get(url)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, msgs)
	msgs.Links = *links
	if err != nil {
		return err
	}
	return nil
}

// Message fetches a message based on the ID provided
func (msg *Message) Get() error {
	body, _, err := get(MessagesResource + "/" + msg.ID)
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
	body, _, err = post(MessagesResource, body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, msg)
	if err != nil {
		return err
	}
	return nil
}
