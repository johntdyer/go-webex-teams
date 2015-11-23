package spark

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	PersonJSON = `{"id":"OTZhYmMy","emails":["johnny.chang@foomail.com"],"displayName":"John Andersen","avatar":"TODO","created":"2015-10-18T07:26:16-07:00"}`
	PeopleJSON = `{"items":[` + PersonJSON + `]}`
)

func TestPeopleSpec(t *testing.T) {
	Convey("Given we want to interact with Spark people", t, func() {
		Convey("For people", func() {
			Convey("Should construct proper query strings", func() {
				Convey("Email query", func() {
					people := &People{Email: "john@doe.com"}
					query := people.buildQueryString()
					So(query, ShouldEqual, "?email=john@doe.com")
				})
				Convey("Display query", func() {
					people := &People{Displayname: "John Doe"}
					query := people.buildQueryString()
					So(query, ShouldEqual, "?displayName=John+Doe")
				})
				Convey("Email & Display query", func() {
					people := &People{
						Email:       "john@doe.com",
						Displayname: "John Doe",
					}
					query := people.buildQueryString()
					So(query, ShouldEqual, "?email=john@doe.com&displayName=John+Doe")
				})
			})
			Convey("Get people", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				InitClient("123")
				people := &People{Email: "john@doe.com"}
				err := people.Get()
				So(err, ShouldBeNil)
				So(people.Items[0].ID, ShouldEqual, "OTZhYmMy")
				So(people.Items[0].Emails[0], ShouldEqual, "johnny.chang@foomail.com")
				So(people.Items[0].Displayname, ShouldEqual, "John Andersen")
				So(people.Items[0].Avatar, ShouldEqual, "TODO")
				BaseURL = previousURL
			})
			Convey("It should generate the proper struct from the JSON", func() {
				people := &People{}
				err := json.Unmarshal([]byte(PeopleJSON)[:], &people)
				So(err, ShouldBeNil)
				So(people.Items[0].ID, ShouldEqual, "OTZhYmMy")
				So(people.Items[0].Emails[0], ShouldEqual, "johnny.chang@foomail.com")
				So(people.Items[0].Displayname, ShouldEqual, "John Andersen")
				So(people.Items[0].Avatar, ShouldEqual, "TODO")
			})
		})
		Convey("For a person", func() {
			Convey("It should generate the proper JSON message", func() {
				person := &Person{
					ID:          "OTZhYmMy",
					Emails:      []string{"johnny.chang@foomail.com"},
					Displayname: "John Andersen",
					Avatar:      "TODO",
					Created:     &CreatedTime,
				}
				body, err := json.Marshal(person)
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, PersonJSON)
			})
			Convey("It should generate the proper struct from the JSON", func() {
				person := &Person{}
				err := json.Unmarshal([]byte(PersonJSON)[:], person)
				So(err, ShouldBeNil)
				So(person.ID, ShouldEqual, "OTZhYmMy")
				So(person.Emails[0], ShouldEqual, "johnny.chang@foomail.com")
				So(person.Displayname, ShouldEqual, "John Andersen")
				So(person.Avatar, ShouldEqual, "TODO")
			})
			Convey("Get person", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				InitClient("123")
				person := &Person{ID: "1"}
				err := person.Get()
				So(err, ShouldBeNil)
				So(person.ID, ShouldEqual, "OTZhYmMy")
				So(person.Emails[0], ShouldEqual, "johnny.chang@foomail.com")
				So(person.Displayname, ShouldEqual, "John Andersen")
				So(person.Avatar, ShouldEqual, "TODO")
				BaseURL = previousURL
			})
			Convey("Get me", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				InitClient("123")
				person := &Person{}
				err := person.GetMe()
				So(err, ShouldBeNil)
				So(person.ID, ShouldEqual, "OTZhYmMy")
				So(person.Emails[0], ShouldEqual, "johnny.chang@foomail.com")
				So(person.Displayname, ShouldEqual, "John Andersen")
				So(person.Avatar, ShouldEqual, "TODO")
				BaseURL = previousURL
			})
		})
	})
}
