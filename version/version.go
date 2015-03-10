package version

import (
	"curbformation-version-service/version_service"
	"encoding/json"
)

func init() {
}

func New() {
	return &version{}
}

type version struct {
	Version         string `json:"version"`
	Name            string `json:"name"`
	PrivateEndPoint string `json:"privateUrl"`
	PublicEndPoint  string `json:"publicUrl"`
}

type Version interface {
}
