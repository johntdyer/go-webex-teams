package spark

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	RoomJSON  = `{"id":"123","title":"foo","members":["foo","bar"],"created":"0001-01-01T00:00:00Z"}`
	RoomsJSON = `{"items":[` + RoomJSON + `]}`
)

func TestRoomsSpec(t *testing.T) {
	Convey("Given we want to interact with Spark rooms", t, func() {
		Convey("For rooms", func() {
			Convey("Get rooms", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				InitClient("123")
				rooms := &Rooms{}
				err := rooms.Get()
				So(err, ShouldBeNil)
				So(rooms.Items[0].Title, ShouldEqual, "foo")
				So(rooms.Items[0].Members[0], ShouldEqual, "foo")
				So(rooms.Items[0].Members[1], ShouldEqual, "bar")
				BaseURL = previousURL
			})
			Convey("It should generate the proper struct from the JSON", func() {
				rooms := &Rooms{}
				err := json.Unmarshal([]byte(RoomsJSON)[:], &rooms)
				So(err, ShouldBeNil)
				So(rooms.Items[0].Title, ShouldEqual, "foo")
				So(rooms.Items[0].Members[0], ShouldEqual, "foo")
				So(rooms.Items[0].Members[1], ShouldEqual, "bar")
			})
		})
		Convey("For room", func() {
			Convey("It should generate the proper JSON message", func() {
				room := &Room{
					ID:      "123",
					Title:   "foo",
					Members: []string{"foo", "bar"},
				}
				body, err := json.Marshal(room)
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, RoomJSON)
			})
			Convey("It should generate the proper struct from the JSON", func() {
				room := &Room{}
				err := json.Unmarshal([]byte(RoomJSON)[:], room)
				So(err, ShouldBeNil)
				So(room.Title, ShouldEqual, "foo")
				So(room.Members[0], ShouldEqual, "foo")
				So(room.Members[1], ShouldEqual, "bar")
			})
			Convey("Get room", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				InitClient("123")
				room := &Room{ID: "1"}
				err := room.Get()
				So(err, ShouldBeNil)
				So(room.Title, ShouldEqual, "foo")
				So(room.Members[0], ShouldEqual, "foo")
				So(room.Members[1], ShouldEqual, "bar")
				BaseURL = previousURL
			})
			Convey("Delete room", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				InitClient("123")
				room := &Room{ID: "1"}
				err := room.Delete()
				So(err, ShouldBeNil)
				BaseURL = previousURL
			})
		})
	})
}
