package spark

import (
	"encoding/json"
	"net/url"
	"time"
)

type Person struct {
	ID          string     `json:"id,omitempty"`
	Emails      []string   `json:"emails,omitempty"`
	Displayname string     `json:"displayName,omitempty"`
	Avatar      string     `json:"avatar,omitempty"`
	Created     *time.Time `json:"created,omitempty"`
}

type People struct {
	Items []struct {
		Person
	} `json:"items"`
	// Used as URL query parameters
	Email       string
	Displayname string
}

// Messages fetches all people
func (people *People) Get() error {
	body, err := get(PeopleResource + people.buildQueryString())
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, people)
	if err != nil {
		return err
	}
	return nil
}

// Message fetches a person
func (person *Person) Get() error {
	body, err := get(PeopleResource + "/" + person.ID)
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
