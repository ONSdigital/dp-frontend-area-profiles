package handlers

import (
	"context"
	"errors"
	"github.com/ONSdigital/dp-frontend-area-profiles/config"
	coreModel "github.com/ONSdigital/dp-renderer/model"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ctx = context.Background()

type testCliError struct{}

func (e *testCliError) Error() string { return "client error" }
func (e *testCliError) Code() int     { return http.StatusNotFound }

func TestUnitHandlers(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	cfg := initialiseMockConfig()

	Convey("test setStatusCode", t, func() {

		Convey("test status code handles 404 response from client", func() {
			req := httptest.NewRequest("GET", "http://localhost:26600", nil)
			w := httptest.NewRecorder()
			err := &testCliError{}

			setStatusCode(req, w, err)

			So(w.Code, ShouldEqual, http.StatusNotFound)
		})

		Convey("test status code handles internal server error", func() {
			req := httptest.NewRequest("GET", "http://localhost:26600", nil)
			w := httptest.NewRecorder()
			err := errors.New("internal server error")

			setStatusCode(req, w, err)

			So(w.Code, ShouldEqual, http.StatusInternalServerError)
		})
	})

	Convey("test GetGeographyStart", t, func() {
		mockConfig := config.Config{}
		mockRenderClient := NewMockRenderClient(mockCtrl)

		router := mux.NewRouter()
		router.HandleFunc("/areas", GeographyStart(mockConfig, mockRenderClient))

		w := httptest.NewRecorder()

		Convey("it returns 200 when rendered successfully", func() {
			mockRenderClient.EXPECT().NewBasePageModel().Return(coreModel.NewPage(cfg.PatternLibraryAssetsPath, cfg.SiteDomain))
			mockRenderClient.EXPECT().BuildPage(gomock.Any(), gomock.Any(), "geography-start")
			req := httptest.NewRequest("GET", "http://localhost:26600/areas", nil)

			router.ServeHTTP(w, req)

			So(w.Code, ShouldEqual, http.StatusOK)
		})
	})
}

func TestGetAreaWithSpies(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	cfg := initialiseMockConfig()
	Convey("test GetArea", t, func() {
		//var wg sync.WaitGroup
		//wg.Add(1)
		//defer wg.Wait()
		mockConfig := config.Config{}
		mockRenderClient := NewMockRenderClient(mockCtrl)
		mockAreaApi := NewMockAreaApiClient(mockCtrl)
		mockRenderer := NewMockRendererClient(mockCtrl)
		router := mux.NewRouter()

		c := Clients{
			HealthCheckHandler: func(w http.ResponseWriter, req *http.Request) {},
			Render:             mockRenderClient,
			AreaApi:            mockAreaApi,
			Renderer:           mockRenderer,
		}

		ctx := context.Background()
		router.HandleFunc("/areas/{id}", GetArea(ctx, mockConfig, c))

		w := httptest.NewRecorder()
		Convey("it returns 200 when rendered successfully", func() {
			// Add moq tests

			mockRenderClient.EXPECT().NewBasePageModel().Return(coreModel.NewPage(cfg.PatternLibraryAssetsPath, cfg.SiteDomain))
			mockRenderClient.EXPECT().BuildPage(gomock.Any(), gomock.Any(), "area-summary")
			mockAreaApi.EXPECT().GetArea(ctx, "", "", "", "E92000001", "").AnyTimes()
			mockAreaApi.EXPECT().GetRelations(ctx, "", "", "", "E92000001", "").AnyTimes()

			req := httptest.NewRequest("GET", "http://localhost:26600/areas/E92000001", nil)

			router.ServeHTTP(w, req)
			So(w.Code, ShouldEqual, http.StatusOK)
		})
	})

}

func TestUnitReadHandlerSuccess(t *testing.T) {
	t.Parallel()

}

func initialiseMockConfig() config.Config {
	return config.Config{
		PatternLibraryAssetsPath: "http://localhost:9000/dist",
		SiteDomain:               "ons",
	}
}
