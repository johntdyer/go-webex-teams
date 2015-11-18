package spark

import (
	"encoding/json"
	"time"
)

type Person struct {
	ID          string     `json:"id"`
	Emails      string     `json:"emails"`
	Displayname string     `json:"displayName"`
	Avatar      string     `json:"avatar"`
	Created     *time.Time `json:"created"`
}

type People struct {
	Items []struct {
		Person
	} `json:"items"`
}

// Messages fetches all people
func (people *People) Get() error {
	body, err := get(PeopleResource)
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
