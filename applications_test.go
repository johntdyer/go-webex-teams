package spark

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	ApplicationJSON         = `{"id":"123","name":"foo","description":"bar","logo":"image.jpg","keywords":["foo","bar"],"contactEmails":["john@doe.com","jane@doe.com"],"redirectUrls":["http://1.com","http://2.com"],"scopes":["scope1","scope2"],"subscriptionCount":1000,"clientId":"123","clientSecret":"456","created":"2015-10-18T07:26:16-07:00"}`
	ApplicationsJSON        = `{"items":[` + ApplicationJSON + `]}`
	ApplicationResponseJSON = `{"id":"123","name":"Out of Office Assistant","description":"Does awesome things","logo":"logo.jpg","keywords":["foo","bar"],"contactEmails":["bob@foo.com","alice@bar.org"],"redirectUrls":["http://myapp.com/verify","http://myapp.fr/verify"],"scopes":["foo","bar"],"subscriptionCount":1000,"clientId":"456","clientSecret":"secret","created":"2015-10-18T07:26:16-07:00"}`
)

func TestApplicationsSpec(t *testing.T) {
	ts := serveHTTP(t)
	defer ts.Close()
	previousURL := BaseURL
	BaseURL = ts.URL
	InitClient("123")
	Convey("Given we want to interact with Spark applications", t, func() {
		Convey("For an application", func() {
			Convey("It should generate the proper JSON message", func() {
				application := &Application{
					ID:                "123",
					Name:              "foo",
					Description:       "bar",
					Logo:              "image.jpg",
					Keywords:          []string{"foo", "bar"},
					Redirecturls:      []string{"http://1.com", "http://2.com"},
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
				err := application.Get()
				So(err, ShouldBeNil)
				validateApplication(t, application)
			})
			Convey("Delete application", func() {
				application := &Application{ID: "1"}
				err := application.Delete()
				So(err, ShouldBeNil)
			})
			Convey("Post application", func() {
				application := &Application{
					Name:          "Out of Office Assistant",
					Description:   "Does awesome things",
					Logo:          "logo.jpg",
					Keywords:      []string{"foo", "bar"},
					Contactemails: []string{"bob@foo.com", "alice@bar.org"},
					Redirecturls:  []string{"http://myapp.com/verify", "http://myapp.fr/verify"},
					Scopes:        []string{"foo", "bar"},
				}
				err := application.Post()
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
					Redirecturls:  []string{"http://myapp.com/verify", "http://myapp.fr/verify"},
					Scopes:        []string{"foo", "bar"},
				}
				err := application.Put()
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
				err := applications.Get()
				So(err, ShouldBeNil)
				validateApplications(t, applications)
			})
			Convey("It should raise an error when no link cursor", func() {
				applications := &Applications{}
				err := applications.First()
				So(err.Error(), ShouldEqual, "first cursor not available")
				err = applications.Next()
				So(err.Error(), ShouldEqual, "next cursor not available")
				err = applications.Last()
				So(err.Error(), ShouldEqual, "last cursor not available")
				err = applications.Previous()
				So(err.Error(), ShouldEqual, "previous cursor not available")
			})
			Convey("Should get our link cursor", func() {
				applications := &Applications{}
				applications.FirstURL = "/applications/first"
				err := applications.First()
				So(err, ShouldBeNil)
				So(applications.Items[0].ID, ShouldEqual, "123")
				applications.LastURL = "/applications/last"
				err = applications.Last()
				So(err, ShouldBeNil)
				So(applications.Items[0].ID, ShouldEqual, "123")
				applications.NextURL = "/applications/next"
				err = applications.Next()
				So(err, ShouldBeNil)
				So(applications.Items[0].ID, ShouldEqual, "123")
				applications.PreviousURL = "/applications/prev"
				err = applications.Previous()
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
	So(application.Redirecturls[0], ShouldEqual, "http://1.com")
	So(application.Redirecturls[1], ShouldEqual, "http://2.com")
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
	So(applications.Items[0].Redirecturls[0], ShouldEqual, "http://1.com")
	So(applications.Items[0].Redirecturls[1], ShouldEqual, "http://2.com")
	So(applications.Items[0].Contactemails[0], ShouldEqual, "john@doe.com")
	So(applications.Items[0].Contactemails[1], ShouldEqual, "jane@doe.com")
	So(applications.Items[0].Scopes[0], ShouldEqual, "scope1")
	So(applications.Items[0].Scopes[1], ShouldEqual, "scope2")
	So(applications.Items[0].SubscriptionCount, ShouldEqual, 1000)
	So(applications.Items[0].ClientID, ShouldEqual, "123")
	So(applications.Items[0].ClientSecret, ShouldEqual, "456")
}
