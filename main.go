package main

import (
	"net/http"

	"github.com/techlogger/iam/controllers"
)

func main() {
	sc := controllers.NewSessionController()
	http.Handle("/signin", *sc)
	http.Handle("/signup", *sc)
	http.ListenAndServe(":4000", nil)
}
