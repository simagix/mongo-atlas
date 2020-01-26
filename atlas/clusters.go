// Copyright 2019 Kuei-chun Chen. All rights reserved.

package atlas

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/simagix/gox"
)

// GetClusters gets clusters by a group
func (api *API) GetClusters(groupID string) (map[string]interface{}, error) {
	var err error
	var doc map[string]interface{}
	var b []byte

	uri := BaseURL + "/groups/" + groupID + "/clusters"
	if b, err = api.GET(uri, ApplicationJSON); err != nil {
		return nil, err
	}
	json.Unmarshal(b, &doc)
	if api.verbose == true {
		fmt.Println(gox.Stringify(doc, "", "  "))
	}
	return doc, err
}

// Resume pauses a cluster
func (api *API) Resume() (string, error) {
	return api.pause(false)
}

// Pause pauses a cluster
func (api *API) Pause() (string, error) {
	return api.pause(true)
}

// pause pauses a cluster
func (api *API) pause(paused bool) (string, error) {
	var err error
	var doc map[string]interface{}
	var b []byte

	if api.groupID == "" || api.clusterName == "" {
		return "", errors.New("invalid format ([atlas://]publicKey:privateKey[@group/cluster])")
	}
	uri := BaseURL + "/groups/" + api.groupID + "/clusters/" + api.clusterName
	body := []byte(`{ "paused": ` + strconv.FormatBool(paused) + ` }`)
	if b, err = api.PATCH(uri, ApplicationJSON, body); err != nil {
		return "", err
	}
	json.Unmarshal(b, &doc)
	return gox.Stringify(doc, "", "  "), err
}
