package spark

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSubscriptionsSpec(t *testing.T) {
	Convey("Given we want to interact with Spark subscriptions", t, func() {
		Convey("For a subscription", func() {
			Convey("It should generate the proper JSON message", func() {
				subscription := &Subscription{
					ID:              "000",
					Personid:        "123",
					Applicationid:   "456",
					Applicationname: "foo",
					Created:         stubNow(),
				}
				body, err := json.Marshal(subscription)
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, SubscriptionJSON)
			})
			Convey("It should generate the proper struct from the JSON", func() {
				subscription := &Subscription{}
				err := json.Unmarshal([]byte(SubscriptionJSON)[:], subscription)
				So(err, ShouldBeNil)
				So(subscription.ID, ShouldEqual, "000")
				So(subscription.Personid, ShouldEqual, "123")
				So(subscription.Applicationid, ShouldEqual, "456")
				So(subscription.Applicationname, ShouldEqual, "foo")
				So(subscription.Created, ShouldHappenOnOrBefore, stubNow())
			})
			Convey("Get subscription", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				subscription, err := client.Subscription("1")
				So(err, ShouldBeNil)
				So(subscription.ID, ShouldEqual, "000")
				So(subscription.Personid, ShouldEqual, "123")
				So(subscription.Applicationid, ShouldEqual, "456")
				So(subscription.Applicationname, ShouldEqual, "foo")
				So(subscription.Created, ShouldHappenOnOrBefore, stubNow())
				BaseURL = previousURL
			})
			Convey("Delete subscription", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				err := client.DeleteSubscription("1")
				So(err, ShouldBeNil)
				BaseURL = previousURL
			})
		})
		Convey("For subscriptions", func() {
			Convey("It should generate the proper JSON message", func() {
				subscriptions := make(Subscriptions, 1)
				subscriptions[0].ID = "000"
				subscriptions[0].Personid = "123"
				subscriptions[0].Applicationid = "456"
				subscriptions[0].Applicationname = "foo"
				subscriptions[0].Created = stubNow()
				body, err := json.Marshal(subscriptions)
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, SubscriptionsJSON)
			})
			Convey("It should generate the proper struct from the JSON", func() {
				subscriptions := make(Subscriptions, 0)
				err := json.Unmarshal([]byte(SubscriptionsJSON)[:], &subscriptions)
				So(err, ShouldBeNil)
				So(subscriptions[0].ID, ShouldEqual, "000")
				So(subscriptions[0].Personid, ShouldEqual, "123")
				So(subscriptions[0].Applicationid, ShouldEqual, "456")
				So(subscriptions[0].Applicationname, ShouldEqual, "foo")
				So(subscriptions[0].Created, ShouldHappenOnOrBefore, stubNow())
			})
			Convey("Get subscriptions", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				subscriptions, err := client.Subscriptions()
				So(err, ShouldBeNil)
				So(subscriptions[0].ID, ShouldEqual, "000")
				So(subscriptions[0].Personid, ShouldEqual, "123")
				So(subscriptions[0].Applicationid, ShouldEqual, "456")
				So(subscriptions[0].Applicationname, ShouldEqual, "foo")
				So(subscriptions[0].Created, ShouldHappenOnOrBefore, stubNow())
				BaseURL = previousURL
			})
		})
	})
}
