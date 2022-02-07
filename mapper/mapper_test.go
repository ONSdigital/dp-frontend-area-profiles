package mapper

import (
	"testing"

	"github.com/ONSdigital/dp-api-clients-go/v2/areas"
	"github.com/ONSdigital/dp-renderer/model"
	. "github.com/smartystreets/goconvey/convey"
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
		ancestors := []areas.Ancestor{areas.Ancestor{
			Name:      "England",
			Level:     "",
			Code:      "E92000001",
			Ancestors: []areas.Ancestor{},
			Siblings:  nil,
			Children:  nil,
		}}
		areaModel := CreateAreaPage(mdl, areaDetails, relations, ancestors)

		So(areaModel.BetaBannerEnabled, ShouldBeTrue)
		So(areaModel.Metadata.Title, ShouldEqual, areaModel.Name+" Summary")
		So(areaModel.Page.Breadcrumb, ShouldHaveLength, 3)
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
		So(areaModel.Ancestors[0].Name, ShouldEqual, "England")
		So(areaModel.Ancestors[0].Level, ShouldEqual, "")
		So(areaModel.Ancestors[0].Code, ShouldEqual, "E92000001")
		So(areaModel.Ancestors[0].Ancestors, ShouldHaveLength, 0)
		So(areaModel.Ancestors[0].Siblings, ShouldEqual, nil)
		So(areaModel.Ancestors[0].Children, ShouldEqual, nil)
	})
}
