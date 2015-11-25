package spark

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	WebhookJSON          = `{"id":"123","resource":"messages","event":"created","filter":"roomId=456","targetUrl":"https://example.com/mywebhook","name":"My Awesome Webhook","created":"2015-10-18T07:26:16-07:00"}`
	WebhooksJSON         = `{"items":[` + WebhookJSON + `]}`
	WebhooksResponseJSON = `{"id":"123","resource":"messages","event":"created","filter":"roomId=456","targetUrl":"https://example.com/mywebhook","name":"My Awesome Webhook","created":"2015-10-18T14:26:16+00:00"}`
)

func TestWebhooksSpec(t *testing.T) {
	ts := serveHTTP(t)
	defer ts.Close()
	previousURL := BaseURL
	BaseURL = ts.URL
	InitClient("123")
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
					Created:   &CreatedTime,
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
			})
			Convey("Get webhook", func() {
				webhook := &Webhook{ID: "1"}
				result, err := webhook.Get()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(webhook.ID, ShouldEqual, "123")
				So(webhook.Resource, ShouldEqual, "messages")
				So(webhook.Event, ShouldEqual, "created")
				So(webhook.Filter, ShouldEqual, "roomId=456")
				So(webhook.Targeturl, ShouldEqual, "https://example.com/mywebhook")
				So(webhook.Name, ShouldEqual, "My Awesome Webhook")
			})
			Convey("Delete webhook", func() {
				webhook := &Webhook{ID: "1"}
				result, err := webhook.Delete()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
			})
			Convey("Post webhook", func() {
				webhook := &Webhook{
					Resource:  "messages",
					Event:     "created",
					Filter:    "room=123",
					Targeturl: "http://foo.com/bar",
					Name:      "My Awesome webhook",
				}
				result, err := webhook.Post()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
			})
			Convey("Put webhook", func() {
				webhook := &Webhook{
					ID:        "1",
					Resource:  "messages",
					Event:     "created",
					Filter:    "room=123",
					Targeturl: "http://foo.com/bar",
					Name:      "My Awesome webhook",
				}
				result, err := webhook.Put()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
			})
		})
		Convey("For webhooks", func() {
			Convey("Get webhooks", func() {
				webhooks := &Webhooks{}
				result, err := webhooks.Get()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(webhooks.Items[0].ID, ShouldEqual, "123")
				So(webhooks.Items[0].Resource, ShouldEqual, "messages")
				So(webhooks.Items[0].Event, ShouldEqual, "created")
				So(webhooks.Items[0].Filter, ShouldEqual, "roomId=456")
				So(webhooks.Items[0].Targeturl, ShouldEqual, "https://example.com/mywebhook")
				So(webhooks.Items[0].Name, ShouldEqual, "My Awesome Webhook")
			})
			Convey("It should raise an error when no link cursor", func() {
				webhooks := &Webhooks{}
				result, err := webhooks.First()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "first cursor not available")
				result, err = webhooks.Next()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "next cursor not available")
				result, err = webhooks.Last()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "last cursor not available")
				result, err = webhooks.Previous()
				So(result, ShouldBeNil)
				So(err.Error(), ShouldEqual, "previous cursor not available")
			})
			Convey("Should get our link cursor", func() {
				webhooks := &Webhooks{}
				webhooks.FirstURL = "/webhooks/first"
				result, err := webhooks.First()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(webhooks.Items[0].ID, ShouldEqual, "123")
				webhooks.LastURL = "/webhooks/last"
				result, err = webhooks.Last()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(webhooks.Items[0].ID, ShouldEqual, "123")
				webhooks.NextURL = "/webhooks/next"
				result, err = webhooks.Next()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(webhooks.Items[0].ID, ShouldEqual, "123")
				webhooks.PreviousURL = "/webhooks/prev"
				result, err = webhooks.Previous()
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
				So(webhooks.Items[0].ID, ShouldEqual, "123")
			})
		})
	})
	BaseURL = previousURL
}
