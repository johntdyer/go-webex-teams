package spark

import (
	"encoding/json"
	"errors"
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
	Links
}

// Memberships fetches all memberships
func (mems *Memberships) Get() error {
	body, _, err := get(MembershipsResource + mems.buildQueryString())
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, mems)
	if err != nil {
		return err
	}
	return nil
}

func (mems *Memberships) Next() error {
	if mems.NextURL != "" {
		err := mems.getCursor(mems.NextURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("next cursor not available")
	}
}

func (mems *Memberships) Last() error {
	if mems.LastURL != "" {
		err := mems.getCursor(mems.LastURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("last cursor not available")
	}
}

func (mems *Memberships) First() error {
	if mems.FirstURL != "" {
		err := mems.getCursor(mems.FirstURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("first cursor not available")
	}
}

func (mems *Memberships) Previous() error {
	if mems.PreviousURL != "" {
		err := mems.getCursor(mems.PreviousURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("previous cursor not available")
	}
}

func (mems *Memberships) getCursor(url string) error {
	body, links, err := get(url)
	if err != nil {
		return err
	}
	if links != nil {
		mems.Links = *links
	}
	err = json.Unmarshal(body, mems)
	if err != nil {
		return err
	}
	return nil
}

// Membership fetches a membership
func (mem *Membership) Get() error {
	body, _, err := get(MembershipsResource + "/" + mem.ID)
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
	body, _, err = post(MembershipsResource, body)
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
	body, _, err = put(MembershipsResource+"/"+mem.ID, body)
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
