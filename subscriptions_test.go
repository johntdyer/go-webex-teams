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
			Convey("It should generate the proper struct from the JSON", func() {
				subscriptions := &Subscriptions{}
				err := json.Unmarshal([]byte(SubscriptionsJSON)[:], &subscriptions)
				So(err, ShouldBeNil)
				So(subscriptions.Items[0].ID, ShouldEqual, "000")
				So(subscriptions.Items[0].Personid, ShouldEqual, "123")
				So(subscriptions.Items[0].Applicationid, ShouldEqual, "456")
				So(subscriptions.Items[0].Applicationname, ShouldEqual, "foo")
				So(subscriptions.Items[0].Created, ShouldHappenOnOrBefore, stubNow())
			})
			Convey("Get subscriptions", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				client := NewClient("123")
				subscriptions, err := client.Subscriptions()
				So(err, ShouldBeNil)
				So(subscriptions.Items[0].ID, ShouldEqual, "000")
				So(subscriptions.Items[0].Personid, ShouldEqual, "123")
				So(subscriptions.Items[0].Applicationid, ShouldEqual, "456")
				So(subscriptions.Items[0].Applicationname, ShouldEqual, "foo")
				So(subscriptions.Items[0].Created, ShouldHappenOnOrBefore, stubNow())
				BaseURL = previousURL
			})
		})
	})
}
