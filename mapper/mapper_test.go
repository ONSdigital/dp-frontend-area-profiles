package mapper

import (
	"testing"

	"github.com/ONSdigital/dp-areas-api/sdk/areas"
	"github.com/ONSdigital/dp-renderer/model"
	coreModel "github.com/ONSdigital/dp-renderer/model"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStartPageMapper(t *testing.T) {
	Convey("test CreateStartPage mapper correctly initialises a StartPageModel", t, func() {
		mdl := model.Page{}
		geoStartModel := CreateStartPage(mdl)

		So(geoStartModel.BetaBannerEnabled, ShouldBeTrue)
		So(geoStartModel.Metadata.Title, ShouldEqual, "Areas")
		So(geoStartModel.Page.Breadcrumb, ShouldHaveLength, 1)
		So(geoStartModel.Page.Breadcrumb[0].Title, ShouldEqual, "Home")
		So(geoStartModel.Page.Breadcrumb[0].URI, ShouldEqual, "/")
	})
}
func TestAreaProfilesMapper(t *testing.T) {
	Convey("test CreateAreaPage mapper correctly initialises an AreaModel", t, func() {
		areaDetails := areas.AreaDetails{
			Code:          "E92000001",
			Name:          "England",
			DateStarted:   "",
			DateEnd:       "",
			WelshName:     "",
			GeometricData: nil,
			AreaType:      "Country",
			Ancestors:     nil,
		}
		areaModel := CreateAreaPage(model.Page{}, areaDetails, []areas.Relation{}, "England")

		So(areaModel.BetaBannerEnabled, ShouldBeTrue)
		So(areaModel.Metadata.Title, ShouldEqual, areaModel.Name+" Summary")
		So(areaModel.Code, ShouldEqual, "E92000001")
		So(areaModel.Name, ShouldEqual, "England")
		So(areaModel.AreaType, ShouldEqual, "Country")
		So(areaModel.NomisLink, ShouldEqual, "https://www.nomisweb.co.uk/reports/localarea?compare=E92000001")
	})

	Convey("Test area profile page has breadcrumbs", t, func() {
		Convey("should contain area & home breadcrumbs", func() {
			expected := []coreModel.TaxonomyNode{
				{Title: "Home", URI: "/"},
				{Title: "Areas", URI: "/areas"},
			}
			ancestorData := []areas.Ancestor{}
			areaDetails := areas.AreaDetails{
				Ancestors: ancestorData, // For "England"
			}
			areaModel := CreateAreaPage(model.Page{}, areaDetails, []areas.Relation{}, "England")

			So(areaModel.Page.Breadcrumb, ShouldResemble, expected)
		})
		Convey("should display country if area is direct child", func() {
			// This test fixes a bug in `createBreadcrumbs()`
			expected := []coreModel.TaxonomyNode{
				{Title: "Home", URI: "/"},
				{Title: "Areas", URI: "/areas"},
				{Title: "Wales", URI: "/areas/W92000004"},
			}
			ancestorData := []areas.Ancestor{{Id: "W92000004", Name: "Wales"}}
			areaDetails := areas.AreaDetails{
				Ancestors: ancestorData, // For "Wales"
			}
			areaModel := CreateAreaPage(model.Page{}, areaDetails, []areas.Relation{}, "Wales")

			So(areaModel.Page.Breadcrumb, ShouldResemble, expected)
		})
		Convey("should contain breadcrumbs with county in the correct order", func() {
			expected := []coreModel.TaxonomyNode{
				{Title: "Home", URI: "/"},
				{Title: "Areas", URI: "/areas"},
				{Title: "England", URI: "/areas/E92000001"},
				{Title: "Yorkshire and the Humber", URI: "/areas/E12000003"},
			}
			ancestorData := []areas.Ancestor{{Id: "E12000003", Name: "Yorkshire and the Humber"}, {Id: "E92000001", Name: "England"}}
			areaDetails := areas.AreaDetails{
				Ancestors: ancestorData, // For "England"
			}
			areaModel := CreateAreaPage(model.Page{}, areaDetails, []areas.Relation{}, "England")

			So(areaModel.Page.Breadcrumb, ShouldResemble, expected)
		})

	})

	Convey("Test area profile page has child relations", t, func() {
		expected := []RelationLink{
			{Name: "Yorkshire and the Humber", Href: "/areas/E92000001"},
			{Name: "Whitby", Href: "/areas/E92000002"},
			{Name: "Hastings", Href: "/areas/E92000003"},
		}
		relations := []areas.Relation{
			{AreaName: "Yorkshire and the Humber", Href: "/v1/area/E92000001", AreaCode: "E92000001"},
			{AreaName: "Whitby", Href: "/v1/area/E92000002", AreaCode: "E92000002"},
			{AreaName: "Hastings", Href: "/v1/area/E92000003", AreaCode: "E92000003"},
		}
		areaModel := CreateAreaPage(model.Page{}, areas.AreaDetails{}, relations, "England")

		So(areaModel.Relations, ShouldResemble, expected)
	})

	Convey("Test GetRelationsHeading selects the correct relations heading", t, func() {
		ancestors := []areas.Ancestor{{}}
		result := GetRelationsHeading(ancestors, "heading1", "heading2", "England")
		exepected := "heading2 England"
		So(result, ShouldEqual, exepected)
		ancestors = []areas.Ancestor{{}, {}}
		result = GetRelationsHeading(ancestors, "heading1", "heading2", "England")
		exepected = "heading2 England"
		So(result, ShouldEqual, exepected)
		ancestors = []areas.Ancestor{}
		result = GetRelationsHeading(ancestors, "heading1", "heading2", "England")
		exepected = "heading1 England"
		So(result, ShouldEqual, exepected)
	})
}
