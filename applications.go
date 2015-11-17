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
func (c Client) Applications() (*Applications, error) {
	applications := &Applications{}
	body, err := c.get(ApplicationsResource)
	if err != nil {
		return applications, err
	}
	err = json.Unmarshal(body, &applications)
	if err != nil {
		return applications, err
	}
	return applications, nil
}

// Application fetches an application
func (c Client) Application(id string) (*Application, error) {
	body, err := c.get(ApplicationsResource + "/" + id)
	if err != nil {
		return nil, err
	}
	application := &Application{}
	err = json.Unmarshal(body, application)
	if err != nil {
		return nil, err
	}
	return application, nil
}

// DeleteApplication deletes an application
func (c Client) DeleteApplication(id string) error {
	return c.delete(ApplicationsResource + "/" + id)
}
