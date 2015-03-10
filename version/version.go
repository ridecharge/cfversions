package version

import ()

func init() {
}

func NewVersion() Version {
	return &version{verServ: NewVersionServ()}
}

type version struct {
	Version         string `json:"version"`
	Name            string `json:"name"`
	PrivateEndPoint string `json:"privateUrl"`
	PublicEndPoint  string `json:"publicUrl"`
	verServ         VersionServ
}

type Version interface {
}
