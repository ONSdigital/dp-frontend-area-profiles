package mapper

import coreModel "github.com/ONSdigital/dp-renderer/model"


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
