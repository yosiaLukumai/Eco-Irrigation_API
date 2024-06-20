package utils

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RouteParam(r *http.Request, paramName string) string {
	vars := mux.Vars(r)
	return vars[paramName]
}
