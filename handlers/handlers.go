package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-frontend-area-profiles/config"
	"github.com/ONSdigital/dp-frontend-area-profiles/mapper"
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

// TODO: remove hello world example handler
// HelloWorld Handler
func HelloWorld(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		helloWorld(w, req, cfg)
	}
}

func helloWorld(w http.ResponseWriter, req *http.Request, cfg config.Config) {
	ctx := req.Context()
	greetingsModel := mapper.HelloModel{Greeting: "Hello", Who: "World"}
	m := mapper.HelloWorld(ctx, greetingsModel, cfg)

	b, err := json.Marshal(m)
	if err != nil {
		setStatusCode(req, w, err)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		log.Error(ctx, "failed to write bytes for http response", err)
		setStatusCode(req, w, err)
		return
	}
	return
}
