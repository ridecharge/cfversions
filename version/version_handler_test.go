package version

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockVersionServ struct {
	EncodeVersionsCalled      bool
	FindAllVersionNamesCalled bool
	FindVersionsCalled        bool
}

type errVersionServ struct {
	mockVersionServ
}

type notFoundVersionServ struct {
	mockVersionServ
}

func (s *notFoundVersionServ) FindVersions(names []string) ([]*version, error) {
	return make([]*version, 0, 0), nil
}

func (s *errVersionServ) FindVersions(names []string) ([]*version, error) {
	return nil, errors.New("error")
}

func (s *errVersionServ) FindAllVersionNames() ([]string, error) {
	return nil, errors.New("error")
}

func (s *mockVersionServ) EncodeVersions(w io.Writer, versions []*version) error {
	s.EncodeVersionsCalled = true
	return nil
}

func (s *mockVersionServ) FindAllVersionNames() ([]string, error) {
	s.FindAllVersionNamesCalled = true
	return nil, nil
}

func (s *mockVersionServ) FindVersions(names []string) ([]*version, error) {
	s.FindVersionsCalled = true
	return []*version{&version{}}, nil
}

func TestHealth(t *testing.T) {
	r := httptest.NewRecorder()
	health(r, nil)
	if r.Code != 200 {
		t.Error("health check should return code 200")
	}
}

func TestServeHTTPSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/versions?name=test", nil)
	s := &mockVersionServ{}
	h := &versionHandler{s: s}
	h.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Error("HTTP 200 should be returned on success.")
	}

	if w.HeaderMap["Content-Type"][0] != "application/json" {

		t.Error("Content-Type should be application/json.")
	}

	if !s.EncodeVersionsCalled {
		t.Error("EncodeVersions should be called on success.")
	}
}

func TestServeHTTPFindVersionErr(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/versions", nil)
	s := &errVersionServ{}
	h := &versionHandler{s: s}
	h.ServeHTTP(w, r)

	if w.Code != 500 {
		t.Error("HTTP 500 should be returned on errors.")
	}
}

func TestServeHTTPNoNames(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/versions", nil)
	s := &mockVersionServ{}
	h := &versionHandler{s: s}
	h.ServeHTTP(w, r)

	if !s.FindAllVersionNamesCalled {
		t.Error("FindAllVersionNames should be called when no names param.")
	}
}

func TestServeHTTPNoNamesErr(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/versions", nil)
	s := &errVersionServ{}
	h := &versionHandler{s: s}
	h.ServeHTTP(w, r)

	if w.Code != 500 {
		t.Error("HTTP 500 should be returned on errors.")
	}
}
