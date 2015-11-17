package spark

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	MessageJSON  = `{"id":"123","roomId":"456","text":"foo","file":"image.jpg","created":"2015-10-18T07:26:16-07:00"}`
	MessagesJSON = `{"items":[` + MessageJSON + `]}`
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
				So(messages.Items[0].Roomid, ShouldEqual, "456")
				So(messages.Items[0].Text, ShouldEqual, "foo")
				So(messages.Items[0].File, ShouldEqual, "image.jpg")
				So(messages.Items[0].Created, ShouldHappenOnOrBefore, stubNow())
				BaseURL = previousURL
			})
			Convey("It should generate the proper struct from the JSON", func() {
				messages := &Messages{}
				err := json.Unmarshal([]byte(MessagesJSON)[:], &messages)
				So(err, ShouldBeNil)
				So(messages.Items[0].Roomid, ShouldEqual, "456")
				So(messages.Items[0].Text, ShouldEqual, "foo")
				So(messages.Items[0].File, ShouldEqual, "image.jpg")
				So(messages.Items[0].Created, ShouldHappenOnOrBefore, stubNow())
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
