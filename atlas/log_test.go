// Copyright 2019 Kuei-chun Chen. All rights reserved.

package atlas

import (
	"os"
	"testing"
)

func TestParseLogURI(t *testing.T) {
	var err error
	publicKey := os.Getenv("ATLAS_PUB")
	privateKey := os.Getenv("ATLAS_PRI")
	groupID := os.Getenv("ATLAS_GROUP")
	uri := "atlas://" + publicKey + ":" + privateKey + "@" + groupID + "/keyhole"
	if _, err = ParseLogURI(uri); err != nil {
		t.Fatal(err)
	}
}

func TestDownload(t *testing.T) {
	var err error
	var atl *Log
	var filenames []string
	publicKey := os.Getenv("ATLAS_PUB")
	privateKey := os.Getenv("ATLAS_PRI")
	groupID := os.Getenv("ATLAS_GROUP")
	uri := "atlas://" + publicKey + ":" + privateKey + "@" + groupID + "/keyhole"
	if atl, err = ParseLogURI(uri); err != nil {
		t.Fatal(err)
	}
	atl.SetVerbose(testing.Verbose())
	if filenames, err = atl.Download(); err != nil {
		t.Fatal(err)
	}
	for _, filename := range filenames {
		os.Remove(filename)
	}
}
