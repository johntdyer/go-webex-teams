package spark

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	MembershipJSON         = `{"id":"000","roomId":"123","personId":"456","isModerator":true,"isMonitor":true,"isLocked":true,"personEmail":"jane@doe.com","created":"2015-10-18T07:26:16Z"}`
	MembershipsJSON        = `{"items":[` + MembershipJSON + `]}`
	MembershipResponseJSON = `{"id":"1","roomId":"123","personId":"456","personEmail":"john@doe.com","isModerator":true,"isMonitor":true,"created":"2015-10-18T07:26:16Z"}`
)

func TestMembershipsSpec(t *testing.T) {
	ts := serveHTTP(t)
	defer ts.Close()
	previousURL := BaseURL
	BaseURL = ts.URL
	InitClient(&Authorization{AccessToken: "123"})
	Convey("Given we want to interact with Spark memberships", t, func() {
		Convey("Should construct proper query strings", func() {
			Convey("Personid query", func() {
				memberships := &Memberships{PersonID: "123"}
				query := memberships.buildQueryString()
				So(query, ShouldEqual, "?personId=123")
			})
			Convey("Roomid query", func() {
				memberships := &Memberships{RoomID: "123"}
				query := memberships.buildQueryString()
				So(query, ShouldEqual, "?roomId=123")
			})
			Convey("PersonEmail query", func() {
				memberships := &Memberships{PersonEmail: "123"}
				query := memberships.buildQueryString()
				So(query, ShouldEqual, "?personEmail=123")
			})
			Convey("Personid & Roomid query", func() {
				memberships := &Memberships{
					PersonID: "123",
					RoomID:   "456",
				}
				query := memberships.buildQueryString()
				So(query, ShouldEqual, "?roomId=456&personId=123")
			})
			Convey("Personid & PersonEmail query", func() {
				memberships := &Memberships{
					PersonID:    "123",
					PersonEmail: "456",
				}
				query := memberships.buildQueryString()
				So(query, ShouldEqual, "?personId=123&personEmail=456")
			})
			Convey("Personid & PersonEmail & Roomid query", func() {
				memberships := &Memberships{
					PersonID:    "123",
					PersonEmail: "456",
					RoomID:      "789",
				}
				query := memberships.buildQueryString()
				So(query, ShouldEqual, "?roomId=789&personId=123&personEmail=456")
			})
		})
		Convey("For a membership", func() {
			Convey("It should generate the proper JSON message", func() {
				membership := &Membership{
					ID:          "000",
					RoomID:      "123",
					PersonID:    "456",
					Ismoderator: true,
					Ismonitor:   true,
					Islocked:    true,
					PersonEmail: "jane@doe.com",
					Created:     &CreatedTime,
				}
				body, err := json.Marshal(membership)
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, MembershipJSON)
			})
			Convey("It should generate the proper struct from the JSON", func() {
				membership := &Membership{}
				err := json.Unmarshal([]byte(MembershipJSON)[:], membership)
				So(err, ShouldBeNil)
				So(membership.ID, ShouldEqual, "000")
				So(membership.RoomID, ShouldEqual, "123")
				So(membership.PersonID, ShouldEqual, "456")
				So(membership.Ismoderator, ShouldBeTrue)
				So(membership.Islocked, ShouldBeTrue)
				So(membership.PersonEmail, ShouldEqual, "jane@doe.com")
			})
			Convey("Get membership", func() {
				membership := &Membership{ID: "1"}
				result, err := membership.Get()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(membership.ID, ShouldEqual, "000")
				So(membership.RoomID, ShouldEqual, "123")
				So(membership.PersonID, ShouldEqual, "456")
				So(membership.Ismoderator, ShouldEqual, true)
			})
			Convey("Delete membership", func() {
				membership := &Membership{ID: "1"}
				result, err := membership.Delete()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
			})
			Convey("Post membership", func() {
				membership := &Membership{
					RoomID:      "123",
					PersonID:    "456",
					PersonEmail: "john@doe.com",
					Ismoderator: true,
				}
				result, err := membership.Post()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
			})
			Convey("Put membership", func() {
				membership := &Membership{
					ID:          "1",
					RoomID:      "123",
					PersonID:    "456",
					PersonEmail: "john@doe.com",
					Ismoderator: true,
				}
				result, err := membership.Put()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
			})
		})
		Convey("For memberships", func() {
			Convey("It should generate the proper struct from the JSON", func() {
				memberships := &Memberships{}
				err := json.Unmarshal([]byte(MembershipsJSON)[:], &memberships)
				So(err, ShouldBeNil)
				So(memberships.Items[0].ID, ShouldEqual, "000")
				So(memberships.Items[0].RoomID, ShouldEqual, "123")
				So(memberships.Items[0].PersonID, ShouldEqual, "456")
				So(memberships.Items[0].Ismoderator, ShouldEqual, true)
				So(memberships.Items[0].Ismonitor, ShouldEqual, true)
				So(memberships.Items[0].Islocked, ShouldEqual, true)
			})
			Convey("Get memberships", func() {
				memberships := &Memberships{RoomID: "123"}
				result, err := memberships.Get()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(memberships.Items[0].ID, ShouldEqual, "000")
				So(memberships.Items[0].RoomID, ShouldEqual, "123")
				So(memberships.Items[0].PersonID, ShouldEqual, "456")
				So(memberships.Items[0].Ismoderator, ShouldEqual, true)
				So(memberships.Items[0].Ismonitor, ShouldEqual, true)
				So(memberships.Items[0].Islocked, ShouldEqual, true)
			})
			Convey("It should raise an error when no link cursor", func() {
				memberships := &Memberships{}
				result, err := memberships.First()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "first cursor not available")
				result, err = memberships.Next()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "next cursor not available")
				result, err = memberships.Last()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "last cursor not available")
				result, err = memberships.Previous()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "previous cursor not available")
			})
			Convey("Should get our link cursor", func() {
				memberships := &Memberships{}
				memberships.FirstURL = "/messages/first"
				result, err := memberships.First()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(memberships.Items[0].ID, ShouldEqual, "123")
				memberships.LastURL = "/messages/last"
				result, err = memberships.Last()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(memberships.Items[0].ID, ShouldEqual, "123")
				memberships.NextURL = "/messages/next"
				result, err = memberships.Next()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(memberships.Items[0].ID, ShouldEqual, "123")
				memberships.PreviousURL = "/messages/prev"
				result, err = memberships.Previous()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(memberships.Items[0].ID, ShouldEqual, "123")
			})
		})
	})
	BaseURL = previousURL
}
