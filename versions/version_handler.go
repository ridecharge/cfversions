package versions

import (
	"net/http"
)

type versionHandler struct {
	s versionService
}

func RegisterHandlers() {
	http.Handle("/versions", &versionHandler{s: newVersionService()})
	http.HandleFunc("/health", health)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (h *versionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		names    []string
		versions []*version
		err      error
	)
	names = r.URL.Query()["names"]
	if versions, err = h.s.findVersions(names); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	h.s.encodeVersions(w, versions)
}
