package main

import (
	"github.com/ridecharge/cfversions/versions"
	"net/http"
)

func main() {
	versions.RegisterHandlers()
	http.ListenAndServe(":8080", nil)
}
