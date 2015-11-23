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

// Memberships represents a collection of Memberships
type Memberships struct {
	Items []struct {
		Membership
	} `json:"items"`
	Roomid      string
	Personid    string
	PersonEmail string
	Links
}

// Get - GETs all memberships
func (memberships *Memberships) Get() error {
	body, _, err := get(MembershipsResource + memberships.buildQueryString())
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, memberships)
	if err != nil {
		return err
	}
	return nil
}

// Next - Moves to the next link from a link header for a large collection
func (memberships *Memberships) Next() error {
	if memberships.NextURL != "" {
		err := memberships.getCursor(memberships.NextURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("next cursor not available")

}

// Last - Moves to the last link from a link header for a large collection
func (memberships *Memberships) Last() error {
	if memberships.LastURL != "" {
		err := memberships.getCursor(memberships.LastURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("last cursor not available")

}

// First - Moves to the first link from a link header for a large collection
func (memberships *Memberships) First() error {
	if memberships.FirstURL != "" {
		err := memberships.getCursor(memberships.FirstURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("first cursor not available")
}

// Previous - Moves to the previous link from a link header for a large collection
func (memberships *Memberships) Previous() error {
	if memberships.PreviousURL != "" {
		err := memberships.getCursor(memberships.PreviousURL)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("previous cursor not available")
}

// getCursor - Gets the appropriate link associated to the link header
func (memberships *Memberships) getCursor(url string) error {
	body, links, err := get(url)
	if err != nil {
		return err
	}
	if links != nil {
		memberships.Links = *links
	}
	err = json.Unmarshal(body, memberships)
	if err != nil {
		return err
	}
	return nil
}

// Get - Membership fetches a membership
func (membership *Membership) Get() error {
	body, _, err := get(MembershipsResource + "/" + membership.ID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, membership)
	if err != nil {
		return err
	}
	return nil
}

// Delete - DELETEs a membership
func (membership *Membership) Delete() error {
	return delete(MembershipsResource + "/" + membership.ID)
}

// Post - Creates (POSTs) a new membership
func (membership *Membership) Post() error {
	body, err := json.Marshal(membership)
	if err != nil {
		return err
	}
	body, _, err = post(MembershipsResource, body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, membership)
	if err != nil {
		return err
	}
	return nil
}

// Put - Updates (PUTs) a membership
func (membership *Membership) Put() error {
	body, err := json.Marshal(membership)
	if err != nil {
		return err
	}
	body, _, err = put(MembershipsResource+"/"+membership.ID, body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, membership)
	if err != nil {
		return err
	}
	return nil
}

// buildQueryString - Builds the query string
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
