package spark

import (
	"encoding/json"
	"errors"
	"time"
)

// Membership represent a relationship between a person and a room.
type Membership struct {
	ID          string     `json:"id,omitempty"`
	RoomID      string     `json:"roomId,omitempty"`
	PersonID    string     `json:"personId,omitempty"`
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
	RoomID      string
	PersonID    string
	PersonEmail string
	Links
}

// Get - GETs all memberships
func (memberships *Memberships) Get() (*Result, error) {
	body, _, err := get(MembershipsResource + memberships.buildQueryString())
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, memberships)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Next - Moves to the next link from a link header for a large collection
func (memberships *Memberships) Next() (*Result, error) {
	if memberships.NextURL != "" {
		body, err := memberships.getCursor(memberships.NextURL)
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
func (memberships *Memberships) Last() (*Result, error) {
	if memberships.LastURL != "" {
		body, err := memberships.getCursor(memberships.LastURL)
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
func (memberships *Memberships) First() (*Result, error) {
	if memberships.FirstURL != "" {
		body, err := memberships.getCursor(memberships.FirstURL)
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
func (memberships *Memberships) Previous() (*Result, error) {
	if memberships.PreviousURL != "" {
		body, err := memberships.getCursor(memberships.PreviousURL)
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
func (memberships *Memberships) getCursor(url string) ([]byte, error) {
	body, links, err := get(url)
	if err != nil {
		return body, err
	}
	if links != nil {
		memberships.Links = *links
	}
	err = json.Unmarshal(body, memberships)
	if err != nil {
		return body, err
	}
	return body, nil
}

// Get - Membership fetches a membership
func (membership *Membership) Get() (*Result, error) {
	body, _, err := get(MembershipsResource + "/" + membership.ID)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, membership)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Delete - DELETEs a membership
func (membership *Membership) Delete() (*Result, error) {
	return delete(MembershipsResource + "/" + membership.ID)
}

// Post - Creates (POSTs) a new membership
func (membership *Membership) Post() (*Result, error) {
	body, err := json.Marshal(membership)
	if err != nil {
		return nil, err
	}
	body, _, err = post(MembershipsResource, body)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, membership)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Put - Updates (PUTs) a membership
func (membership *Membership) Put() (*Result, error) {
	body, err := json.Marshal(membership)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	body, _, err = put(MembershipsResource+"/"+membership.ID, body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, membership)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// buildQueryString - Builds the query string
func (memberships *Memberships) buildQueryString() string {
	query := ""
	if memberships.RoomID != "" {
		query = "?roomId=" + memberships.RoomID
		if memberships.PersonID != "" {
			query += "&personId=" + memberships.PersonID
		}
		if memberships.PersonEmail != "" {
			query += "&personEmail=" + memberships.PersonEmail
		}
	} else {
		if memberships.PersonID != "" {
			query = "?personId=" + memberships.PersonID
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
