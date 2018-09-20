package webexTeams

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	RoomJSON         = `{"id":"123","title":"foo","sipAddress":"foo@bar.com","members":["foo","bar"],"created":"2015-10-18T07:26:16Z"}`
	RoomsJSON        = `{"items":[` + RoomJSON + `]}`
	RoomResponseJSON = `{"title":"Project Unicorn - Sprint 0","members":["john@example.com","123"]}`
)

func TestRoomsSpec(t *testing.T) {
	ts := serveHTTP(t)
	defer ts.Close()
	previousURL := BaseURL
	BaseURL = ts.URL
	InitClient(&Authorization{AccessToken: "123"})
	Convey("Given we want to interact with Teams rooms", t, func() {
		Convey("For rooms", func() {
			Convey("Get rooms", func() {
				rooms := &Rooms{}
				result, err := rooms.Get()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(rooms.Items[0].Title, ShouldEqual, "foo")
				So(rooms.Items[0].Members[0], ShouldEqual, "foo")
				So(rooms.Items[0].Members[1], ShouldEqual, "bar")
			})
			Convey("It should generate the proper struct from the JSON", func() {
				rooms := &Rooms{}
				err := json.Unmarshal([]byte(RoomsJSON)[:], &rooms)
				So(err, ShouldBeNil)
				So(rooms.Items[0].Title, ShouldEqual, "foo")
				So(rooms.Items[0].Members[0], ShouldEqual, "foo")
				So(rooms.Items[0].Members[1], ShouldEqual, "bar")
			})
			Convey("It should raise an error when no link cursor", func() {
				rooms := &Rooms{}
				result, err := rooms.First()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "first cursor not available")
				result, err = rooms.Next()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "next cursor not available")
				result, err = rooms.Last()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "last cursor not available")
				result, err = rooms.Previous()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "previous cursor not available")
			})
			Convey("Should get our link cursor", func() {
				rooms := &Rooms{}
				rooms.FirstURL = "/rooms/first"
				result, err := rooms.First()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(rooms.Items[0].ID, ShouldEqual, "123")
				rooms.LastURL = "/rooms/last"
				result, err = rooms.Last()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(rooms.Items[0].ID, ShouldEqual, "123")
				rooms.NextURL = "/rooms/next"
				result, err = rooms.Next()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(rooms.Items[0].ID, ShouldEqual, "123")
				rooms.PreviousURL = "/rooms/prev"
				result, err = rooms.Previous()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(rooms.Items[0].ID, ShouldEqual, "123")
			})
		})
		Convey("For room", func() {
			Convey("It should generate the proper JSON message", func() {
				room := &Room{
					ID:         "123",
					Title:      "foo",
					SIPAddress: "foo@bar.com",
					Members:    []string{"foo", "bar"},
					Created:    &CreatedTime,
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
				So(room.SIPAddress, ShouldEqual, "foo@bar.com")
				So(room.Members[0], ShouldEqual, "foo")
				So(room.Members[1], ShouldEqual, "bar")
			})
			Convey("Get room", func() {
				room := &Room{ID: "1"}
				result, err := room.Get()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(room.Title, ShouldEqual, "foo")
				So(room.SIPAddress, ShouldEqual, "foo@bar.com")
				So(room.Members[0], ShouldEqual, "foo")
				So(room.Members[1], ShouldEqual, "bar")
			})
			Convey("Delete room", func() {
				room := &Room{ID: "1"}
				result, err := room.Delete()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
			})
			Convey("Post room", func() {
				room := &Room{
					Title:      "Project Unicorn",
					SIPAddress: "foo@bar.com",
					Members:    []string{"john@doe.com", "456"},
				}
				result, err := room.Post()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
			})
			Convey("Put room", func() {
				room := &Room{
					ID:         "1",
					Title:      "Project Unicorn",
					SIPAddress: "foo@bar.com",
					Members:    []string{"john@doe.com", "456"},
				}
				result, err := room.Put()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
			})
		})
	})
	BaseURL = previousURL
}
