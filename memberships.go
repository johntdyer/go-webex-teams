package spark

import (
	"encoding/json"
	"time"
)

// Membership represent a relationship between a person and a room.
type Membership struct {
	ID          string    `json:"id,omitempty"`
	Roomid      string    `json:"roomId,omitempty"`
	Personid    string    `json:"personId,omitempty"`
	Ismoderator bool      `json:"isModerator,omitempty"`
	Ismonitor   bool      `json:"isMonitor,omitempty"`
	Islocked    bool      `json:"isLocked,omitempty"`
	PersonEmail string    `json:"personEmail,omitempty"`
	Created     time.Time `json:"created,omitempty"`
}

type Memberships struct {
	Items []struct {
		Membership
	} `json:"items"`
}

// Memberships fetches all memberships
func (c Client) Memberships() (*Memberships, error) {
	memberships := &Memberships{}
	body, err := c.get(MembershipsResource)
	if err != nil {
		return memberships, err
	}
	err = json.Unmarshal(body, memberships)
	if err != nil {
		return nil, err
	}
	return memberships, nil
}

// Membership fetches a membership
func (c Client) Membership(id string) (*Membership, error) {
	body, err := c.get(MembershipsResource + "/" + id)
	if err != nil {
		return nil, err
	}
	membership := &Membership{}
	err = json.Unmarshal(body, membership)
	if err != nil {
		return nil, err
	}
	return membership, nil
}

// DeleteMembership deletes a membership
func (c Client) DeleteMembership(id string) error {
	return c.delete(MembershipsResource + "/" + id)
}
