// Copyright 2019 Kuei-chun Chen. All rights reserved.

package atlas

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/simagix/gox"
)

// GetClustersSummary retrieve clusters summary for a user
func (api *API) GetClustersSummary() (string, error) {
	var err error
	var doc map[string]interface{}
	var buffers []string
	var b []byte

	uri := BaseURL + "/users/byName/" + api.publicKey
	if b, err = api.GET(uri, ApplicationJSON); err != nil {
		return "", err
	}
	json.Unmarshal(b, &doc)
	if api.verbose == true {
		fmt.Println(gox.Stringify(doc, "", "  "))
	}

	roles, ok := doc["roles"]
	if !ok {
		return "", errors.New(gox.Stringify(doc))
	}

	for _, role := range roles.([]interface{}) {
		groupID := role.(map[string]interface{})["groupId"]
		if groupID != nil {
			var clusters []interface{}
			if clusters, err = api.getClusters(groupID.(string)); err != nil {
				return "", err
			}
			var processes []interface{}
			if processes, err = api.getProcesses(groupID.(string)); err != nil {
				return "", err
			}
			buffers = append(buffers, fmt.Sprint("- Group: ", groupID))
			for _, cluster := range clusters {
				m := cluster.(map[string]interface{})
				buffers = append(buffers, fmt.Sprint("  - cluster name: ", m["name"]))
				buffers = append(buffers, fmt.Sprint("    - ", m["mongoDBVersion"], ", ", m["clusterType"], ", ", m["srvAddress"]))
				buffers = append(buffers, fmt.Sprint("    - Hosts:"))
				for _, process := range processes {
					maps := process.(map[string]interface{})
					if strings.Index(strings.ToLower(maps["hostname"].(string)), strings.ToLower(m["name"].(string)+"-")) == 0 {
						buffers = append(buffers, fmt.Sprint("      - ", maps["hostname"], " (", maps["typeName"], ")"))
					}
				}
			}
			buffers = append(buffers, "\n")
		}
	}
	return strings.Join(buffers, "\n"), err
}

func (api *API) getClusters(groupID string) ([]interface{}, error) {
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
	results, ok := doc["results"]
	if !ok {
		return nil, errors.New(gox.Stringify(doc))
	}
	return results.([]interface{}), err
}

func (api *API) getProcesses(groupID string) ([]interface{}, error) {
	var err error
	var doc map[string]interface{}
	var b []byte

	uri := BaseURL + "/groups/" + groupID + "/processes"
	if b, err = api.GET(uri, ApplicationJSON); err != nil {
		return nil, err
	}
	json.Unmarshal(b, &doc)
	if api.verbose == true {
		fmt.Println(gox.Stringify(doc, "", "  "))
	}
	results, ok := doc["results"]
	if !ok {
		return nil, errors.New(gox.Stringify(doc))
	}
	return results.([]interface{}), err
}
