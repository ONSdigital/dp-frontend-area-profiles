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
}
