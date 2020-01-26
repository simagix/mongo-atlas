// Copyright 2019 Kuei-chun Chen. All rights reserved.

package atlas

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/simagix/gox"
)

// BaseURL -
const BaseURL = "https://cloud.mongodb.com/api/atlas/v1.0"

// ApplicationJSON -
const ApplicationJSON = "application/json"

// ApplicationGZip -
const ApplicationGZip = "application/gzip"

// API stores Atlas API key
type API struct {
	publicKey   string
	privateKey  string
	groupID     string
	clusterName string
	verbose     bool
}

// NewKey returns API struct
func NewKey(publicKey string, privateKey string) *API {
	return &API{publicKey: publicKey, privateKey: privateKey}
}

// ParseURI returns API struct from a URI
func ParseURI(uri string) (*API, error) {
	api := &API{}
	if strings.HasPrefix(uri, "atlas://") == true {
		uri = uri[8:]
	}
	i := strings.Index(uri, "@")
	if i > 0 {
		tailer := strings.Split(uri[i+1:], "/")
		api.groupID = tailer[0]
		api.clusterName = tailer[1]
		uri = uri[:i]
	}
	i = strings.LastIndex(uri, ":")
	if i < 0 {
		return nil, errors.New("invalid format ([atlas://]publicKey:privateKey[@group/cluster])")
	}
	api.publicKey = uri[:i]
	api.privateKey = uri[i+1:]
	return api, nil
}

// SetVerbose sets verbose
func (api *API) SetVerbose(verbose bool) {
	api.verbose = verbose
}

// GET performs HTTP GET function
func (api *API) GET(uri string, accept string) ([]byte, error) {
	var err error
	var resp *http.Response
	var b []byte

	headers := map[string]string{}
	headers["Content-Type"] = ApplicationJSON
	headers["Accept"] = accept
	if resp, err = gox.HTTPDigest("GET", uri, api.publicKey, api.privateKey, headers); err != nil {
		return b, err
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	return b, err
}

// PATCH performs HTTP PATCH function
func (api *API) PATCH(uri string, accept string, body []byte) ([]byte, error) {
	var err error
	var resp *http.Response
	var b []byte

	headers := map[string]string{}
	headers["Content-Type"] = ApplicationJSON
	headers["Accept"] = accept
	if resp, err = gox.HTTPDigest("PATCH", uri, api.publicKey, api.privateKey, headers, body); err != nil {
		return b, err
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	return b, err
}
