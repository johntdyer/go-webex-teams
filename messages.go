package webexTeams

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

)

// Message represents how people communicate in rooms. Individual messages are timestamped and represented in the Teams app followed by a line break.
type Message struct {
	ID          string     `json:"id,omitempty"`
	PersonID    string     `json:"personId,omitempty"`
	PersonEmail string     `json:"personEmail,omitempty"`
	RoomID      string     `json:"roomId,omitempty"`
	Text        string     `json:"text,omitempty"`
	Files       []string   `json:"files,omitempty"`
	Created     *time.Time `json:"created,omitempty"`
}

// Messages represents a collection of Messages
type Messages struct {
	Items []struct {
		Message
	} `json:"items"`
	// Used as a URL query paramter
	RoomID string
	Links
}

// Get - GETs all messages based on the provided Roomid
func (msgs *Messages) Get() (*Result, error) {
	body, links, err := get(MessagesResource + "?roomId=" + msgs.RoomID)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, msgs)
	if links != nil {
		msgs.Links = *links
	}
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Next - Moves to the next link from a link header for a large collection
func (msgs *Messages) Next() (*Result, error) {
	if msgs.NextURL != "" {
		body, err := msgs.getCursor(msgs.NextURL)
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
func (msgs *Messages) Last() (*Result, error) {
	if msgs.LastURL != "" {
		body, err := msgs.getCursor(msgs.LastURL)
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
func (msgs *Messages) First() (*Result, error) {
	if msgs.FirstURL != "" {
		body, err := msgs.getCursor(msgs.FirstURL)
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
func (msgs *Messages) Previous() (*Result, error) {
	if msgs.PreviousURL != "" {
		body, err := msgs.getCursor(msgs.PreviousURL)
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
func (msgs *Messages) getCursor(url string) ([]byte, error) {
	body, links, err := get(url)
	if err != nil {
		return body, err
	}
	if links != nil {
		msgs.Links = *links
	}
	err = json.Unmarshal(body, msgs)
	if err != nil {
		return body, err
	}
	return body, nil
}

// Get - GETs a message by ID
func (msg *Message) Get() (*Result, error) {
	body, _, err := get(MessagesResource + "/" + msg.ID)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, msg)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Delete - Deletes a message based on the ID provided
func (msg *Message) Delete() (*Result, error) {
	return delete(MessagesResource + "/" + msg.ID)
}

// Post - Creates (POSTs) a new message
func (msg *Message) Post() (*Result, error) {
	fmt.Println("ERRRRR")
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	body, _, err = post(MessagesResource, body)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, msg)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
