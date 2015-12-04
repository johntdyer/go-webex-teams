package spark

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	stubNow     = func() time.Time { return time.Unix(1445178376, 0) }
	CreatedTime = time.Unix(1445178376, 0)
	APIError    = `{"message":"Failed to get conversation.","errors":[{"description":"Failed to get conversation."}],"trackingId":"go-spark_6D15B3DA-BF4B-0601-C7DB-F9315AE0783E_76"}`
)

func TestClientSpec(t *testing.T) {
	Convey("Sould parse a link header", t, func() {
		Convey("Parse url", func() {
			url := `<https://api.github.com/search/code?q=addClass+user_mozilla&page=34>`
			So(parseURL(url), ShouldEqual, "https://api.github.com/search/code?q=addClass+user_mozilla&page=34")
		})
		Convey("Parse header itself", func() {
			link := `<https://api.ciscospark.com/v1/applications/first>; rel="first", <https://api.ciscospark.com/v1/applications/prev>; rel="prev", <https://api.ciscospark.com/v1/applications/next>; rel="next", <https://api.ciscospark.com/v1/applications/last>; rel="last"`
			links := parseLink(link)
			So(links.NextURL, ShouldEqual, "https://api.ciscospark.com/v1/applications/next")
			So(links.LastURL, ShouldEqual, "https://api.ciscospark.com/v1/applications/last")
			So(links.FirstURL, ShouldEqual, "https://api.ciscospark.com/v1/applications/first")
			So(links.PreviousURL, ShouldEqual, "https://api.ciscospark.com/v1/applications/prev")
		})
	})
	Convey("Constants should be set", t, func() {
		So(BaseURL, ShouldStartWith, "https://")
		So(BaseURL, ShouldNotEndWith, "/")
		So(ApplicationsResource, ShouldStartWith, "/")
		So(MembershipsResource, ShouldStartWith, "/")
		So(MessagesResource, ShouldStartWith, "/")
		So(RoomsResource, ShouldStartWith, "/")
		So(SubscriptionsResource, ShouldStartWith, "/")
	})

	Convey("Should create an *http.Client", t, func() {
		InitClient(&Authorization{AccessToken: "123"})
		So(ActiveClient.Authorization.AccessToken, ShouldEqual, "123")
		So(reflect.TypeOf(ActiveClient.HTTP).String(), ShouldEqual, "*http.Client")
	})
	Convey("Should set the HTTP headers properly", t, func() {
		InitClient(&Authorization{AccessToken: "123"})
		req, _ := http.NewRequest("GET", "http://foo.com", nil)
		setHeaders(req)
		So(req.Header.Get("Authorization"), ShouldEqual, "Bearer 123")
		So(req.Header.Get("Content-Type"), ShouldEqual, "application/json")
		So(req.Header.Get("Accept"), ShouldEqual, "application/json")
		So(req.Header.Get("TrackingID"), ShouldEqual, TrackingID())
	})
	Convey("Should DELETE, GET, POST and PUT request", t, func() {
		ts := serveHTTP(t)
		defer ts.Close()
		previousURL := BaseURL
		BaseURL = ts.URL
		InitClient(&Authorization{AccessToken: "123"})
		Convey("DELETE", func() {
			Convey("Happy case", func() {
				result, err := delete("/foo")
				So(result, ShouldBeNil)
				So(err, ShouldBeNil)
			})
			Convey("Negative case", func() {
				result, err := delete("/negative")
				So(err.Error(), ShouldEqual, "400 Bad Request")
				So(result.Errors[0].Description, ShouldEqual, "Failed to get conversation.")
				So(result.Message, ShouldEqual, "Failed to get conversation.")
				So(result.Trackingid, ShouldEqual, "go-spark_6D15B3DA-BF4B-0601-C7DB-F9315AE0783E_76")
			})
		})
		Convey("GET", func() {
			Convey("Happy case", func() {
				body, _, err := get("/foo")
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, "you GOT")
			})
			Convey("Negative case", func() {
				body, _, err := get("/negative")
				result := &Result{}
				json.Unmarshal(body, result)
				So(err.Error(), ShouldEqual, "400 Bad Request")
				So(result.Errors[0].Description, ShouldEqual, "Failed to get conversation.")
				So(result.Message, ShouldEqual, "Failed to get conversation.")
				So(result.Trackingid, ShouldEqual, "go-spark_6D15B3DA-BF4B-0601-C7DB-F9315AE0783E_76")
			})
		})
		message := "foo-bar"
		Convey("POST", func() {
			Convey("Happy case", func() {
				body, _, err := post("/foo", []byte(message))
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, "you POSTED")
			})
			Convey("Negative case", func() {
				body, _, err := post("/negative", []byte(message))
				result := &Result{}
				json.Unmarshal(body, result)
				So(err.Error(), ShouldEqual, "400 Bad Request")
				So(result.Errors[0].Description, ShouldEqual, "Failed to get conversation.")
				So(result.Message, ShouldEqual, "Failed to get conversation.")
				So(result.Trackingid, ShouldEqual, "go-spark_6D15B3DA-BF4B-0601-C7DB-F9315AE0783E_76")
			})
		})
		Convey("PUT", func() {
			Convey("Happy case", func() {
				body, _, err := put("/foo", []byte(message))
				So(err, ShouldBeNil)
				So(string(body), ShouldEqual, "you PUT")
			})
			Convey("Negative case", func() {
				body, _, err := put("/negative", []byte(message))
				result := &Result{}
				json.Unmarshal(body, result)
				So(err.Error(), ShouldEqual, "400 Bad Request")
				So(result.Errors[0].Description, ShouldEqual, "Failed to get conversation.")
				So(result.Message, ShouldEqual, "Failed to get conversation.")
				So(result.Trackingid, ShouldEqual, "go-spark_6D15B3DA-BF4B-0601-C7DB-F9315AE0783E_76")
			})
		})
		BaseURL = previousURL
	})
}

// serveHTTP serves up a test server emulating the Tropo Gateway
func serveHTTP(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		switch req.URL.String() {
		case "/foo":
			Convey("Should receive the correct body from a POST/PUT request", t, func() {
				switch req.Method {
				case "DELETE":
					w.WriteHeader(200)
				case "GET":
					w.WriteHeader(200)
					w.Write([]byte("you GOT"))
				case "POST":
					So(string(body), ShouldEqual, "foo-bar")
					w.WriteHeader(200)
					w.Write([]byte("you POSTED"))
				case "PUT":
					So(string(body), ShouldEqual, "foo-bar")
					w.WriteHeader(200)
					w.Write([]byte("you PUT"))
				}
			})
		case "/negative":
			w.WriteHeader(400)
			w.Write([]byte(APIError))
		case ApplicationsResource:
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(ApplicationsJSON))
			case "POST":
				w.WriteHeader(200)
				w.Write([]byte(ApplicationResponseJSON))
			}
		case "/applications/next", "/applications/last", "/applications/prev", "/applications/first":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(ApplicationsJSON))
			}
		case ApplicationsResource + "/1":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(ApplicationJSON))
			case "DELETE":
				w.WriteHeader(200)
			case "PUT":
				w.WriteHeader(200)
				w.Write([]byte(ApplicationResponseJSON))
			default:
				w.WriteHeader(404)
				t.Error("Unknown HTTP method for Applications")
			}
		case MembershipsResource:
			switch req.Method {
			case "POST":
				w.WriteHeader(200)
				w.Write([]byte(MembershipResponseJSON))
			}
		case MembershipsResource + "?roomId=123":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(MembershipsJSON))
			}
		case "/memberships/next", "/memberships/last", "/memberships/prev", "/memberships/first":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(MessagesJSON))
			}
		case MembershipsResource + "/1":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(MembershipJSON))
			case "DELETE":
				w.WriteHeader(200)
			case "PUT":
				w.WriteHeader(200)
				w.Write([]byte(MembershipResponseJSON))
			default:
				w.WriteHeader(404)
				t.Error("Unknown HTTP method for Memberships")
			}
		case MessagesResource:
			switch req.Method {
			case "POST":
				w.WriteHeader(200)
				w.Write([]byte(MessageResponseJSON))
			}
		case MessagesResource + "?roomId=1":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(MessagesJSON))
			}
		case MessagesResource + "/1":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(MessageJSON))
			case "DELETE":
				w.WriteHeader(200)
			default:
				w.WriteHeader(404)
				t.Error("Unknown HTTP method for Messages")
			}
		case "/messages/next", "/messages/last", "/messages/prev", "/messages/first":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(MessagesJSON))
			}
		case PeopleResource + "?email=john@doe.com":
			if req.Method == "GET" {
				w.WriteHeader(200)
				w.Write([]byte(PeopleJSON))
			}
		case "/people/next", "/people/last", "/people/prev", "/people/first":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(PeopleJSON))
			}
		case PeopleResource + "/1", PeopleResource + "/me":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(PersonJSON))
			case "DELETE":
				w.WriteHeader(200)
			default:
				w.WriteHeader(404)
				t.Error("Unknown HTTP method for People")
			}
		case RoomsResource:
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(RoomsJSON))
			case "POST":
				w.WriteHeader(200)
				w.Write([]byte(RoomResponseJSON))
			}
		case "/rooms/next", "/rooms/last", "/rooms/prev", "/rooms/first":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(RoomsJSON))
			}
		case RoomsResource + "/1":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(RoomJSON))
			case "PUT":
				w.WriteHeader(200)
				w.Write([]byte(RoomResponseJSON))
			case "DELETE":
				w.WriteHeader(200)
			default:
				w.WriteHeader(404)
				t.Error("Unknown HTTP method for Room")
			}
		case SubscriptionsResource:
			if req.Method == "GET" {
				w.WriteHeader(200)
				w.Write([]byte(SubscriptionsJSON))
			}
		case "/subscriptions/next", "/subscriptions/last", "/subscriptions/prev", "/subscriptions/first":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(SubscriptionsJSON))
			}
		case SubscriptionsResource + "?personId=123":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(SubscriptionsJSON))
			}
		case SubscriptionsResource + "/1":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(SubscriptionJSON))
			case "DELETE":
				w.WriteHeader(200)
			default:
				w.WriteHeader(404)
				t.Error("Unknown HTTP method for Subscriptions")
			}
		case WebhooksResource:
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(WebhooksJSON))
			case "POST":
				w.WriteHeader(200)
				w.Write([]byte(WebhooksResponseJSON))
			}
		case "/webhooks/next", "/webhooks/last", "/webhooks/prev", "/webhooks/first":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(WebhooksJSON))
			}
		case WebhooksResource + "/1":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(WebhookJSON))
			case "DELETE":
				w.WriteHeader(200)
			case "PUT":
				w.WriteHeader(200)
				w.Write([]byte(WebhooksResponseJSON))
			default:
				w.WriteHeader(404)
				t.Error("Unknown HTTP method for Webhooks")
			}
		default:
			w.WriteHeader(404)
			t.Error("Unknown HTTP request")
		}
	}))
}
