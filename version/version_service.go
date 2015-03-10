package version

import (
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/cloudformation"
	"log"
	"time"
)

var (
	verServ VersionServ
)

func init() {
	verServ = &versionServ{}
	auth, err := aws.GetAuth("", "", "", time.Now())
	if err != nil {
		log.Panic("Could not find AWS Credentials")
	}
	cloudformation.New(auth, aws.USEast)
}

func NewVersionServ() VersionServ {
	return verServ
}

type versionServ struct {
}

type VersionServ interface {
}
