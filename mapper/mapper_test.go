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
		So(len(geoStartModel.Page.Breadcrumb), ShouldEqual, 1)
		So(geoStartModel.Page.Breadcrumb[0].Title, ShouldEqual, "Home")
		So(geoStartModel.Page.Breadcrumb[0].URI, ShouldEqual, "/")
	})

	Convey("test CreateAreaPage mapper correctly initialises an AreaModel", t, func() {
		areaModel := CreateAreaPage(mdl)

		So(areaModel.BetaBannerEnabled, ShouldBeTrue)
		So(areaModel.Metadata.Title, ShouldEqual, areaModel.Name + " Summary")
		So(len(areaModel.Page.Breadcrumb), ShouldEqual, len(areaModel.Ancestors) + 2)
		So(areaModel.Page.Breadcrumb[0].Title, ShouldEqual, "Home")
		So(areaModel.Page.Breadcrumb[0].URI, ShouldEqual, "/")
		So(areaModel.Page.Breadcrumb[1].Title, ShouldEqual, "Areas")
		So(areaModel.Page.Breadcrumb[1].URI, ShouldEqual, "/areas")
		if len(areaModel.Ancestors) > 0 {
			So(areaModel.Page.Breadcrumb[2].Title, ShouldEqual, areaModel.Ancestors[0].Name)
			So(areaModel.Page.Breadcrumb[2].URI, ShouldEqual, "/areas/" + areaModel.Ancestors[0].Code)
		}
	})
}
