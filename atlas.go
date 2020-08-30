// Copyright 2019 Kuei-chun Chen. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/simagix/keyhole/atlas"
)

var version = "self-built"

func main() {
	alertsFile := flag.String("addAlerts", "", "add all alerts")
	alerts := flag.Bool("alerts", false, "get all alerts")
	ftdc := flag.Bool("ftdc", false, "download FTDC")
	info := flag.Bool("info", false, "list Atlas clusters")
	loginfo := flag.Bool("loginfo", false, "download mongo logs of a group")
	pause := flag.Bool("pause", false, "pause a cluster")
	request := flag.String("request", "", "HTTP command")
	resume := flag.Bool("resume", false, "resume a cluster")
	verbose := flag.Bool("v", false, "verbose")
	ver := flag.Bool("version", false, "print version number")

	flag.Parse()
	if *ver {
		fmt.Println("matlas", version)
		os.Exit(0)
	} else if len(flag.Args()) == 0 {
		fmt.Println("Usage: atlas [flags] uri")
		os.Exit(1)
	}
	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })
	var err error
	var api *atlas.API
	if api, err = atlas.ParseURI(flag.Arg(0)); err != nil {
		log.Fatal(err)
	}
	api.SetAlertsFile(*alertsFile)
	api.SetAlerts(*alerts)
	api.SetArgs(flag.Args())
	api.SetFTDC(*ftdc)
	api.SetInfo(*info)
	api.SetPause(*pause)
	api.SetResume(*resume)
	api.SetLoginfo(*loginfo)
	api.SetRequest(*request)
	api.SetVerbose(*verbose)
	fmt.Println(api.Execute())
}
