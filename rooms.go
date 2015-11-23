package spark

import (
	"encoding/json"
	"errors"
	"time"
)

// Room is virtual meeting places for getting work done.
type Room struct {
	ID         string     `json:"id,omitempty"`
	Title      string     `json:"title,omitempty"`
	SIPAddress string     `json:"sipAddress,omitempty"`
	Members    []string   `json:"members,omitempty"`
	Created    *time.Time `json:"created,omitempty"`
}

type Rooms struct {
	Items []struct {
		Room
	} `json:"items"`
	Links
}

// Rooms fetches all rooms
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

func (rooms *Rooms) Next() error {
	if rooms.NextURL != "" {
		err := rooms.getCursor(rooms.NextURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("next cursor not available")
	}
}

func (rooms *Rooms) Last() error {
	if rooms.LastURL != "" {
		err := rooms.getCursor(rooms.LastURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("last cursor not available")
	}
}

func (rooms *Rooms) First() error {
	if rooms.FirstURL != "" {
		err := rooms.getCursor(rooms.FirstURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("first cursor not available")
	}
}

func (rooms *Rooms) Previous() error {
	if rooms.PreviousURL != "" {
		err := rooms.getCursor(rooms.PreviousURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("previous cursor not available")
	}
}

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

// Room fetchs a room
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

// DeleteRoom deletes a room
func (room *Room) Delete() error {
	return delete(RoomsResource + "/" + room.ID)
}

// Post creates a new room
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

// Post updates a room
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
