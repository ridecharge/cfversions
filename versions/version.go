package versions

import (
	"log"
)

func newVersion(props map[string]string, key_prefix string) *version {
	log.Print("Creating New Version Object")
	return &version{
		Version:         props[key_prefix+"version"],
		ApplicationName: props[key_prefix+"application_name"],
		PrivateEndPoint: props[key_prefix+"private_end_point"],
		PublicEndPoint:  props[key_prefix+"public_end_point"]}
}

type version struct {
	Version         string `json:"version"`
	ApplicationName string `json:"applicationName"`
	PrivateEndPoint string `json:"privateEndPoint"`
	PublicEndPoint  string `json:"publicEndPoint"`
}
