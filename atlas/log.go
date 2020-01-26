// Copyright 2019 Kuei-chun Chen. All rights reserved.

package atlas

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/simagix/gox"
)

// Log stores Atlas logs API info
type Log struct {
	apiKey      API
	clusterName string
	groupID     string
	startTime   time.Time
	endTime     time.Time
	verbose     bool
}

// ParseLogURI returns API struct from a URI
func ParseLogURI(uri string) (*Log, error) {
	if strings.HasPrefix(uri, "atlas://") == true {
		uri = uri[8:]
	}
	matched := regexp.MustCompile(`^(\S+):(\S+)@([^\/]+)\/([^\/]+)(\/([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01])))?`)
	if matched.MatchString(uri) == false {
		return nil, errors.New(`Incorrect format, should be "[atlas://]publicKey:privateKey@groupID/clusterName[/yyyy-mm-dd]"`)
	}
	results := matched.FindStringSubmatch(uri)
	endTime := time.Now()
	startTime := endTime.Add(time.Hour * -24)
	if len(results) > 6 && results[6] != "" {
		endTime, _ = time.Parse("2006-01-02", results[6])
		startTime = endTime.Add(time.Hour * -24)
		endTime = endTime.Add(time.Hour * 24)
	}
	return &Log{apiKey: API{publicKey: results[1], privateKey: results[2]}, groupID: results[3],
		clusterName: results[4], startTime: startTime, endTime: endTime}, nil
}

// SetVerbose sets verbose
func (atl *Log) SetVerbose(verbose bool) {
	atl.verbose = verbose
}

// Download downloads logs
func (atl *Log) Download() ([]string, error) {
	var err error
	var filenames []string
	var api *API
	var doc map[string]interface{}

	api = NewKey(atl.apiKey.publicKey, atl.apiKey.privateKey)
	if doc, err = api.GetProcesses(atl.groupID); err != nil {
		return filenames, err
	}
	_, ok := doc["results"]
	if !ok {
		return filenames, errors.New(gox.Stringify(doc))
	}
	processes := doc["results"]
	log.Println("download files from", atl.startTime.Format("2006.01.02 15:04:05"), "to", atl.endTime.Format("2006.01.02 15:04:05"))
	for _, process := range processes.([]interface{}) {
		maps := process.(map[string]interface{})
		if strings.Index(strings.ToLower(maps["hostname"].(string)), strings.ToLower(atl.clusterName+"-")) == 0 &&
			strings.Index(maps["typeName"].(string), "REPLICA_") == 0 {
			hostname := maps["hostname"].(string)
			filename := "./mongodb.log." + hostname + ".gz"
			uri := BaseURL + "/groups/" + atl.groupID + "/clusters/" + hostname + "/logs/mongodb.gz"
			uri += "?startDate=" + fmt.Sprintf("%v", atl.startTime.Unix()) + "&endDate=" + fmt.Sprintf("%v", atl.endTime.Unix())
			if atl.verbose {
				log.Println("download from", uri)
			}
			var b []byte
			api.SetAcceptType(ApplicationGZip)
			if b, err = api.Get(uri); err != nil {
				log.Println(err)
				continue
			}
			if len(b) > 0 {
				if err = ioutil.WriteFile(filename, b, 0644); err != nil {
					log.Println(err)
					continue
				}
				filenames = append(filenames, filename)
			} else {
				log.Println("No content, skipped", hostname)
			}
		}
	}
	return filenames, err
}
