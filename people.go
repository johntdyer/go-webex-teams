package spark

import (
	"encoding/json"
	"errors"
	"net/url"
	"time"
)

// Represents a person
type Person struct {
	ID          string     `json:"id,omitempty"`
	Emails      []string   `json:"emails,omitempty"`
	Displayname string     `json:"displayName,omitempty"`
	Avatar      string     `json:"avatar,omitempty"`
	Created     *time.Time `json:"created,omitempty"`
}

// Represents a collection of Persons
type People struct {
	Items []struct {
		Person
	} `json:"items"`
	// Used as URL query parameters
	Email       string
	Displayname string
	Links
}

// GETs all people
func (people *People) Get() error {
	body, _, err := get(PeopleResource + people.buildQueryString())
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, people)
	if err != nil {
		return err
	}
	return nil
}

// Moves to the next link from a link header for a large collection
func (people *People) Next() error {
	if people.NextURL != "" {
		err := people.getCursor(people.NextURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("next cursor not available")
	}
}

// Moves to the last link from a link header for a large collection
func (people *People) Last() error {
	if people.LastURL != "" {
		err := people.getCursor(people.LastURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("last cursor not available")
	}
}

// Moves to the first link from a link header for a large collection
func (people *People) First() error {
	if people.FirstURL != "" {
		err := people.getCursor(people.FirstURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("first cursor not available")
	}
}

// Moves to the previous link from a link header for a large collection
func (people *People) Previous() error {
	if people.PreviousURL != "" {
		err := people.getCursor(people.PreviousURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("previous cursor not available")
	}
}

// Gets the appropriate link associated to the link header
func (people *People) getCursor(url string) error {
	body, links, err := get(url)
	if err != nil {
		return err
	}
	if links != nil {
		people.Links = *links
	}
	err = json.Unmarshal(body, people)
	if err != nil {
		return err
	}
	return nil
}

// GETs a person by ID
func (person *Person) Get() error {
	body, _, err := get(PeopleResource + "/" + person.ID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, person)
	if err != nil {
		return err
	}
	return nil
}

// GETs the current authenticated user
func (person *Person) GetMe() error {
	body, _, err := get(PeopleResource + "/me")
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, person)
	if err != nil {
		return err
	}
	return nil
}

// Builds the query string
func (people *People) buildQueryString() string {
	query := ""
	if people.Email != "" {
		query = "?email=" + people.Email
		if people.Displayname != "" {
			query += "&displayName=" + url.QueryEscape(people.Displayname)
		}
	} else {
		if people.Displayname != "" {
			query = "?displayName=" + url.QueryEscape(people.Displayname)
		}
	}
	return query
}
