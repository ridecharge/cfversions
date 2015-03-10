package version

import "github.com/goamz/goamz/cloudformation"
import "github.com/goamz/goamz/aws"

var verServ

func init() {
	verServ = &versionServ{}
	cloudformation.New(aws.GetAuth(nil, nil, nil, nil), aws.USEast)
}

func New() {
	return verServ
}

type versionServ struct {
}

type VersionServ interface {
}
