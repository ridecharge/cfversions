package main

import (
	"github.com/ridecharge/curbformation-version-service/version"
	"net/http"
)

func main() {
	version.RegisterHandlers()
	http.ListenAndServe(":8080", nil)
}
