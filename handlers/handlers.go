package handlers

import (
	"context"
	"github.com/ONSdigital/dp-frontend-area-profiles/mapper"
	dphandlers "github.com/ONSdigital/dp-net/handlers"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/ONSdigital/dp-frontend-area-profiles/config"
	"github.com/ONSdigital/log.go/v2/log"
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
		func(w http.ResponseWriter, req *http.Request, lang, collectionID, accessToken string) {
			vars := mux.Vars(req)
			areaID := vars["id"]
			// TODO add when available: AcceptedLang := req.Header.Get("Accepted-Language")
			areaData, err := c.AreaApi.GetArea(ctx, accessToken, "", collectionID, areaID)
			if err != nil {
				log.Error(ctx, "Fetching Area Data", err)
			}

			basePage := c.Render.NewBasePageModel()
			model := mapper.CreateAreaPage(basePage, areaData)
			c.Render.BuildPage(w, model, "area-summary")
		}(w, req, lang, collectionID, accessToken)
	})
}
