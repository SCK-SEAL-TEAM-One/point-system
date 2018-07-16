package api

import (
	"fmt"
	"net/http"
)

type PointHandler struct {
	Point string
}

func (ph PointHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"point":"%s"}`, ph.Point)
}
