package spark

import (
	"encoding/json"
	"time"
)

// Membership represent a relationship between a person and a room.
type Membership struct {
	ID          string     `json:"id,omitempty"`
	Roomid      string     `json:"roomId,omitempty"`
	Personid    string     `json:"personId,omitempty"`
	Ismoderator bool       `json:"isModerator,omitempty"`
	Ismonitor   bool       `json:"isMonitor,omitempty"`
	Islocked    bool       `json:"isLocked,omitempty"`
	PersonEmail string     `json:"personEmail,omitempty"`
	Created     *time.Time `json:"created,omitempty"`
}

type Memberships struct {
	Items []struct {
		Membership
	} `json:"items"`
	Roomid      string
	Personid    string
	PersonEmail string
}

// Memberships fetches all memberships
func (mems *Memberships) Get() error {
	body, err := get(MembershipsResource + mems.buildQueryString())
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, mems)
	if err != nil {
		return err
	}
	return nil
}

// Membership fetches a membership
func (mem *Membership) Get() error {
	body, err := get(MembershipsResource + "/" + mem.ID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, mem)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMembership deletes a membership
func (mem *Membership) Delete() error {
	return delete(MembershipsResource + "/" + mem.ID)
}

// Post creates a new membership
func (mem *Membership) Post() error {
	body, err := json.Marshal(mem)
	if err != nil {
		return err
	}
	body, err = post(MembershipsResource, body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, mem)
	if err != nil {
		return err
	}
	return nil
}

// Post updates a membership
func (mem *Membership) Put() error {
	body, err := json.Marshal(mem)
	if err != nil {
		return err
	}
	body, err = put(MembershipsResource+"/"+mem.ID, body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, mem)
	if err != nil {
		return err
	}
	return nil
}

// Builds the query string
func (memberships *Memberships) buildQueryString() string {
	query := ""
	if memberships.Roomid != "" {
		query = "?roomId=" + memberships.Roomid
		if memberships.Personid != "" {
			query += "&personId=" + memberships.Personid
		}
		if memberships.PersonEmail != "" {
			query += "&personEmail=" + memberships.PersonEmail
		}
	} else {
		if memberships.Personid != "" {
			query = "?personId=" + memberships.Personid
			if memberships.PersonEmail != "" {
				query += "&personEmail=" + memberships.PersonEmail
			}
		} else {
			if memberships.PersonEmail != "" {
				query += "?personEmail=" + memberships.PersonEmail
			}
		}
	}
	return query
}
