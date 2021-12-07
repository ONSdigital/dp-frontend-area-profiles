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
	model.Greeting = "Welcome to area profiles"
	return model
}
