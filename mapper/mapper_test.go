package mapper

import (
	"github.com/ONSdigital/dp-api-clients-go/v2/areas"
	"github.com/ONSdigital/dp-renderer/model"
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
		areaModel := CreateAreaPage(mdl, areaDetails)

		So(areaModel.BetaBannerEnabled, ShouldBeTrue)
		So(areaModel.Metadata.Title, ShouldEqual, areaModel.Name+" Summary")
		So(areaModel.Page.Breadcrumb, ShouldHaveLength, 2)
		So(areaModel.Page.Breadcrumb[0].Title, ShouldEqual, "Home")
		So(areaModel.Page.Breadcrumb[0].URI, ShouldEqual, "/")
		So(areaModel.Page.Breadcrumb[1].Title, ShouldEqual, "Areas")
		So(areaModel.Page.Breadcrumb[1].URI, ShouldEqual, "/areas")
	})
}
