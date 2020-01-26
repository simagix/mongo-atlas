// Copyright 2019 Kuei-chun Chen. All rights reserved.

package atlas

import (
	"encoding/json"
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
	acceptType  string
	contentType string
	clusterName string
	groupID     string
	params      []string
	privateKey  string
	publicKey   string
	verbose     bool
}

// NewKey returns API struct
func NewKey(publicKey string, privateKey string) *API {
	return &API{publicKey: publicKey, privateKey: privateKey, contentType: ApplicationJSON, acceptType: ApplicationJSON}
}

// ParseURI returns API struct from a URI
func ParseURI(uri string) (*API, error) {
	api := &API{contentType: ApplicationJSON, acceptType: ApplicationJSON, params: []string{}}
	if strings.HasPrefix(uri, "atlas://") == true {
		uri = uri[8:]
	}
	i := strings.Index(uri, "@")
	if i > 0 {
		tailer := uri[i+1:]
		if q := strings.Index(tailer, "?"); q > 0 {
			api.params = strings.Split(tailer[q+1:], "&")
			tailer = tailer[:q]
		}
		if n := strings.Index(tailer, "/"); n > 0 {
			api.groupID = tailer[:n]
			api.clusterName = tailer[n+1:]
		} else {
			api.groupID = tailer
		}
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

// SetAcceptType sets acceptType
func (api *API) SetAcceptType(acceptType string) {
	api.acceptType = acceptType
}

// SetContentType sets contentType
func (api *API) SetContentType(contentType string) {
	api.contentType = contentType
}

// SetVerbose sets verbose
func (api *API) SetVerbose(verbose bool) {
	api.verbose = verbose
}

// Get performs HTTP GET function
func (api *API) Get(uri string) ([]byte, error) {
	var err error
	var resp *http.Response
	var b []byte

	headers := map[string]string{}
	headers["Content-Type"] = api.contentType
	headers["Accept"] = api.acceptType
	if resp, err = gox.HTTPDigest("GET", uri, api.publicKey, api.privateKey, headers); err != nil {
		return b, err
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	return b, err
}

// Patch performs HTTP PATCH function
func (api *API) Patch(uri string, body []byte) ([]byte, error) {
	var err error
	var resp *http.Response
	var b []byte

	headers := map[string]string{}
	headers["Content-Type"] = api.contentType
	headers["Accept"] = api.acceptType
	if resp, err = gox.HTTPDigest("PATCH", uri, api.publicKey, api.privateKey, headers, body); err != nil {
		return b, err
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	return b, err
}

// Do execute a command
func (api *API) Do(method string, data string) (string, error) {
	var err error
	var resp *http.Response
	var doc map[string]interface{}
	var b []byte

	if api.groupID == "" {
		return "", errors.New("invalid format ([atlas://]publicKey:privateKey@group)")
	}
	uri := BaseURL + "/groups/" + api.groupID + "/clusters"
	if api.clusterName != "" {
		uri += "/" + api.clusterName
	}
	body := []byte(data)
	headers := map[string]string{}
	headers["Content-Type"] = api.contentType
	headers["Accept"] = api.acceptType
	if resp, err = gox.HTTPDigest(method, uri, api.publicKey, api.privateKey, headers, body); err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	json.Unmarshal(b, &doc)
	return gox.Stringify(doc, "", "  "), err
}
