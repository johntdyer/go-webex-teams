package webexTeams

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	ApplicationJSON         = `{"id":"123","name":"foo","description":"bar","logo":"image.jpg","keywords":["foo","bar"],"contactEmails":["john@doe.com","jane@doe.com"],"redirectUrls":["http://1.com","http://2.com"],"scopes":["scope1","scope2"],"subscriptionCount":1000,"clientId":"123","clientSecret":"456","created":"2015-10-18T07:26:16Z"}`
	ApplicationsJSON        = `{"items":[` + ApplicationJSON + `]}`
	ApplicationResponseJSON = `{"id":"123","name":"Out of Office Assistant","description":"Does awesome things","logo":"logo.jpg","keywords":["foo","bar"],"contactEmails":["bob@foo.com","alice@bar.org"],"redirectUrls":["http://myapp.com/verify","http://myapp.fr/verify"],"scopes":["foo","bar"],"subscriptionCount":1000,"clientId":"456","clientSecret":"secret","created":"2015-10-18T07:26:16Z"}`
)

func TestApplicationsSpec(t *testing.T) {
	ts := serveHTTP(t)
	defer ts.Close()
	previousURL := BaseURL
	BaseURL = ts.URL
	InitClient(&Authorization{AccessToken: "123"})
	Convey("Given we want to interact with Teams applications", t, func() {
		Convey("For an application", func() {
			Convey("It should generate the proper JSON message", func() {
				application := &Application{
					ID:                "123",
					Name:              "foo",
					Description:       "bar",
					Logo:              "image.jpg",
					Keywords:          []string{"foo", "bar"},
					RedirectURLs:      []string{"http://1.com", "http://2.com"},
					Contactemails:     []string{"john@doe.com", "jane@doe.com"},
					Scopes:            []string{"scope1", "scope2"},
					SubscriptionCount: 1000,
					ClientID:          "123",
					ClientSecret:      "456",
					Created:           &CreatedTime,
				}
				body, err := json.Marshal(application)
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, ApplicationJSON)
			})
			Convey("It should generate the proper struct from the JSON", func() {
				application := &Application{}
				err := json.Unmarshal([]byte(ApplicationJSON)[:], application)
				So(err, ShouldBeNil)
				validateApplication(t, application)
			})
			Convey("Get application", func() {
				application := &Application{ID: "1"}
				result, err := application.Get()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				validateApplication(t, application)
			})
			Convey("Delete application", func() {
				application := &Application{ID: "1"}
				result, err := application.Delete()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
			})
			Convey("Post application", func() {
				application := &Application{
					Name:          "Out of Office Assistant",
					Description:   "Does awesome things",
					Logo:          "logo.jpg",
					Keywords:      []string{"foo", "bar"},
					Contactemails: []string{"bob@foo.com", "alice@bar.org"},
					RedirectURLs:  []string{"http://myapp.com/verify", "http://myapp.fr/verify"},
					Scopes:        []string{"foo", "bar"},
				}
				result, err := application.Post()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
			})
			Convey("Put application", func() {
				application := &Application{
					ID:            "1",
					Name:          "Out of Office Assistant",
					Description:   "Does awesome things",
					Logo:          "logo.jpg",
					Keywords:      []string{"foo", "bar"},
					Contactemails: []string{"bob@foo.com", "alice@bar.org"},
					RedirectURLs:  []string{"http://myapp.com/verify", "http://myapp.fr/verify"},
					Scopes:        []string{"foo", "bar"},
				}
				result, err := application.Put()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
			})
		})
		Convey("For applications", func() {
			Convey("It should generate the proper struct from the JSON", func() {
				applications := &Applications{}
				err := json.Unmarshal([]byte(ApplicationsJSON)[:], &applications)
				So(err, ShouldBeNil)
				validateApplications(t, applications)
			})
			Convey("Get applications", func() {
				applications := &Applications{}
				result, err := applications.Get()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				validateApplications(t, applications)
			})
			Convey("It should raise an error when no link cursor", func() {
				applications := &Applications{}
				result, err := applications.First()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "first cursor not available")
				result, err = applications.Next()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "next cursor not available")
				result, err = applications.Last()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "last cursor not available")
				result, err = applications.Previous()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "previous cursor not available")
			})
			Convey("Should get our link cursor", func() {
				applications := &Applications{}
				applications.FirstURL = "/applications/first"
				result, err := applications.First()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(applications.Items[0].ID, ShouldEqual, "123")
				applications.LastURL = "/applications/last"
				result, err = applications.Last()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(applications.Items[0].ID, ShouldEqual, "123")
				applications.NextURL = "/applications/next"
				result, err = applications.Next()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(applications.Items[0].ID, ShouldEqual, "123")
				applications.PreviousURL = "/applications/prev"
				result, err = applications.Previous()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(applications.Items[0].ID, ShouldEqual, "123")
			})
		})
	})
	BaseURL = previousURL
}

func validateApplication(t *testing.T, application *Application) {
	So(application.ID, ShouldEqual, "123")
	So(application.Name, ShouldEqual, "foo")
	So(application.Description, ShouldEqual, "bar")
	So(application.Logo, ShouldEqual, "image.jpg")
	So(application.Keywords[0], ShouldEqual, "foo")
	So(application.Keywords[1], ShouldEqual, "bar")
	So(application.RedirectURLs[0], ShouldEqual, "http://1.com")
	So(application.RedirectURLs[1], ShouldEqual, "http://2.com")
	So(application.Contactemails[0], ShouldEqual, "john@doe.com")
	So(application.Contactemails[1], ShouldEqual, "jane@doe.com")
	So(application.Scopes[0], ShouldEqual, "scope1")
	So(application.Scopes[1], ShouldEqual, "scope2")
	So(application.SubscriptionCount, ShouldEqual, 1000)
	So(application.ClientID, ShouldEqual, "123")
	So(application.ClientSecret, ShouldEqual, "456")
}

func validateApplications(t *testing.T, applications *Applications) {
	So(applications.Items[0].ID, ShouldEqual, "123")
	So(applications.Items[0].Name, ShouldEqual, "foo")
	So(applications.Items[0].Description, ShouldEqual, "bar")
	So(applications.Items[0].Logo, ShouldEqual, "image.jpg")
	So(applications.Items[0].Keywords[0], ShouldEqual, "foo")
	So(applications.Items[0].Keywords[1], ShouldEqual, "bar")
	So(applications.Items[0].RedirectURLs[0], ShouldEqual, "http://1.com")
	So(applications.Items[0].RedirectURLs[1], ShouldEqual, "http://2.com")
	So(applications.Items[0].Contactemails[0], ShouldEqual, "john@doe.com")
	So(applications.Items[0].Contactemails[1], ShouldEqual, "jane@doe.com")
	So(applications.Items[0].Scopes[0], ShouldEqual, "scope1")
	So(applications.Items[0].Scopes[1], ShouldEqual, "scope2")
	So(applications.Items[0].SubscriptionCount, ShouldEqual, 1000)
	So(applications.Items[0].ClientID, ShouldEqual, "123")
	So(applications.Items[0].ClientSecret, ShouldEqual, "456")
}
