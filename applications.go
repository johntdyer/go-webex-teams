package spark

import (
	"encoding/json"
	"errors"
	"time"
)

// Application is how a developer requests access to perform Spark operations on behalf of a user
type Application struct {
	ID                string     `json:"id,omitempty"`
	Name              string     `json:"name,omitempty"`
	Description       string     `json:"description,omitempty"`
	Logo              string     `json:"logo,omitempty"`
	Keywords          []string   `json:"keywords,omitempty"`
	Contactemails     []string   `json:"contactEmails,omitempty"`
	Redirecturls      []string   `json:"redirectUrls,omitempty"`
	Scopes            []string   `json:"scopes,omitempty"`
	SubscriptionCount int        `json:"subscriptionCount,omitempty"`
	ClientID          string     `json:"clientId,omitempty"`
	ClientSecret      string     `json:"clientSecret,omitempty"`
	Created           *time.Time `json:"created,omitempty"`
}

// Applications represent a slice of Application
type Applications struct {
	Items []struct {
		Application
	} `json:"items"`
	Links
}

// Get fetches all applications
func (applications *Applications) Get() error {
	body, _, err := get(ApplicationsResource)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, applications)
	if err != nil {
		return err
	}
	return nil
}

func (applications *Applications) Next() error {
	if applications.NextURL != "" {
		err := applications.getCursor(applications.NextURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("next cursor not available")
	}
}

func (applications *Applications) Last() error {
	if applications.LastURL != "" {
		err := applications.getCursor(applications.LastURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("last cursor not available")
	}
}

func (applications *Applications) First() error {
	if applications.FirstURL != "" {
		err := applications.getCursor(applications.FirstURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("first cursor not available")
	}
}

func (applications *Applications) Previous() error {
	if applications.PreviousURL != "" {
		err := applications.getCursor(applications.PreviousURL)
		if err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("previous cursor not available")
	}
}

func (applications *Applications) getCursor(url string) error {
	body, links, err := get(url)
	if err != nil {
		return err
	}
	if links != nil {
		applications.Links = *links
	}
	err = json.Unmarshal(body, applications)
	if err != nil {
		return err
	}
	return nil
}

// Get fetches an application
func (app *Application) Get() error {
	body, _, err := get(ApplicationsResource + "/" + app.ID)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, app)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes an application
func (app *Application) Delete() error {
	return delete(ApplicationsResource + "/" + app.ID)
}

// Post creates a new application
func (app *Application) Post() error {
	body, err := json.Marshal(app)
	if err != nil {
		return err
	}
	body, _, err = post(ApplicationsResource, body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, app)
	if err != nil {
		return err
	}
	return nil
}

// Post updates a new application
func (app *Application) Put() error {
	body, err := json.Marshal(app)
	if err != nil {
		return err
	}
	body, _, err = put(ApplicationsResource+"/"+app.ID, body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, app)
	if err != nil {
		return err
	}
	return nil
}
