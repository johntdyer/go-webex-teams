package spark

import (
	"encoding/json"
	"time"
)

type Person struct {
	ID          string    `json:"id"`
	Emails      string    `json:"emails"`
	Displayname string    `json:"displayName"`
	Avatar      string    `json:"avatar"`
	Created     time.Time `json:"created"`
}

type People struct {
	Items []struct {
		Person
	} `json:"items"`
}

// Messages fetches all people
func (c Client) People() (*People, error) {
	people := &People{}
	body, err := c.get(PeopleResource)
	if err != nil {
		return people, err
	}
	err = json.Unmarshal(body, &people)
	if err != nil {
		return nil, err
	}
	return people, nil
}

// Message fetches a person
func (c Client) Person(id string) (*Person, error) {
	body, err := c.get(PeopleResource + "/" + id)
	if err != nil {
		return nil, err
	}
	person := &Person{}
	err = json.Unmarshal(body, person)
	if err != nil {
		return nil, err
	}
	return person, nil
}
