package versions

import (
	"testing"
)

func TestNewVersion(t *testing.T) {
	props := map[string]string{
		"version":           "Version",
		"application_name":  "ApplicationName",
		"private_end_point": "PrivateEndPoint",
		"public_end_point":  "PublicEndPoint",
	}

	version := newVersion(props, "")

	if version.Version != "Version" {
		t.Error("Should set Version from map.")
	}
	if version.ApplicationName != "ApplicationName" {
		t.Error("Should set ApplicationName from map.")
	}
	if version.PrivateEndPoint != "PrivateEndPoint" {
		t.Error("Should set PrivateEndPoint from map.")
	}
	if version.PublicEndPoint != "PublicEndPoint" {
		t.Error("Should set PublicEndPoint from map.")
	}
}
