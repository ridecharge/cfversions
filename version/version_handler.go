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
	names := make([]string, 0)
	if len(r.URL.Query()["names"]) > 0 {
		names = r.URL.Query()["names"]
	}

	versions, err := h.s.FindVersions(names)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	if versions == nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	h.s.EncodeVersions(w, versions)
}
