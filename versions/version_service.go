package versions

import (
	"encoding/json"
	"github.com/goamz/goamz/autoscaling"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/cloudformation"
	"github.com/hashicorp/consul/api"
	"io"
	"log"
	"os"
	"time"
)

var (
	s VersionServ
)

type versionServ struct {
	cf  *cloudformation.CloudFormation
	as  *autoscaling.AutoScaling
	env string
}

type VersionServ interface {
	FindVersions(names []string) ([]*version, error)
	FindAllVersionNames() ([]string, error)
	EncodeVersions(w io.Writer, versions []*version) error
}

func NewVersionServ() VersionServ {
	if s == nil {
		s = newVersionServ()
	}
	return s
}

func newConsulConfig() *api.Config {
	address := os.Getenv("CONSUL_PORT_8500_TCP_ADDR") +
		":" + os.Getenv("CONSUL_PORT_8500_TCP_PORT")
	return &api.Config{
		Address: address,
		Scheme:  "http"}
}

func newVersionServ() VersionServ {
	client, err := api.NewClient(newConsulConfig())
	if err != nil {
		log.Fatal("Could not configure consul api.", err)
	}

	kvpair, _, err := client.KV().Get("environment", nil)
	if err != nil || kvpair == nil {
		log.Fatal("Could not get the environment key from consul.")
	}

	auth, err := aws.GetAuth("", "", "", time.Now())
	if err != nil {
		log.Fatal("Could not find AWS Credentials")
	}

	env := string(kvpair.Value)
	cf := cloudformation.New(auth, aws.USEast)
	as := autoscaling.New(auth, aws.USEast)
	return &versionServ{cf: cf, env: env, as: as}
}

func (s *versionServ) EncodeVersions(w io.Writer, versions []*version) error {
	log.Print("Encoding Version Objects to JSON")
	return json.NewEncoder(w).Encode(versions)
}

func (s *versionServ) FindAllVersionNames() ([]string, error) {
	filter := autoscaling.NewFilter()
	filter.Add("key", "Role")
	r, err := s.as.DescribeTags(filter, 0, "")
	if err != nil {
		return nil, err
	}

	numTags := len(r.Tags)
	namesSet := make(map[string]bool, numTags)
	names := make([]string, 0, numTags)
	for _, tag := range r.Tags {
		name := tag.Value
		if namesSet[name] == false &&
			name != "ntp" && name != "nat" && name != "bastion" {
			namesSet[name] = true
			names = append(names, name)
		}
	}
	return names, nil
}

func (s *versionServ) FindVersions(names []string) ([]*version, error) {
	versions := make([]*version, 0, len(names))
	for _, name := range names {
		stack_name := s.env + "-" + name
		log.Printf("DescribingStack %s", stack_name)
		response, err := s.cf.DescribeStacks(stack_name, "")
		if err == nil {
			outputs := make(map[string]string,
				len(response.Stacks[0].Outputs))
			for _, output := range response.Stacks[0].Outputs {
				outputs[output.OutputKey] = output.OutputValue
			}
			versions = append(versions, NewVersion(outputs))
		} else {
			if err.(*cloudformation.Error).StatusCode == 400 {
				log.Print(err)
			} else {
				return nil, err
			}
		}
	}
	return versions, nil
}
