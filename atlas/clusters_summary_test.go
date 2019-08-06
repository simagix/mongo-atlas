// Copyright 2019 Kuei-chun Chen. All rights reserved.

package atlas

import (
	"os"
	"testing"

	"github.com/simagix/gox"
)

func TestGetClustersSummary(t *testing.T) {
	publicKey := os.Getenv("ATLAS_USER")
	privateKey := os.Getenv("ATLAS_KEY")
	api := NewKey(publicKey, privateKey)
	if str, err := api.GetClustersSummary(); err != nil {
		t.Fatal(err)
	} else {
		t.Log(str)
	}
}

func TestGetClusters(t *testing.T) {
	publicKey := os.Getenv("ATLAS_USER")
	privateKey := os.Getenv("ATLAS_KEY")
	groupID := os.Getenv("ATLAS_GROUP")
	api := NewKey(publicKey, privateKey)
	if doc, err := api.getClusters(groupID); err != nil {
		t.Fatal(err)
	} else {
		t.Log(gox.Stringify(doc, "", "  "))
	}
}

func TestGetProcesses(t *testing.T) {
	publicKey := os.Getenv("ATLAS_USER")
	privateKey := os.Getenv("ATLAS_KEY")
	groupID := os.Getenv("ATLAS_GROUP")
	api := NewKey(publicKey, privateKey)
	if doc, err := api.getProcesses(groupID); err != nil {
		t.Fatal(err)
	} else {
		t.Log(gox.Stringify(doc, "", "  "))
	}
}
