package mapper

import (
	"fmt"

	"github.com/ONSdigital/dp-api-clients-go/v2/areas"
	coreModel "github.com/ONSdigital/dp-renderer/model"
)

const (
	pageType = "homepage"
)

type StartPageModel struct {
	coreModel.Page
	Greeting string `json:"greeting"`
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
	return model
}

type AreaModel struct {
	coreModel.Page
	Name      string           `json:"name"`
	Level     string           `json:"level"`
	Code      string           `json:"code"`
	Ancestors []areas.Ancestor `json:"ancestors"`
	Siblings  []AreaModel      `json:"siblings"`
	Children  []AreaModel      `json:"children"`
	Relations []areas.Relation `json:"relations"`
}

// CreateAreaPage maps request area profile data to frontend view
func CreateAreaPage(basePage coreModel.Page, areaDetails areas.AreaDetails, relations []areas.Relation, ancestors []areas.Ancestor, lang string) AreaModel {
	// TODO - load the area data for the requested area once the API has been developed
	model := AreaModel{
		Page: basePage,
	}

	model.Page.Type = pageType
	model.Page.Language = lang
	// Area Details
	model.Name = areaDetails.Name
	model.Code = areaDetails.Code
	// Relations
	model.Relations = relations
	model.Metadata = coreModel.Metadata{
		Title: fmt.Sprintf("%s Summary", model.Name),
	}
	model.Page.Breadcrumb = append(model.Page.Breadcrumb, coreModel.TaxonomyNode{
		Title: "Home",
		URI:   "/",
	})
	model.Page.Breadcrumb = append(model.Page.Breadcrumb, coreModel.TaxonomyNode{
		Title: "Areas",
		URI:   "/areas",
	})
	model.Ancestors = ancestors
	// Only add children to breadcrumbs if we are not on a country / root area page
	if len(ancestors) > 1 {
		for _, ancestor := range ancestors {
			model.Page.Breadcrumb = append(model.Page.Breadcrumb, coreModel.TaxonomyNode{
				Title: ancestor.Name,
				URI:   "/areas/" + ancestor.Name,
			})
		}
	}
	model.Page.BetaBannerEnabled = true
	return model
}
