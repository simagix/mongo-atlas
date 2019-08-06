// Copyright 2019 Kuei-chun Chen. All rights reserved.

package atlas

import (
	"os"
	"testing"
)

func TestNewKey(t *testing.T) {
	publicKey := os.Getenv("ATLAS_USER")
	privateKey := os.Getenv("ATLAS_KEY")
	api := NewKey(publicKey, privateKey)
	if api.publicKey != publicKey || api.privateKey != privateKey {
		t.Fatal("parsing error")
	}
}

func TestParseURI(t *testing.T) {
	var err error
	var api *API
	publicKey := os.Getenv("ATLAS_USER")
	privateKey := os.Getenv("ATLAS_KEY")
	if api, err = ParseURI(os.Getenv("ATLAS_AUTH")); err != nil {
		t.Fatal(err)
	}
	if api.publicKey != publicKey || api.privateKey != privateKey {
		t.Fatal(publicKey, privateKey, os.Getenv("ATLAS_AUTH"))
	}
}

func TestSetVerbose(t *testing.T) {
	publicKey := os.Getenv("ATLAS_USER")
	privateKey := os.Getenv("ATLAS_KEY")
	api := NewKey(publicKey, privateKey)
	api.SetVerbose(true)
	if api.verbose != true {
		t.Fatal("SetVerbose failed")
	}
}

func TestGET(t *testing.T) {
	publicKey := os.Getenv("ATLAS_USER")
	privateKey := os.Getenv("ATLAS_KEY")
	groupID := os.Getenv("ATLAS_GROUP")
	api := NewKey(publicKey, privateKey)
	uri := BaseURL + "/groups/" + groupID + "/clusters"
	if _, err := api.GET(uri, ApplicationJSON); err != nil {
		t.Fatal(err)
	}
}
