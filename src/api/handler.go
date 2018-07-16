package api

import (
	"net/http"
)

func ShowPointHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte(`{"point":"100"}`))
}
