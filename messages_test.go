package spark

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	MessageJSON         = `{"id":"123","roomId":"456","text":"foo","files":["image1.jpg","image2.jpg"],"created":"2015-10-18T07:26:16-07:00"}`
	MessagesJSON        = `{"items":[` + MessageJSON + `]}`
	MessageResponseJSON = `{"id":"123","personId":"456","personEmail":"matt@example.com","roomId":"789","text":"PROJECT UPDATE - A new project project plan has been published on Box","files":["http://www.example.com/images/media.png"],"created":"2015-10-18T14:26:16+00:00"}`
)

func TestMessagesSpec(t *testing.T) {
	ts := serveHTTP(t)
	defer ts.Close()
	previousURL := BaseURL
	BaseURL = ts.URL
	InitClient("123")
	Convey("Given we want to interact with Spark memberships", t, func() {
		Convey("For messages", func() {
			Convey("Get messages", func() {
				messages := &Messages{Roomid: "1"}
				err := messages.Get()
				So(err, ShouldBeNil)
				So(messages.Items[0].Roomid, ShouldEqual, "456")
				So(messages.Items[0].Text, ShouldEqual, "foo")
				So(messages.Items[0].Files[0], ShouldEqual, "image1.jpg")
				So(messages.Items[0].Files[1], ShouldEqual, "image2.jpg")
			})
			Convey("It should generate the proper struct from the JSON", func() {
				messages := &Messages{}
				err := json.Unmarshal([]byte(MessagesJSON)[:], &messages)
				So(err, ShouldBeNil)
				So(messages.Items[0].Roomid, ShouldEqual, "456")
				So(messages.Items[0].Text, ShouldEqual, "foo")
				So(messages.Items[0].Files[0], ShouldEqual, "image1.jpg")
				So(messages.Items[0].Files[1], ShouldEqual, "image2.jpg")
			})
		})
		Convey("For a message", func() {
			Convey("It should generate the proper JSON message", func() {
				message := &Message{
					ID:      "123",
					Roomid:  "456",
					Text:    "foo",
					Files:   []string{"image1.jpg", "image2.jpg"},
					Created: &CreatedTime,
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
				So(message.Files[0], ShouldEqual, "image1.jpg")
				So(message.Files[1], ShouldEqual, "image2.jpg")
			})
			Convey("Get message", func() {
				message := &Message{ID: "1"}
				err := message.Get()
				So(err, ShouldBeNil)
				So(message.Roomid, ShouldEqual, "456")
				So(message.Text, ShouldEqual, "foo")
				So(message.Files[0], ShouldEqual, "image1.jpg")
				So(message.Files[1], ShouldEqual, "image2.jpg")
			})
			Convey("Delete message", func() {
				message := &Message{ID: "1"}
				err := message.Delete()
				So(err, ShouldBeNil)
			})
			Convey("Post message", func() {
				message := &Message{
					Roomid: "123",
					Text:   "foobar",
					Files:  []string{"foo.txt", "bar.txt"},
				}
				err := message.Post()
				So(err, ShouldBeNil)
			})
		})
	})
	BaseURL = previousURL
}
