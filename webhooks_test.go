package spark

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	WebhookJSON  = `{"id":"123","resource":"messages","event":"created","filter":"roomId=456","targetUrl":"https://example.com/mywebhook","name":"My Awesome Webhook","created":"2015-10-18T07:26:16-07:00"}`
	WebhooksJSON = `{"items":[` + WebhookJSON + `]}`
)

func TestWebhooksSpec(t *testing.T) {
	Convey("Given we want to interact with Spark webhooks", t, func() {
		Convey("For a webhook", func() {
			Convey("It should generate the proper JSON message", func() {
				webhook := &Webhook{
					ID:        "123",
					Resource:  "messages",
					Event:     "created",
					Filter:    "roomId=456",
					Targeturl: "https://example.com/mywebhook",
					Name:      "My Awesome Webhook",
					Created:   stubNow(),
				}
				body, err := json.Marshal(webhook)
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, WebhookJSON)
			})
			Convey("It should generate the proper struct from the JSON", func() {
				webhook := &Webhook{}
				err := json.Unmarshal([]byte(WebhookJSON)[:], webhook)
				So(err, ShouldBeNil)
				So(webhook.ID, ShouldEqual, "123")
				So(webhook.Resource, ShouldEqual, "messages")
				So(webhook.Event, ShouldEqual, "created")
				So(webhook.Filter, ShouldEqual, "roomId=456")
				So(webhook.Targeturl, ShouldEqual, "https://example.com/mywebhook")
				So(webhook.Name, ShouldEqual, "My Awesome Webhook")
				So(webhook.Created, ShouldHappenOnOrBefore, stubNow())
			})
			Convey("Get webhook", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				InitClient("123")
				webhook := &Webhook{ID: "1"}
				err := webhook.Get()
				So(err, ShouldBeNil)
				So(webhook.ID, ShouldEqual, "123")
				So(webhook.Resource, ShouldEqual, "messages")
				So(webhook.Event, ShouldEqual, "created")
				So(webhook.Filter, ShouldEqual, "roomId=456")
				So(webhook.Targeturl, ShouldEqual, "https://example.com/mywebhook")
				So(webhook.Name, ShouldEqual, "My Awesome Webhook")
				So(webhook.Created, ShouldHappenOnOrBefore, stubNow())
				BaseURL = previousURL
			})
			Convey("Delete webhook", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				InitClient("123")
				webhook := &Webhook{ID: "1"}
				err := webhook.Delete()
				So(err, ShouldBeNil)
				BaseURL = previousURL
			})
		})
		Convey("For webhooks", func() {
			Convey("Get webhooks", func() {
				ts := serveHTTP(t)
				defer ts.Close()
				previousURL := BaseURL
				BaseURL = ts.URL
				InitClient("123")
				webhooks := &Webhooks{}
				err := webhooks.Get()
				So(err, ShouldBeNil)
				So(webhooks.Items[0].ID, ShouldEqual, "123")
				So(webhooks.Items[0].Resource, ShouldEqual, "messages")
				So(webhooks.Items[0].Event, ShouldEqual, "created")
				So(webhooks.Items[0].Filter, ShouldEqual, "roomId=456")
				So(webhooks.Items[0].Targeturl, ShouldEqual, "https://example.com/mywebhook")
				So(webhooks.Items[0].Name, ShouldEqual, "My Awesome Webhook")
				So(webhooks.Items[0].Created, ShouldHappenOnOrBefore, stubNow())
				BaseURL = previousURL
			})
		})
	})
}
