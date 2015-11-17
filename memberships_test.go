package spark

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	MembershipJSON         = `{"id":"000","roomId":"123","personId":"456","isModerator":true,"isMonitor":true,"isLocked":true,"personEmail":"jane@doe.com","created":"0001-01-01T00:00:00Z"}`
	MembershipsJSON        = `{"items":[` + MembershipJSON + `]}`
	MembershipResponseJSON = `{"id":"1","roomId":"123","personId":"456","personEmail":"john@doe.com","isModerator":true,"isMonitor":true,"created":"0001-01-01T00:00:00Z"}`
)

func TestMembershipsSpec(t *testing.T) {
	ts := serveHTTP(t)
	defer ts.Close()
	previousURL := BaseURL
	BaseURL = ts.URL
	InitClient("123")
	Convey("Given we want to interact with Spark memberships", t, func() {
		Convey("For a membership", func() {
			Convey("It should generate the proper JSON message", func() {
				membership := &Membership{
					ID:          "000",
					Roomid:      "123",
					Personid:    "456",
					Ismoderator: true,
					Ismonitor:   true,
					Islocked:    true,
					PersonEmail: "jane@doe.com",
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
				So(membership.Roomid, ShouldEqual, "123")
				So(membership.Personid, ShouldEqual, "456")
				So(membership.Ismoderator, ShouldBeTrue)
				So(membership.Islocked, ShouldBeTrue)
				So(membership.PersonEmail, ShouldEqual, "jane@doe.com")
			})
			Convey("Get membership", func() {
				membership := &Membership{ID: "1"}
				err := membership.Get()
				So(err, ShouldBeNil)
				So(membership.ID, ShouldEqual, "000")
				So(membership.Roomid, ShouldEqual, "123")
				So(membership.Personid, ShouldEqual, "456")
				So(membership.Ismoderator, ShouldEqual, true)
			})
			Convey("Delete membership", func() {
				membership := &Membership{ID: "1"}
				err := membership.Delete()
				So(err, ShouldBeNil)
			})
			Convey("Post membership", func() {
				membership := &Membership{
					Roomid:      "123",
					Personid:    "456",
					PersonEmail: "john@doe.com",
					Ismoderator: true,
				}
				err := membership.Post()
				So(err, ShouldBeNil)
			})
			Convey("Put membership", func() {
				membership := &Membership{
					ID:          "1",
					Roomid:      "123",
					Personid:    "456",
					PersonEmail: "john@doe.com",
					Ismoderator: true,
				}
				err := membership.Put()
				So(err, ShouldBeNil)
			})
		})
		Convey("For memberships", func() {
			Convey("It should generate the proper struct from the JSON", func() {
				memberships := &Memberships{}
				err := json.Unmarshal([]byte(MembershipsJSON)[:], &memberships)
				So(err, ShouldBeNil)
				So(memberships.Items[0].ID, ShouldEqual, "000")
				So(memberships.Items[0].Roomid, ShouldEqual, "123")
				So(memberships.Items[0].Personid, ShouldEqual, "456")
				So(memberships.Items[0].Ismoderator, ShouldEqual, true)
				So(memberships.Items[0].Ismonitor, ShouldEqual, true)
				So(memberships.Items[0].Islocked, ShouldEqual, true)
				So(memberships.Items[0].Created, ShouldHappenOnOrBefore, stubNow())
			})
			Convey("Get memberships", func() {
				memberships := &Memberships{}
				err := memberships.Get()
				So(err, ShouldBeNil)
				So(memberships.Items[0].ID, ShouldEqual, "000")
				So(memberships.Items[0].Roomid, ShouldEqual, "123")
				So(memberships.Items[0].Personid, ShouldEqual, "456")
				So(memberships.Items[0].Ismoderator, ShouldEqual, true)
				So(memberships.Items[0].Ismonitor, ShouldEqual, true)
				So(memberships.Items[0].Islocked, ShouldEqual, true)
				So(memberships.Items[0].Created, ShouldHappenOnOrBefore, stubNow())
			})
		})
	})
	BaseURL = previousURL
}
