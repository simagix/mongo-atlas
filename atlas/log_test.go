// Copyright 2019 Kuei-chun Chen. All rights reserved.

package atlas

import (
	"os"
	"testing"
)

func TestParseLogURI(t *testing.T) {
	var err error
	publicKey := os.Getenv("ATLAS_USER")
	privateKey := os.Getenv("ATLAS_KEY")
	groupID := os.Getenv("ATLAS_GROUP")
	uri := "atlas://" + publicKey + ":" + privateKey + "@" + groupID + "/keyhole"
	if _, err = ParseLogURI(uri); err != nil {
		t.Fatal(err)
	}
}
