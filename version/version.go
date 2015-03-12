package version

import (
	"log"
)

func NewVersion(props map[string]string) *version {
	log.Print("Creating New Version Object")
	return &version{
		Version:         props["Version"],
		ApplicationName: props["ApplicationName"],
		PrivateEndPoint: props["PrivateEndPoint"],
		PublicEndPoint:  props["PublicEndPoint"]}
}

type version struct {
	Version         string `json:"version"`
	ApplicationName string `json:"name"`
	PrivateEndPoint string `json:"privateEndPoint"`
	PublicEndPoint  string `json:"publicEndPoint"`
}
