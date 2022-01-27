package mapper

import (
	"github.com/ONSdigital/dp-api-clients-go/v2/areas"
	"github.com/ONSdigital/dp-renderer/model"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUnitMapper(t *testing.T) {
	mdl := model.Page{}

	Convey("test CreateStartPage mapper correctly initialises a StartPageModel", t, func() {
		geoStartModel := CreateStartPage(mdl)

		So(geoStartModel.BetaBannerEnabled, ShouldBeTrue)
		So(geoStartModel.Metadata.Title, ShouldEqual, "Areas")
		So(geoStartModel.Page.Breadcrumb, ShouldHaveLength, 1)
		So(geoStartModel.Page.Breadcrumb[0].Title, ShouldEqual, "Home")
		So(geoStartModel.Page.Breadcrumb[0].URI, ShouldEqual, "/")
	})

	Convey("test CreateAreaPage mapper correctly initialises an AreaModel", t, func() {
		areaDetails := areas.AreaDetails{
			Code:          "",
			Name:          "",
			DateStarted:   "",
			DateEnd:       "",
			WelshName:     "",
			GeometricData: nil,
			Visible:       false,
			AreaType:      "",
		}
		relations := []areas.Relation{
			{
				AreaCode: "E12000001",
				AreaName: "North East",
				Href:     "/v1/area/E12000001",
			},
			{
				AreaCode: "E12000002",
				AreaName: "North West",
				Href:     "/v1/area/E12000002",
			},
		}
		areaModel := CreateAreaPage(mdl, areaDetails, relations)

		So(areaModel.BetaBannerEnabled, ShouldBeTrue)
		So(areaModel.Metadata.Title, ShouldEqual, areaModel.Name+" Summary")
		So(areaModel.Page.Breadcrumb, ShouldHaveLength, 2)
		So(areaModel.Page.Breadcrumb[0].Title, ShouldEqual, "Home")
		So(areaModel.Page.Breadcrumb[0].URI, ShouldEqual, "/")
		So(areaModel.Page.Breadcrumb[1].Title, ShouldEqual, "Areas")
		So(areaModel.Page.Breadcrumb[1].URI, ShouldEqual, "/areas")
		So(areaModel.Relations[0].AreaCode, ShouldEqual, "E12000001")
		So(areaModel.Relations[0].AreaName, ShouldEqual, "North East")
		So(areaModel.Relations[0].Href, ShouldEqual, "/v1/area/E12000001")
		So(areaModel.Relations[1].AreaCode, ShouldEqual, "E12000002")
		So(areaModel.Relations[1].AreaName, ShouldEqual, "North West")
		So(areaModel.Relations[1].Href, ShouldEqual, "/v1/area/E12000002")
	})
}
