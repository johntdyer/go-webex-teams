package spark

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMessagesSpec(t *testing.T) {
	Convey("Given we want to interact with Spark memberships", t, func() {
		Convey("For messages", func() {
			Convey("Get messages", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				messages, err := client.Messages()
				So(err, ShouldBeNil)
				So(messages[0].Roomid, ShouldEqual, "456")
				So(messages[0].Text, ShouldEqual, "foo")
				So(messages[0].File, ShouldEqual, "image.jpg")
				So(messages[0].Created, ShouldHappenOnOrBefore, stubNow())
				BaseURL = previousURL
			})
			Convey("It should generate the proper JSON message", func() {
				messages := make(Messages, 1)
				messages[0].ID = "123"
				messages[0].Roomid = "456"
				messages[0].Text = "foo"
				messages[0].File = "image.jpg"
				messages[0].Created = stubNow()
				body, err := json.Marshal(messages)
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, MessagesJSON)
			})
			Convey("It should generate the proper struct from the JSON", func() {
				messages := make(Messages, 0)
				err := json.Unmarshal([]byte(MessagesJSON)[:], &messages)
				So(err, ShouldBeNil)
				So(messages[0].Roomid, ShouldEqual, "456")
				So(messages[0].Text, ShouldEqual, "foo")
				So(messages[0].File, ShouldEqual, "image.jpg")
				So(messages[0].Created, ShouldHappenOnOrBefore, stubNow())
			})
		})
		Convey("For a message", func() {
			Convey("It should generate the proper JSON message", func() {
				message := &Message{
					ID:      "123",
					Roomid:  "456",
					Text:    "foo",
					File:    "image.jpg",
					Created: stubNow(),
				}
				body, err := json.Marshal(message)
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, MessageJSON)
			})
			Convey("It should generate the proper struct from the JSON", func() {
				message := &Message{}
				err := json.Unmarshal([]byte(MessageJSON)[:], message)
				So(err, ShouldBeNil)
				So(message.Roomid, ShouldEqual, "456")
				So(message.Text, ShouldEqual, "foo")
				So(message.File, ShouldEqual, "image.jpg")
			})
			Convey("Get message", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				message, err := client.Message("1")
				So(err, ShouldBeNil)
				So(message.Roomid, ShouldEqual, "456")
				So(message.Text, ShouldEqual, "foo")
				So(message.File, ShouldEqual, "image.jpg")
				So(message.Created, ShouldHappenOnOrBefore, stubNow())
				BaseURL = previousURL
			})
			Convey("Delete message", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				err := client.DeleteMessage("1")
				So(err, ShouldBeNil)
				BaseURL = previousURL
			})
		})
	})
}
