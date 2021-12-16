package mapper

import (
	"fmt"
	coreModel "github.com/ONSdigital/dp-renderer/model"
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
		Title:       "Areas",
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
	Name string `json:"name"`
	Level string `json:"level"`
	Code string `json:"code"`
	Ancestors []AreaModel `json:"ancestors"`
	Siblings []AreaModel `json:"siblings"`
	Children []AreaModel `json:"children"`
}

func CreateAreaPage(basePage coreModel.Page) AreaModel {
	// TODO - load the area data for the requested area once the API has been developed
	model := AreaModel{
		Page: basePage,
	}
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
	for _, ancestor := range model.Ancestors {
		model.Page.Breadcrumb = append(model.Page.Breadcrumb, coreModel.TaxonomyNode{
			Title: ancestor.Name,
			URI:   "/areas/" + ancestor.Code,
		})
	}
	model.Page.BetaBannerEnabled = true
	return model
}
