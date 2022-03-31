package utils

import (
	"net/http"

	"github.com/ONSdigital/log.go/v2/log"
)

type ClientError interface {
	Error() string
	Code() int
}

// SetStatusCode returns the response status code
func SetStatusCode(req *http.Request, w http.ResponseWriter, err error) int {
	status := http.StatusInternalServerError
	if err, ok := err.(ClientError); ok {
		if err.Code() == http.StatusNotFound {
			status = err.Code()
		}
	}
	log.Error(req.Context(), "setting-response-status", err)
	w.WriteHeader(status)
	return status
}
