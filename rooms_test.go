package spark

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRoomsSpec(t *testing.T) {
	Convey("Given we want to interact with Spark rooms", t, func() {
		Convey("For rooms", func() {
			Convey("Get rooms", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				rooms, err := client.Rooms()
				So(err, ShouldBeNil)
				So(rooms[0].Title, ShouldEqual, "foo")
				BaseURL = previousURL
			})
			Convey("It should generate the proper JSON message", func() {
				rooms := make(Rooms, 1)
				rooms[0].ID = "123"
				rooms[0].Title = "foo"
				body, err := json.Marshal(rooms)
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, RoomsJSON)
			})
			Convey("It should generate the proper struct from the JSON", func() {
				rooms := make(Rooms, 0)
				err := json.Unmarshal([]byte(RoomsJSON)[:], &rooms)
				So(err, ShouldBeNil)
				So(rooms[0].Title, ShouldEqual, "foo")
			})
		})
		Convey("For room", func() {
			Convey("It should generate the proper JSON message", func() {
				room := &Room{
					ID:    "123",
					Title: "foo",
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
			})
			Convey("Get room", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				room, err := client.Room("1")
				So(err, ShouldBeNil)
				So(room.Title, ShouldEqual, "foo")
				BaseURL = previousURL
			})
			Convey("Delete room", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				err := client.DeleteRoom("1")
				So(err, ShouldBeNil)
				BaseURL = previousURL
			})
		})
	})
}
