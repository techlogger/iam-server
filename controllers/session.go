package controllers

import "net/http"

type sessionController struct{}

func (sc sessionController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}

func NewSessionController() *sessionController {
	return &sessionController{}
}
