package mapper

import (
	"fmt"
	"os"

	"github.com/ONSdigital/dp-api-clients-go/v2/areas"
	coreModel "github.com/ONSdigital/dp-renderer/model"
)

const (
	pageType = "homepage"
)

type StartPageModel struct {
	coreModel.Page
	Greeting string `json:"greeting"`
	Version  string `json:"area_profiles_version"`
}

func CreateStartPage(basePage coreModel.Page) StartPageModel {
	model := StartPageModel{
		Page: basePage,
	}
	model.Metadata = coreModel.Metadata{
		Title: "Areas",
	}
	model.Page.Breadcrumb = append(model.Page.Breadcrumb, coreModel.TaxonomyNode{
		Title: "Home",
		URI:   "/",
	})
	model.Page.BetaBannerEnabled = true
	model.Version = getAreaProfilesVersion()
	return model
}

type ErrorDetails struct {
	Description string
	Title       string
}

type ErrorModel struct {
	coreModel.Page
	Error ErrorDetails
}

type AreaModel struct {
	coreModel.Page
	Name      string           `json:"name"`
	Code      string           `json:"code"`
	Level     string           `json:"level"`
	AreaType  string           `json:"area_type"`
	NomisLink string           `json:"nomis_link"`
	Ancestors []areas.Ancestor `json:"ancestors"`
	Siblings  []AreaModel      `json:"siblings"`
	Children  []AreaModel      `json:"children"`
	Relations []RelationLink   `json:"relations"`
	Version   string           `json:"area_profiles_version"`
}

// CreateAreaPage maps request area profile data to frontend view
func CreateAreaPage(basePage coreModel.Page, areaDetails areas.AreaDetails, relations []areas.Relation, lang string) AreaModel {
	// TODO - load the area data for the requested area once the API has been developed
	model := AreaModel{
		Page: basePage,
	}

	model.Page.Type = pageType
	model.Page.Language = lang
	// Area Details
	model.Name = areaDetails.Name
	model.AreaType = areaDetails.AreaType
	model.Code = areaDetails.Code
	// Relations
	model.Relations = createRelationLinks(relations)
	model.NomisLink = getNOMISLink(areaDetails.Code)
	model.Metadata = coreModel.Metadata{
		Title: fmt.Sprintf("%s Summary", model.Name),
	}

	model.Ancestors = areaDetails.Ancestors
	pageBreadcrumb := []coreModel.TaxonomyNode{{Title: "Home", URI: "/"}, {Title: "Areas", URI: "/areas"}}
	model.Page.Breadcrumb = createBreadcrumbs(areaDetails.Ancestors, pageBreadcrumb)
	model.Page.BetaBannerEnabled = true
	model.Version = getAreaProfilesVersion()

	return model
}

// Create404Page returns an ErrorModel struct for use with the error template
func Create404Page(basePage coreModel.Page, errorDetails ErrorDetails) ErrorModel {
	return ErrorModel{
		Page:  basePage,
		Error: errorDetails,
	}
}

func getNOMISLink(geoCode string) string {
	return fmt.Sprintf(`https://www.nomisweb.co.uk/reports/localarea?compare=%s`, geoCode)
}

func createBreadcrumbs(ancestors []areas.Ancestor, pageBreadcrumb []coreModel.TaxonomyNode) []coreModel.TaxonomyNode {
	// If we are on a country page ("England", "Wales") the []areas.Ancestor length will be 0
	if len(ancestors) > 0 {
		for i := len(ancestors) - 1; i >= 0; i-- {
			pageBreadcrumb = append(pageBreadcrumb, coreModel.TaxonomyNode{
				Title: ancestors[i].Name,
				URI:   "/areas/" + ancestors[i].Id,
			})
		}
	}
	return pageBreadcrumb
}

type RelationLink struct {
	Name string
	Href string
}

func createRelationLinks(relations []areas.Relation) []RelationLink {
	var relationLinks []RelationLink
	for _, relation := range relations {
		href := fmt.Sprintf("/areas/%s", relation.AreaCode)
		relationLinks = append(relationLinks, RelationLink{
			Name: relation.AreaName,
			Href: href,
		})
	}
	return relationLinks
}

func getAreaProfilesVersion() string {
	return os.Getenv("VERSION")
}
