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
				subscriptions := &Subscriptions{PersonID: "123"}
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
					PersonID: "123",
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
					PersonID:        "123",
					ApplicationID:   "456",
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
				So(subscription.PersonID, ShouldEqual, "123")
				So(subscription.ApplicationID, ShouldEqual, "456")
				So(subscription.Applicationname, ShouldEqual, "foo")
			})
			Convey("Get subscription", func() {
				subscription := &Subscription{ID: "1"}
				result, err := subscription.Get()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(subscription.ID, ShouldEqual, "000")
				So(subscription.PersonID, ShouldEqual, "123")
				So(subscription.ApplicationID, ShouldEqual, "456")
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
				So(subscriptions.Items[0].PersonID, ShouldEqual, "123")
				So(subscriptions.Items[0].ApplicationID, ShouldEqual, "456")
				So(subscriptions.Items[0].Applicationname, ShouldEqual, "foo")
			})
			Convey("Get subscriptions", func() {
				subscriptions := &Subscriptions{PersonID: "123"}
				result, err := subscriptions.Get()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(subscriptions.Items[0].ID, ShouldEqual, "000")
				So(subscriptions.Items[0].PersonID, ShouldEqual, "123")
				So(subscriptions.Items[0].ApplicationID, ShouldEqual, "456")
				So(subscriptions.Items[0].Applicationname, ShouldEqual, "foo")
			})
			Convey("It should raise an error when no link cursor", func() {
				subscriptions := &Subscriptions{}
				result, err := subscriptions.First()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "first cursor not available")
				result, err = subscriptions.Next()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "next cursor not available")
				result, err = subscriptions.Last()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "last cursor not available")
				result, err = subscriptions.Previous()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "previous cursor not available")
			})
			Convey("Should get our link cursor", func() {
				subscriptions := &Subscriptions{}
				subscriptions.FirstURL = "/subscriptions/first"
				result, err := subscriptions.First()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(subscriptions.Items[0].ID, ShouldEqual, "000")
				subscriptions.LastURL = "/subscriptions/last"
				result, err = subscriptions.Last()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(subscriptions.Items[0].ID, ShouldEqual, "000")
				subscriptions.NextURL = "/subscriptions/next"
				result, err = subscriptions.Next()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(subscriptions.Items[0].ID, ShouldEqual, "000")
				subscriptions.PreviousURL = "/subscriptions/prev"
				result, err = subscriptions.Previous()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(subscriptions.Items[0].ID, ShouldEqual, "000")
			})
		})
	})
	BaseURL = previousURL
}
