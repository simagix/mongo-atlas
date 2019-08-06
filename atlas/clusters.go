// Copyright 2019 Kuei-chun Chen. All rights reserved.

package atlas

import (
	"encoding/json"
	"fmt"

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
