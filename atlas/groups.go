// Copyright 2019 Kuei-chun Chen. All rights reserved.

package atlas

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/simagix/gox"
)

// GetGroups get processes of a user
func (api *API) GetGroups() (map[string]interface{}, error) {
	var err error
	var doc map[string]interface{}
	var b []byte

	uri := BaseURL + "/groups"
	if b, err = api.GET(uri, ApplicationJSON); err != nil {
		return doc, err
	}
	json.Unmarshal(b, &doc)
	if api.verbose == true {
		fmt.Println(gox.Stringify(doc, "", "  "))
	}
	return doc, err
}

// GetGroupsByID get processes of a user
func (api *API) GetGroupsByID(groupID string) (map[string]interface{}, error) {
	var err error
	var doc map[string]interface{}
	var b []byte

	uri := BaseURL + "/groups/" + groupID
	if b, err = api.GET(uri, ApplicationJSON); err != nil {
		return nil, err
	}
	json.Unmarshal(b, &doc)
	return doc, err
}

// GetGroupIDs returns an array of group IDs
func (api *API) GetGroupIDs() ([]string, error) {
	var groupIDs []string
	var err error
	var doc map[string]interface{}
	if doc, err = api.GetGroups(); err != nil {
		return groupIDs, err
	}
	_, ok := doc["results"]
	if !ok {
		return groupIDs, errors.New(gox.Stringify(doc))
	}
	results := doc["results"].([]interface{})
	for _, result := range results {
		groupIDs = append(groupIDs, result.(map[string]interface{})["id"].(string))
	}
	return groupIDs, err
}
