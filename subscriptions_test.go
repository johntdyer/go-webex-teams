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
		Convey("Should construct proper query strings", func() {
			Convey("Personid query", func() {
				subscriptions := &Subscriptions{Personid: "123"}
				query := subscriptions.buildQueryString()
				So(query, ShouldEqual, "?personId=123")
			})
			Convey("Type query", func() {
				subscriptions := &Subscriptions{Type: "string"}
				query := subscriptions.buildQueryString()
				So(query, ShouldEqual, "?type=string")
			})
			Convey("Personid & Type query", func() {
				subscriptions := &Subscriptions{
					Personid: "123",
					Type:     "string",
				}
				query := subscriptions.buildQueryString()
				So(query, ShouldEqual, "?personId=123&type=string")
			})
		})
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
				result, err := subscription.Delete()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
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
				subscriptions := &Subscriptions{Personid: "123"}
				err := subscriptions.Get()
				So(err, ShouldBeNil)
				So(subscriptions.Items[0].ID, ShouldEqual, "000")
				So(subscriptions.Items[0].Personid, ShouldEqual, "123")
				So(subscriptions.Items[0].Applicationid, ShouldEqual, "456")
				So(subscriptions.Items[0].Applicationname, ShouldEqual, "foo")
			})
			Convey("It should raise an error when no link cursor", func() {
				subscriptions := &Subscriptions{}
				err := subscriptions.First()
				So(err.Error(), ShouldEqual, "first cursor not available")
				err = subscriptions.Next()
				So(err.Error(), ShouldEqual, "next cursor not available")
				err = subscriptions.Last()
				So(err.Error(), ShouldEqual, "last cursor not available")
				err = subscriptions.Previous()
				So(err.Error(), ShouldEqual, "previous cursor not available")
			})
			Convey("Should get our link cursor", func() {
				subscriptions := &Subscriptions{}
				subscriptions.FirstURL = "/subscriptions/first"
				err := subscriptions.First()
				So(err, ShouldBeNil)
				So(subscriptions.Items[0].ID, ShouldEqual, "000")
				subscriptions.LastURL = "/subscriptions/last"
				err = subscriptions.Last()
				So(err, ShouldBeNil)
				So(subscriptions.Items[0].ID, ShouldEqual, "000")
				subscriptions.NextURL = "/subscriptions/next"
				err = subscriptions.Next()
				So(err, ShouldBeNil)
				So(subscriptions.Items[0].ID, ShouldEqual, "000")
				subscriptions.PreviousURL = "/subscriptions/prev"
				err = subscriptions.Previous()
				So(err, ShouldBeNil)
				So(subscriptions.Items[0].ID, ShouldEqual, "000")
			})
		})
	})
	BaseURL = previousURL
}
