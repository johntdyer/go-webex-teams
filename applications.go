package spark

import (
	"encoding/json"
	"errors"
	"time"
)

// Application represents how a developer requests access to perform Spark operations on behalf of a user
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

// Get - GETs all applications
func (applications *Applications) Get() (*Result, error) {
	body, _, err := get(ApplicationsResource)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, applications)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Next - Moves to the next link from a link header for a large collection of Applications
func (applications *Applications) Next() (*Result, error) {
	if applications.NextURL != "" {
		body, err := applications.getCursor(applications.NextURL)
		if err != nil {
			result := &Result{}
			json.Unmarshal(body, result)
			return result, err
		}
		return nil, nil
	}
	return nil, errors.New("next cursor not available")
}

// Last - Moves to the last link from a link header for a large collection of Applications
func (applications *Applications) Last() (*Result, error) {
	if applications.LastURL != "" {
		body, err := applications.getCursor(applications.LastURL)
		if err != nil {
			result := &Result{}
			json.Unmarshal(body, result)
			return result, err
		}
		return nil, nil
	}
	return nil, errors.New("last cursor not available")
}

// First - Moves to the next first from a link header for a large collection of Applications
func (applications *Applications) First() (*Result, error) {
	if applications.FirstURL != "" {
		body, err := applications.getCursor(applications.FirstURL)
		if err != nil {
			result := &Result{}
			json.Unmarshal(body, result)
			return result, err
		}
		return nil, nil
	}
	return nil, errors.New("first cursor not available")
}

// Previous - Moves to the prev link from a link header for a large collection
func (applications *Applications) Previous() (*Result, error) {
	if applications.PreviousURL != "" {
		body, err := applications.getCursor(applications.PreviousURL)
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
func (applications *Applications) getCursor(url string) ([]byte, error) {
	body, links, err := get(url)
	if err != nil {
		return body, err
	}
	if links != nil {
		applications.Links = *links
	}
	err = json.Unmarshal(body, applications)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Get - GETs an application by ID
func (app *Application) Get() (*Result, error) {
	body, _, err := get(ApplicationsResource + "/" + app.ID)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, app)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Delete - DELETEs an application by ID
func (app *Application) Delete() (*Result, error) {
	return delete(ApplicationsResource + "/" + app.ID)
}

// Post - Creates (POSTs) a new application
func (app *Application) Post() (*Result, error) {
	body, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}
	body, _, err = post(ApplicationsResource, body)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, app)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Put - Updates (PUTs) an existing application
func (app *Application) Put() (*Result, error) {
	body, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}
	body, _, err = put(ApplicationsResource+"/"+app.ID, body)
	if err != nil {
		result := &Result{}
		json.Unmarshal(body, result)
		return result, err
	}
	err = json.Unmarshal(body, app)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
