package webexTeams

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	PersonJSON = `{"id":"OTZhYmMy","emails":["johnny.chang@foomail.com"],"displayName":"John Andersen","avatar":"TODO","created":"2015-10-18T07:26:16Z"}`
	PeopleJSON = `{"items":[` + PersonJSON + `]}`
)

func TestPeopleSpec(t *testing.T) {
	ts := serveHTTP(t)
	defer ts.Close()
	previousURL := BaseURL
	BaseURL = ts.URL
	InitClient(&Authorization{AccessToken: "123"})
	Convey("Given we want to interact with Teams people", t, func() {
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
				people := &People{Email: "john@doe.com"}
				result, err := people.Get()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(people.Items[0].ID, ShouldEqual, "OTZhYmMy")
				So(people.Items[0].Emails[0], ShouldEqual, "johnny.chang@foomail.com")
				So(people.Items[0].Displayname, ShouldEqual, "John Andersen")
				So(people.Items[0].Avatar, ShouldEqual, "TODO")
			})
			Convey("It should raise an error when no link cursor", func() {
				people := &People{}
				result, err := people.First()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "first cursor not available")
				result, err = people.Next()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "next cursor not available")
				result, err = people.Last()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "last cursor not available")
				result, err = people.Previous()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "previous cursor not available")
			})
			Convey("Should get our link cursor", func() {
				people := &People{}
				people.FirstURL = "/messages/first"
				result, err := people.First()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(people.Items[0].ID, ShouldEqual, "123")
				people.LastURL = "/messages/last"
				result, err = people.Last()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(people.Items[0].ID, ShouldEqual, "123")
				people.NextURL = "/messages/next"
				result, err = people.Next()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(people.Items[0].ID, ShouldEqual, "123")
				people.PreviousURL = "/messages/prev"
				result, err = people.Previous()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(people.Items[0].ID, ShouldEqual, "123")
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
				person := &Person{ID: "1"}
				result, err := person.Get()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(person.ID, ShouldEqual, "OTZhYmMy")
				So(person.Emails[0], ShouldEqual, "johnny.chang@foomail.com")
				So(person.Displayname, ShouldEqual, "John Andersen")
				So(person.Avatar, ShouldEqual, "TODO")
			})
			Convey("Get me", func() {
				person := &Person{}
				result, err := person.GetMe()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(person.ID, ShouldEqual, "OTZhYmMy")
				So(person.Emails[0], ShouldEqual, "johnny.chang@foomail.com")
				So(person.Displayname, ShouldEqual, "John Andersen")
				So(person.Avatar, ShouldEqual, "TODO")
			})
		})
	})
	BaseURL = previousURL
}
