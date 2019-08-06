// Copyright 2019 Kuei-chun Chen. All rights reserved.

package atlas

import (
	"errors"
	"fmt"
	"strings"

	"github.com/simagix/gox"
)

// GetClustersSummary retrieve clusters summary for a user
func (api *API) GetClustersSummary() (string, error) {
	var err error
	var buffers []string
	var groupIDs []string
	var doc map[string]interface{}

	if groupIDs, err = api.GetGroupIDs(); err != nil {
		return "", err
	}

	for _, groupID := range groupIDs {
		if doc, err = api.GetClusters(groupID); err != nil {
			return "", err
		}
		_, ok := doc["results"]
		if !ok {
			return "", errors.New(gox.Stringify(doc))
		}
		clusters := doc["results"]
		if doc, err = api.GetProcesses(groupID); err != nil {
			return "", err
		}
		_, ok = doc["results"]
		if !ok {
			return "", errors.New(gox.Stringify(doc))
		}
		processes := doc["results"]
		buffers = append(buffers, fmt.Sprint("- Group: ", groupID))
		for _, cluster := range clusters.([]interface{}) {
			m := cluster.(map[string]interface{})
			buffers = append(buffers, fmt.Sprint("  - cluster name: ", m["name"]))
			buffers = append(buffers, fmt.Sprint("    - ", m["mongoDBVersion"], ", ", m["clusterType"], ", ", m["srvAddress"]))
			buffers = append(buffers, fmt.Sprint("    - Hosts:"))
			for _, process := range processes.([]interface{}) {
				maps := process.(map[string]interface{})
				if strings.Index(strings.ToLower(maps["hostname"].(string)), strings.ToLower(m["name"].(string)+"-")) == 0 {
					buffers = append(buffers, fmt.Sprint("      - ", maps["hostname"], " (", maps["typeName"], ")"))
				}
			}
		}
		buffers = append(buffers, "\n")
	}
	return strings.Join(buffers, "\n"), err
}
