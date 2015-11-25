package spark

import (
	"encoding/json"
	"errors"
	"time"
)

// Room represets a virtual meeting places for getting work done.
type Room struct {
	ID         string     `json:"id,omitempty"`
	Title      string     `json:"title,omitempty"`
	SIPAddress string     `json:"sipAddress,omitempty"`
	Members    []string   `json:"members,omitempty"`
	Created    *time.Time `json:"created,omitempty"`
}

// Rooms represents a collection of Rooms
type Rooms struct {
	Items []struct {
		Room
	} `json:"items"`
	Links
}

// Get - GETs all rooms
func (rooms *Rooms) Get() error {
	body, _, err := get(RoomsResource)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, rooms)
	if err != nil {
		return err
	}
	return nil
}

// Next - Moves to the next link from a link header for a large collection
func (rooms *Rooms) Next() error {
	if rooms.NextURL != "" {
		err := rooms.getCursor(rooms.NextURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("next cursor not available")
}

// Last - Moves to the last link from a link header for a large collection
func (rooms *Rooms) Last() error {
	if rooms.LastURL != "" {
		err := rooms.getCursor(rooms.LastURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("last cursor not available")
}

// First - Moves to the first link from a link header for a large collection
func (rooms *Rooms) First() error {
	if rooms.FirstURL != "" {
		err := rooms.getCursor(rooms.FirstURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("first cursor not available")
}

// Previous - Moves to the previous link from a link header for a large collection
func (rooms *Rooms) Previous() error {
	if rooms.PreviousURL != "" {
		err := rooms.getCursor(rooms.PreviousURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("previous cursor not available")
}

// getCursor - Gets the appropriate link associated to the link header
func (rooms *Rooms) getCursor(url string) error {
	body, links, err := get(url)
	if err != nil {
		return err
	}
	if links != nil {
		rooms.Links = *links
	}
	err = json.Unmarshal(body, rooms)
	if err != nil {
		return err
	}
	return nil
}

// Get - GETs a room by ID
func (room *Room) Get() error {
	body, _, err := get(RoomsResource + "/" + room.ID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, room)
	if err != nil {
		return err
	}
	return nil
}

// Delete - DELETEs a room
func (room *Room) Delete() (*Result, error) {
	return delete(RoomsResource + "/" + room.ID)
}

// Post - Creates (POSTs) a new room
func (room *Room) Post() error {
	body, err := json.Marshal(room)
	if err != nil {
		return err
	}
	body, _, err = post(RoomsResource, body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, room)
	if err != nil {
		return err
	}
	return nil
}

// Put - Updates (PUTs) a room
func (room *Room) Put() error {
	body, err := json.Marshal(room)
	if err != nil {
		return err
	}
	body, _, err = put(RoomsResource+"/"+room.ID, body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, room)
	if err != nil {
		return err
	}
	return nil
}
