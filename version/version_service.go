package version

import (
	"encoding/json"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/cloudformation"
	"github.com/hashicorp/consul/api"
	"io"
	"log"
	"time"
)

var (
	s VersionServ
)

type versionServ struct {
	cf          *cloudformation.CloudFormation
	environment string
}

type VersionServ interface {
	FindVersions(names []string) ([]*version, error)
	EncodeVersions(w io.Writer, versions []*version) error
}

func NewVersionServ() VersionServ {
	if s == nil {
		s = newVersionServ()
	}
	return s
}

func newConsulConfig() *api.Config {
	return &api.Config{
		Address: "consul:8500",
		Scheme:  "http"}
}

func newVersionServ() VersionServ {
	client, err := api.NewClient(newConsulConfig())
	if err != nil {
		log.Fatal("Could not configure consul api.")
	}

	kvpair, _, err := client.KV().Get("environment", nil)
	if err != nil {
		log.Fatal("Could not get the environment key from consul.")
	}

	auth, err := aws.GetAuth("", "", "", time.Now())
	if err != nil {
		log.Fatal("Could not find AWS Credentials")
	}

	cfn := cloudformation.New(auth, aws.USEast)
	return &versionServ{cf: cfn, environment: string(kvpair.Value)}
}

func (s *versionServ) EncodeVersions(w io.Writer, versions []*version) error {
	log.Print("Encoding Version Objects to JSON")
	return json.NewEncoder(w).Encode(versions)
}

func (s *versionServ) FindVersions(names []string) ([]*version, error) {
	versions := make([]*version, 0)
	for _, name := range names {
		stack_name := s.environment + "-" + name
		log.Printf("DescribingStack %s", stack_name)
		response, err := s.cf.DescribeStacks(stack_name, "")
		if err == nil {
			outputs := make(map[string]string)
			for _, output := range response.Stacks[0].Outputs {
				outputs[output.OutputKey] = output.OutputValue
			}
			versions = append(versions, NewVersion(outputs))
		} else {
			log.Print(err)
		}
	}
	return versions, nil
}
