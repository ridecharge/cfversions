package version

import (
	"net/http"
)

type versionHandler struct {
	s VersionServ
}

func RegisterHandlers() {
	http.Handle("/versions", &versionHandler{s: NewVersionServ()})
	http.HandleFunc("/health", health)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (h *versionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		names    []string
		err      error
		versions []*version
	)

	if len(r.URL.Query()["names"]) > 0 {
		names = r.URL.Query()["names"]
	} else {
		if names, err = h.s.FindAllVersionNames(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if versions, err = h.s.FindVersions(names); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(versions) < 1 {
		http.NotFound(w, r)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	h.s.EncodeVersions(w, versions)
}
