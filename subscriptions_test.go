package spark

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	SubscriptionJSON  = `{"id":"000","personId":"123","applicationId":"456","applicationName":"foo","created":"2015-10-18T07:26:16-07:00"}`
	SubscriptionsJSON = `{"items":[` + SubscriptionJSON + `]}`
)

func TestSubscriptionsSpec(t *testing.T) {
	ts := serveHTTP(t)
	defer ts.Close()
	previousURL := BaseURL
	BaseURL = ts.URL
	InitClient("123")
	Convey("Given we want to interact with Spark subscriptions", t, func() {
		Convey("For a subscription", func() {
			Convey("It should generate the proper JSON message", func() {
				subscription := &Subscription{
					ID:              "000",
					Personid:        "123",
					Applicationid:   "456",
					Applicationname: "foo",
					Created:         &CreatedTime,
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
			})
			Convey("Get subscription", func() {
				subscription := &Subscription{ID: "1"}
				err := subscription.Get()
				So(err, ShouldBeNil)
				So(subscription.ID, ShouldEqual, "000")
				So(subscription.Personid, ShouldEqual, "123")
				So(subscription.Applicationid, ShouldEqual, "456")
				So(subscription.Applicationname, ShouldEqual, "foo")
			})
			Convey("Delete subscription", func() {
				subscription := &Subscription{ID: "1"}
				err := subscription.Delete()
				So(err, ShouldBeNil)
				BaseURL = previousURL
			})
		})
		Convey("For subscriptions", func() {
			Convey("It should generate the proper struct from the JSON", func() {
				subscriptions := &Subscriptions{}
				err := json.Unmarshal([]byte(SubscriptionsJSON)[:], &subscriptions)
				So(err, ShouldBeNil)
				So(subscriptions.Items[0].ID, ShouldEqual, "000")
				So(subscriptions.Items[0].Personid, ShouldEqual, "123")
				So(subscriptions.Items[0].Applicationid, ShouldEqual, "456")
				So(subscriptions.Items[0].Applicationname, ShouldEqual, "foo")
			})
			Convey("Get subscriptions", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				BaseURL = ts.URL
				InitClient("123")
				subscriptions := &Subscriptions{}
				err := subscriptions.Get()
				So(err, ShouldBeNil)
				So(subscriptions.Items[0].ID, ShouldEqual, "000")
				So(subscriptions.Items[0].Personid, ShouldEqual, "123")
				So(subscriptions.Items[0].Applicationid, ShouldEqual, "456")
				So(subscriptions.Items[0].Applicationname, ShouldEqual, "foo")
			})
		})
	})
	BaseURL = previousURL
}
