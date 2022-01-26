package handlers

import (
	"context"
	"errors"
	"github.com/ONSdigital/dp-frontend-area-profiles/config"
	mocks "github.com/ONSdigital/dp-frontend-area-profiles/handlers/mocks"
	coreModel "github.com/ONSdigital/dp-renderer/model"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
		mockRenderClient := mocks.NewMockRenderClient(mockCtrl)

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

	Convey("test GetArea", t, func() {

		mockConfig := config.Config{}
		mockRenderClient := mocks.NewMockRenderClient(mockCtrl)
		mockAreaApi := mocks.NewMockAreaApiClient(mockCtrl)
		mockRenderer := mocks.NewMockRendererClient(mockCtrl)
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
			mockRenderClient.EXPECT().NewBasePageModel().Return(coreModel.NewPage(cfg.PatternLibraryAssetsPath, cfg.SiteDomain))
			mockRenderClient.EXPECT().BuildPage(gomock.Any(), gomock.Any(), "area-summary")
			mockAreaApi.EXPECT().GetArea(ctx, "", "", "", "abc123")
			req := httptest.NewRequest("GET", "http://localhost:26600/areas/abc123", nil)

			router.ServeHTTP(w, req)

			So(w.Code, ShouldEqual, http.StatusOK)
		})
	})
}

func initialiseMockConfig() config.Config {
	return config.Config{
		PatternLibraryAssetsPath: "http://localhost:9000/dist",
		SiteDomain:               "ons",
	}
}
