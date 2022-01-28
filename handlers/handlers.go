package handlers

import (
	"context"
	"github.com/ONSdigital/dp-api-clients-go/v2/areas"
	"github.com/ONSdigital/dp-frontend-area-profiles/config"
	"github.com/ONSdigital/dp-frontend-area-profiles/mapper"
	dphandlers "github.com/ONSdigital/dp-net/handlers"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
)

func setStatusCode(req *http.Request, w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	if err, ok := err.(ClientError); ok {
		if err.Code() == http.StatusNotFound {
			status = err.Code()
		}
	}
	log.Error(req.Context(), "setting-response-status", err)
	w.WriteHeader(status)
}

// GeographyStart Handler
func GeographyStart(cfg config.Config, rc RenderClient) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		basePage := rc.NewBasePageModel()
		model := mapper.CreateStartPage(basePage)
		rc.BuildPage(w, model, "geography-start")
	}
}

// GetArea Handler
func GetArea(ctx context.Context, cfg config.Config, c Clients) http.HandlerFunc {
	return dphandlers.ControllerHandler(func(w http.ResponseWriter, req *http.Request, lang, collectionID, accessToken string) {
		GetAreaViewHandler(w, req, ctx, c, lang, collectionID, accessToken)
	})
}

func GetAreaViewHandler(w http.ResponseWriter, req *http.Request, ctx context.Context, c Clients, lang, collectionID, accessToken string) {
	var err error
	ctx, cancel := context.WithCancel(req.Context())
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
		areaData, err = c.AreaApi.GetArea(ctx, accessToken, "", collectionID, areaID, acceptedLang)
		if err != nil {
			log.Error(ctx, "Fetching Area Data", err)
			cancel()
			return
		}
	}()
	go func() {
		defer wg.Done()
		// Create a new local error variable otherwise we will incur a race condition when other goroutines access it
		var relationsErr error
		relationsData, relationsErr = c.AreaApi.GetRelations(ctx, accessToken, "", collectionID, areaID, acceptedLang)
		if relationsErr != nil {
			log.Error(ctx, "Fetching area relations data", err)
			cancel()
			return
		}
	}()

	wg.Wait()
	if err != nil {
		log.Error(ctx, "Error fetching from remote", err)
		//setStatusCode(w, req, err)
		return
	}
	relationsData = []areas.Relation{}
	//  View logic
	basePage := c.Render.NewBasePageModel()
	model := mapper.CreateAreaPage(basePage, areaData, relationsData)
	c.Render.BuildPage(w, model, "area-summary")
}
