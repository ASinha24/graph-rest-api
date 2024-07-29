package utils

import (
	"encoding/json"
	"net/http"

	"github.com/asinha24/graph-rest-api/api"
)

func WriteResponse(status int, response interface{}, rw http.ResponseWriter) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	if response != nil {
		json.NewEncoder(rw).Encode(response)
	}
}

func WriteErrorResponse(status int, err error, rw http.ResponseWriter) {
	graphErr, ok := err.(*api.GraphError)
	if !ok {
		graphErr = &api.GraphError{
			Code:        0,
			Message:     "failed in serving request",
			Description: graphErr.Error(),
		}
	} else {
		status = graphErr.Code.HTTPStatus()
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(graphErr)
}
