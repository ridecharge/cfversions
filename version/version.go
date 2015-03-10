package version

import (
	"encoding/json"
)

type version struct {
	Version         String `json:"version"`
	Name            String `json:"appName"`
	PrivateEndPoint String `json:"privateUrl"`
	PublicEndPoint  String `json:"publicUrl"`
}

type Version interface {
}
