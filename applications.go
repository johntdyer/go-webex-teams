package spark

import (
	"encoding/json"
	"time"
)

// Application is how a developer requests access to perform Spark operations on behalf of a user
type Application struct {
	ID                string    `json:"id,omitempty"`
	Name              string    `json:"name,omitempty"`
	Description       string    `json:"description,omitempty"`
	Logo              string    `json:"logo,omitempty"`
	Keywords          []string  `json:"keywords,omitempty"`
	Contactemails     []string  `json:"contactEmails,omitempty"`
	Redirecturls      []string  `json:"redirectUrls,omitempty"`
	Scopes            []string  `json:"scopes,omitempty"`
	SubscriptionCount int       `json:"subscriptionCount,omitempty"`
	ClientID          string    `json:"clientId,omitempty"`
	ClientSecret      string    `json:"clientSecret,omitempty"`
	Created           time.Time `json:"created,omitempty"`
}

// Applications represent a slice of Application
type Applications struct {
	Items []struct {
		Application
	} `json:"items"`
}

// Applications fetches all applications
func (applications *Applications) Get() error {
	body, err := get(ApplicationsResource)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, applications)
	if err != nil {
		return err
	}
	return nil
}

// Application fetches an application
func (app *Application) Get() error {
	body, err := get(ApplicationsResource + "/" + app.ID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, app)
	if err != nil {
		return err
	}
	return nil
}

// DeleteApplication deletes an application
func (app *Application) Delete() error {
	return delete(ApplicationsResource + "/" + app.ID)
}
