package handlers

import (
	"fmt"
	"github.com/ONSdigital/dp-frontend-area-profiles/mapper"
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
		//ctx := req.Context()
		basePage := rc.NewBasePageModel()
		fmt.Println()
		model := mapper.CreateStartPage(basePage)
		rc.BuildPage(w, model, "geography-start")
	}
}
