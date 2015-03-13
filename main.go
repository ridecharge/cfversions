package main

import (
	"github.com/ridecharge/cf-versions/version"
	"net/http"
)

func main() {
	version.RegisterHandlers()
	http.ListenAndServe(":8080", nil)
}
