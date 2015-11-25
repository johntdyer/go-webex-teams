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
func (rooms *Rooms) Get() (*Result, error) {
	body, _, err := get(RoomsResource)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, rooms)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Next - Moves to the next link from a link header for a large collection
func (rooms *Rooms) Next() (*Result, error) {
	if rooms.NextURL != "" {
		body, err := rooms.getCursor(rooms.NextURL)
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
func (rooms *Rooms) Last() (*Result, error) {
	if rooms.LastURL != "" {
		body, err := rooms.getCursor(rooms.LastURL)
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
func (rooms *Rooms) First() (*Result, error) {
	if rooms.FirstURL != "" {
		body, err := rooms.getCursor(rooms.FirstURL)
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
func (rooms *Rooms) Previous() (*Result, error) {
	if rooms.PreviousURL != "" {
		body, err := rooms.getCursor(rooms.PreviousURL)
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
func (rooms *Rooms) getCursor(url string) ([]byte, error) {
	body, links, err := get(url)
	if err != nil {
		return body, err
	}
	if links != nil {
		rooms.Links = *links
	}
	err = json.Unmarshal(body, rooms)
	if err != nil {
		return body, err
	}
	return body, nil
}

// Get - GETs a room by ID
func (room *Room) Get() (*Result, error) {
	body, _, err := get(RoomsResource + "/" + room.ID)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, room)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Delete - DELETEs a room
func (room *Room) Delete() (*Result, error) {
	return delete(RoomsResource + "/" + room.ID)
}

// Post - Creates (POSTs) a new room
func (room *Room) Post() (*Result, error) {
	body, err := json.Marshal(room)
	if err != nil {
		return nil, err
	}
	body, _, err = post(RoomsResource, body)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, room)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Put - Updates (PUTs) a room
func (room *Room) Put() (*Result, error) {
	body, err := json.Marshal(room)
	if err != nil {
		return nil, err
	}
	body, _, err = put(RoomsResource+"/"+room.ID, body)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, room)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
