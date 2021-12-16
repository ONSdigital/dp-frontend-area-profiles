package mapper

import (
	"github.com/ONSdigital/dp-renderer/model"
	"testing"

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
		areaModel := CreateAreaPage(mdl)

		So(areaModel.BetaBannerEnabled, ShouldBeTrue)
		So(areaModel.Metadata.Title, ShouldEqual, areaModel.Name + " Summary")
		So(areaModel.Page.Breadcrumb, ShouldHaveLength, 2)
		So(areaModel.Page.Breadcrumb[0].Title, ShouldEqual, "Home")
		So(areaModel.Page.Breadcrumb[0].URI, ShouldEqual, "/")
		So(areaModel.Page.Breadcrumb[1].Title, ShouldEqual, "Areas")
		So(areaModel.Page.Breadcrumb[1].URI, ShouldEqual, "/areas")
	})
}
