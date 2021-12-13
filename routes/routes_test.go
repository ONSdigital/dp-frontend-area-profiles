package routes

import (
	"context"
	"github.com/ONSdigital/dp-frontend-area-profiles/config"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetup(t *testing.T) {
	Convey("When Setup is called", t, func() {
		r := mux.NewRouter()
		ctx := context.Background()
		cfg, err := config.Get()
		clients := Clients{
			HealthCheckHandler: func(w http.ResponseWriter, req *http.Request){ return },
		}

		So(err, ShouldBeNil)

		Setup(ctx, r, cfg, clients)

		Convey("The following route(s) should have been added", func() {
			So(hasRoute(r, "/health", http.MethodGet), ShouldBeTrue)
			So(hasRoute(r, "/areas", http.MethodGet), ShouldBeTrue)
		})
	})
}

func hasRoute(r *mux.Router, path, method string) bool {
	req := httptest.NewRequest(method, path, nil)
	match := &mux.RouteMatch{}
	return r.Match(req, match)
}

