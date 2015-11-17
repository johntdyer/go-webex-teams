package spark

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	PersonJSON = `{"id":"OTZhYmMy","emails":"johnny.chang@foomail.com","displayName":"John Andersen","avatar":"TODO","created":"2015-10-18T07:26:16-07:00"}`
	PeopleJSON = `{"items":[` + PersonJSON + `]}`
)

func TestPeopleSpec(t *testing.T) {
	Convey("Given we want to interact with Spark people", t, func() {
		Convey("For people", func() {
			Convey("Get people", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				people, err := client.People()
				So(err, ShouldBeNil)
				So(people.Items[0].ID, ShouldEqual, "OTZhYmMy")
				So(people.Items[0].Emails, ShouldEqual, "johnny.chang@foomail.com")
				So(people.Items[0].Displayname, ShouldEqual, "John Andersen")
				So(people.Items[0].Avatar, ShouldEqual, "TODO")
				So(people.Items[0].Created, ShouldHappenOnOrBefore, stubNow())
				BaseURL = previousURL
			})
			Convey("It should generate the proper struct from the JSON", func() {
				people := &People{}
				err := json.Unmarshal([]byte(PeopleJSON)[:], &people)
				So(err, ShouldBeNil)
				So(people.Items[0].ID, ShouldEqual, "OTZhYmMy")
				So(people.Items[0].Emails, ShouldEqual, "johnny.chang@foomail.com")
				So(people.Items[0].Displayname, ShouldEqual, "John Andersen")
				So(people.Items[0].Avatar, ShouldEqual, "TODO")
				So(people.Items[0].Created, ShouldHappenOnOrBefore, stubNow())
			})
		})
		Convey("For a person", func() {
			Convey("It should generate the proper JSON message", func() {
				person := &Person{
					ID:          "OTZhYmMy",
					Emails:      "johnny.chang@foomail.com",
					Displayname: "John Andersen",
					Avatar:      "TODO",
					Created:     stubNow(),
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
				So(person.Emails, ShouldEqual, "johnny.chang@foomail.com")
				So(person.Displayname, ShouldEqual, "John Andersen")
				So(person.Avatar, ShouldEqual, "TODO")
				So(person.Created, ShouldHappenOnOrBefore, stubNow())
			})
			Convey("Get person", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				person, err := client.Person("1")
				So(err, ShouldBeNil)
				So(person.ID, ShouldEqual, "OTZhYmMy")
				So(person.Emails, ShouldEqual, "johnny.chang@foomail.com")
				So(person.Displayname, ShouldEqual, "John Andersen")
				So(person.Avatar, ShouldEqual, "TODO")
				So(person.Created, ShouldHappenOnOrBefore, stubNow())
				BaseURL = previousURL
			})
		})
	})
}
