package spark

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	ApplicationJSON  = `{"id":"123","name":"foo","description":"bar","logo":"image.jpg","keywords":["foo","bar"],"contactEmails":["john@doe.com","jane@doe.com"],"redirectUrls":["http://1.com","http://2.com"],"scopes":["scope1","scope2"],"subscriptionCount":1000,"clientId":"123","clientSecret":"456","created":"0001-01-01T00:00:00Z"}`
	ApplicationsJSON = `{"items":[` + ApplicationJSON + `]}`
)

func TestApplicationsSpec(t *testing.T) {
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
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				InitClient("123")
				application := &Application{ID: "1"}
				err := application.Get()
				So(err, ShouldBeNil)
				validateApplication(t, application)
				BaseURL = previousURL
			})
			Convey("Delete application", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				InitClient("123")
				application := &Application{ID: "1"}
				err := application.Delete()
				So(err, ShouldBeNil)
				BaseURL = previousURL
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
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				InitClient("123")
				applications := &Applications{}
				err := applications.Get()
				So(err, ShouldBeNil)
				fmt.Println("+++++")
				fmt.Println(applications)
				validateApplications(t, applications)
				BaseURL = previousURL
			})
		})
	})
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
	So(application.Created, ShouldHappenOnOrBefore, stubNow())
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
	So(applications.Items[0].Created, ShouldHappenOnOrBefore, stubNow())
}
