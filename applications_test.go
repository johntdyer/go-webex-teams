package spark

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestApplicationsSpec(t *testing.T) {
	Convey("Given we want to interact with Spark applications", t, func() {
		Convey("For an application", func() {
			Convey("It should generate the proper JSON message", func() {
				application := &Application{
					ID:            "123",
					Name:          "foo",
					Description:   "bar",
					Logo:          "image.jpg",
					Keywords:      []string{"foo", "bar"},
					Redirecturls:  []string{"http://1.com", "http://2.com"},
					Contactemails: []string{"john@doe.com", "jane@doe.com"},
					Scopes:        []string{"scope1", "scope2"},
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
				client := NewClient("123")
				application, err := client.Application("1")
				So(err, ShouldBeNil)
				validateApplication(t, application)
				BaseURL = previousURL
			})
			Convey("Delete application", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				err := client.DeleteApplication("1")
				So(err, ShouldBeNil)
				BaseURL = previousURL
			})
		})
		Convey("For applications", func() {
			Convey("It should generate the proper JSON message", func() {
				applications := make(Applications, 1)
				applications[0].ID = "123"
				applications[0].Name = "foo"
				applications[0].Description = "bar"
				applications[0].Logo = "image.jpg"
				applications[0].Keywords = []string{"foo", "bar"}
				applications[0].Redirecturls = []string{"http://1.com", "http://2.com"}
				applications[0].Contactemails = []string{"john@doe.com", "jane@doe.com"}
				applications[0].Scopes = []string{"scope1", "scope2"}
				body, err := json.Marshal(applications)
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, ApplicationsJSON)
			})
			Convey("It should generate the proper struct from the JSON", func() {
				applications := make(Applications, 0)
				err := json.Unmarshal([]byte(ApplicationsJSON)[:], &applications)
				So(err, ShouldBeNil)
				validateApplications(t, applications)
			})
			Convey("Get applications", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				applications, err := client.Applications()
				So(err, ShouldBeNil)
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
	So(application.Created, ShouldHappenOnOrBefore, stubNow())
}

func validateApplications(t *testing.T, applications Applications) {
	So(applications[0].ID, ShouldEqual, "123")
	So(applications[0].Name, ShouldEqual, "foo")
	So(applications[0].Description, ShouldEqual, "bar")
	So(applications[0].Logo, ShouldEqual, "image.jpg")
	So(applications[0].Keywords[0], ShouldEqual, "foo")
	So(applications[0].Keywords[1], ShouldEqual, "bar")
	So(applications[0].Redirecturls[0], ShouldEqual, "http://1.com")
	So(applications[0].Redirecturls[1], ShouldEqual, "http://2.com")
	So(applications[0].Contactemails[0], ShouldEqual, "john@doe.com")
	So(applications[0].Contactemails[1], ShouldEqual, "jane@doe.com")
	So(applications[0].Scopes[0], ShouldEqual, "scope1")
	So(applications[0].Scopes[1], ShouldEqual, "scope2")
	So(applications[0].Created, ShouldHappenOnOrBefore, stubNow())
}
