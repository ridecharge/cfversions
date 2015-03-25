package versions

import (
	"encoding/json"
	"github.com/hashicorp/consul/api"
	"io"
	"log"
)

type versionServ struct {
	c            *api.Client
	env          string
	versionNames []string
}

type versionService interface {
	findVersions(names []string) ([]*version, error)
	encodeVersions(w io.Writer, versions []*version) error
}

func newVersionService() versionService {
	client, err := api.NewClient(&api.Config{
		Address: "192.168.59.103:8500",
		Scheme:  "http"})
	if err != nil {
		log.Fatal("Could not configure consul api.", err)
	}

	kvpair, _, err := client.KV().Get("cf/env/environment", nil)
	if err != nil || kvpair == nil {
		log.Fatal("Could not get the environment key from consul.")
	}
	vs := &versionServ{
		c:   client,
		env: string(kvpair.Value)}
	go vs.watchVersionNames()
	return vs
}

func (s *versionServ) findVersionNames(index uint64) ([]string, uint64, error) {
	qo := &api.QueryOptions{WaitIndex: index}
	names, meta, err := s.c.KV().Keys("cf/", "/", qo)
	return names, meta.LastIndex, err
}

func (s *versionServ) watchVersionNames() {
	var (
		index uint64
		err   error
	)
	index = 0
	for {
		s.versionNames, index, err = s.findVersionNames(index)
		if err != nil {
			log.Fatal("Error watching versionNames.", err)
		}
	}
}

func (s *versionServ) encodeVersions(w io.Writer, versions []*version) error {
	log.Print("Encoding Version Objects to JSON")
	return json.NewEncoder(w).Encode(versions)
}

func (s *versionServ) findVersions(names []string) ([]*version, error) {
	vn := names
	if len(vn) == 0 {
		vn = s.versionNames
	}
	vs := make([]*version, 0, len(names))
	for _, name := range vn {
		println("key_prefix: ", name)
		kvs, _, err := s.c.KV().List(name, nil)
		if err != nil {
			return nil, err
		}
		if kvs == nil {
			continue
		}

		props := make(map[string]string, len(kvs))
		for _, kv := range kvs {
			println(kv.Key, string(kv.Value))
			props[kv.Key] = string(kv.Value)
		}
		if props[name+"application_name"] != "" {
			vs = append(vs, newVersion(props, name))
		}
	}
	return vs, nil
}
