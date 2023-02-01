package handlers

import (
	"context"
	"net/http"
	"sync"

	"github.com/ONSdigital/dp-areas-api/sdk/areas"
	"github.com/ONSdigital/dp-frontend-area-profiles/config"
	"github.com/ONSdigital/dp-frontend-area-profiles/mapper"
	"github.com/ONSdigital/dp-frontend-area-profiles/utils"
	dphandlers "github.com/ONSdigital/dp-net/v2/handlers"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
)

// GeographyStart Handler
func GeographyStart(ctx context.Context, cfg config.Config, c Clients) http.HandlerFunc {
	return dphandlers.ControllerHandler(func(w http.ResponseWriter, req *http.Request, lang, collectionID, accessToken string) {
		basePage := c.Render.NewBasePageModel()
		model := mapper.CreateStartPage(basePage)
		if cfg.EnableNewNavBar {
			mappedNavContent, err := c.CacheHelper.GetMappedNavigationContent(ctx, lang)
			if err == nil {
				model.NavigationContent = mappedNavContent
			}
		}
		c.Render.BuildPage(w, model, "geography-start")
	})
}

// GetArea Handler
func GetArea(ctx context.Context, cfg config.Config, c Clients) http.HandlerFunc {
	return dphandlers.ControllerHandler(func(w http.ResponseWriter, req *http.Request, lang, collectionID, accessToken string) {
		GetAreaViewHandler(w, req, ctx, c, cfg, lang, collectionID, accessToken)
	})
}

func GetAreaViewHandler(w http.ResponseWriter, req *http.Request, ctx context.Context, c Clients, cfg config.Config, lang, collectionID, accessToken string) {
	var err error
	var relationsErr error
	var areaData areas.AreaDetails
	var relationsData []areas.Relation
	vars := mux.Vars(req)
	areaID := vars["id"]
	acceptedLang := req.Header.Get("Accept-Language")
	var wg sync.WaitGroup
	wg.Add(2)
	// Remote requests
	go func() {
		defer wg.Done()
		areaData, err = c.AreasSDKClient.GetArea(ctx, accessToken, "", collectionID, areaID, acceptedLang)
		if err != nil {
			log.Error(ctx, "fetching Area Data", err)
			return
		}
	}()
	go func() {
		defer wg.Done()
		// Create a new local error variable otherwise we will incur a race condition when other goroutines access it
		relationsData, relationsErr = c.AreasSDKClient.GetRelations(ctx, accessToken, "", collectionID, areaID, acceptedLang)
		if relationsErr != nil {
			log.Error(ctx, "fetching area relations data", relationsErr)
			return
		}
	}()
	wg.Wait()
	basePage := c.Render.NewBasePageModel()
	if err != nil {
		// We only care about AreaDetails data errors for setting the status code & rendering the error template
		var errorDetails mapper.ErrorDetails
		resWriterStatusCode := utils.SetStatusCode(req, w, err)
		utils.SetErrorDetails(resWriterStatusCode, &errorDetails)
		c.Render.BuildPage(w, mapper.Create404Page(basePage, errorDetails), "error")
		return
	}
	//  View logic
	model := mapper.CreateAreaPage(basePage, areaData, relationsData, lang)
	if cfg.EnableNewNavBar {
		mappedNavContent, err := c.CacheHelper.GetMappedNavigationContent(ctx, lang)
		if err == nil {
			model.NavigationContent = mappedNavContent
		}
	}

	c.Render.BuildPage(w, model, "area-summary")
}
