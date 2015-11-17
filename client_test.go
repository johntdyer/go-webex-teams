package spark

import (
	// "fmt"
	. "github.com/smartystreets/goconvey/convey"
	// "io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

var (
	stubNow    = func() time.Time { return time.Unix(1445178376, 0) }
	PersonJSON = `{"id":"OTZhYmMyYWEtM2RjYy0xMWU1LWExNTItZmUzNDgxOWNkYzlh","emails":"johnny.chang@foomail.com","displayName":"John Andersen","avatar":"TODO","created":"2015-10-18T14:26:16+00:00"}`
	PeopleJSON = `{"items":[` + PersonJSON + `]}`
)

func TestClientSpec(t *testing.T) {
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
		client := NewClient("1234")
		So(reflect.TypeOf(client).String(), ShouldEqual, "spark.Client")
		So(client.Token, ShouldEqual, "1234")
		So(reflect.TypeOf(client.HTTP).String(), ShouldEqual, "*http.Client")
	})
}

// serveHTTP serves up a test server emulating the Tropo Gateway
func serveHTTP(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// body, _ := ioutil.ReadAll(req.Body)
		// req.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		switch req.URL.String() {
		case ApplicationsResource:
			if req.Method == "GET" {
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
			default:
				w.WriteHeader(404)
				t.Error("Unknown HTTP method for Applications")
			}
		case MembershipsResource:
			if req.Method == "GET" {
				w.WriteHeader(200)
				w.Write([]byte(MembershipsJSON))
			}
		case MembershipsResource + "/1":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(MembershipJSON))
			case "DELETE":
				w.WriteHeader(200)
			default:
				w.WriteHeader(404)
				t.Error("Unknown HTTP method for Memberships")
			}
		case MessagesResource:
			if req.Method == "GET" {
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
		case RoomsResource:
			if req.Method == "GET" {
				w.WriteHeader(200)
				w.Write([]byte(RoomsJSON))
			}
		case RoomsResource + "/1":
			switch req.Method {
			case "GET":
				w.WriteHeader(200)
				w.Write([]byte(RoomJSON))
			case "DELETE":
				w.WriteHeader(200)
			default:
				w.WriteHeader(404)
				t.Error("Unknown HTTP method for Rooms")
			}
		case SubscriptionsResource:
			if req.Method == "GET" {
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
		default:
			w.WriteHeader(404)
			t.Error("Unknown HTTP request")
		}
	}))
}
