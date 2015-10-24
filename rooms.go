package spark

import (
	"encoding/json"
	"time"
)

// Room is virtual meeting places for getting work done.
type Room struct {
	ID      string    `json:"id,omitempty"`
	Title   string    `json:"title,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

// Rooms represent a slice of Room
type Rooms []struct {
	Room
}

// Rooms fetches all rooms
func (c Client) Rooms() (Rooms, error) {
	body, err := c.get(RoomsResource)
	if err != nil {
		return nil, err
	}
	rooms := make(Rooms, 0)
	err = json.Unmarshal(body, &rooms)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

// Room fetchs a room
func (c Client) Room(id string) (*Room, error) {
	body, err := c.get(RoomsResource + "/" + id)
	if err != nil {
		return nil, err
	}
	room := &Room{}
	err = json.Unmarshal(body, room)
	if err != nil {
		return nil, err
	}
	return room, nil
}

// DeleteRoom deletes a room
func (c Client) DeleteRoom(id string) error {
	return c.delete(RoomsResource + "/" + id)
}
