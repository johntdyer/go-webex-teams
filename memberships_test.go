package spark

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMembershipsSpec(t *testing.T) {
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
					Email:       "jane@doe.com",
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
				So(membership.Email, ShouldEqual, "jane@doe.com")
			})
			Convey("Get membership", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				membership, err := client.Membership("1")
				So(err, ShouldBeNil)
				So(membership.ID, ShouldEqual, "000")
				So(membership.Roomid, ShouldEqual, "123")
				So(membership.Personid, ShouldEqual, "456")
				So(membership.Ismoderator, ShouldEqual, true)
				BaseURL = previousURL
			})
			Convey("Delete membership", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				err := client.DeleteMembership("1")
				So(err, ShouldBeNil)
				BaseURL = previousURL
			})
		})
		Convey("For memberships", func() {
			Convey("It should generate the proper JSON message", func() {
				memberships := make(Memberships, 1)
				memberships[0].ID = "000"
				memberships[0].Roomid = "123"
				memberships[0].Personid = "456"
				memberships[0].Ismoderator = true
				memberships[0].Ismonitor = true
				memberships[0].Islocked = true
				memberships[0].Email = "jane@doe.com"
				body, err := json.Marshal(memberships)
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, MembershipsJSON)
			})
			Convey("It should generate the proper struct from the JSON", func() {
				memberships := make(Memberships, 0)
				err := json.Unmarshal([]byte(MembershipsJSON)[:], &memberships)
				So(err, ShouldBeNil)
				So(memberships[0].ID, ShouldEqual, "000")
				So(memberships[0].Roomid, ShouldEqual, "123")
				So(memberships[0].Personid, ShouldEqual, "456")
				So(memberships[0].Ismoderator, ShouldEqual, true)
				So(memberships[0].Ismonitor, ShouldEqual, true)
				So(memberships[0].Islocked, ShouldEqual, true)
				So(memberships[0].Created, ShouldHappenOnOrBefore, stubNow())
			})
			Convey("Get memberships", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				memberships, err := client.Memberships()
				So(err, ShouldBeNil)
				So(memberships[0].ID, ShouldEqual, "000")
				So(memberships[0].Roomid, ShouldEqual, "123")
				So(memberships[0].Personid, ShouldEqual, "456")
				So(memberships[0].Ismoderator, ShouldEqual, true)
				So(memberships[0].Ismonitor, ShouldEqual, true)
				So(memberships[0].Islocked, ShouldEqual, true)
				So(memberships[0].Created, ShouldHappenOnOrBefore, stubNow())
				BaseURL = previousURL
			})
		})
	})
}
