// Copyright 2019 Kuei-chun Chen. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/simagix/mongo-atlas/atlas"
)

func main() {
	info := flag.Bool("info", false, "list atlas clusters")
	loginfo := flag.Bool("loginfo", false, "download mongo logs of a group")
	pause := flag.Bool("pause", false, "pause a cluster")
	resume := flag.Bool("resume", false, "resume a cluster")
	verbose := flag.Bool("v", false, "verbose")

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Usage: atlas [flags] uri")
		os.Exit(1)
	}
	flagset := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { flagset[f.Name] = true })
	var err error
	var api *atlas.API
	if *info == true {
		if api, err = atlas.ParseURI(flag.Arg(0)); err != nil {
			log.Fatal(err)
		}
		api.SetVerbose(*verbose)
		var str string
		if str, err = api.GetClustersSummary(); err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	} else if *loginfo == true {
		var atl *atlas.Log
		if atl, err = atlas.ParseLogURI(flag.Arg(0)); err != nil {
			log.Fatal(err)
		}
		atl.SetVerbose(*verbose)
		var filenames []string
		if filenames, err = atl.Download(); err != nil {
			log.Fatal(err)
		}
		if len(filenames) > 0 {
			fmt.Println("Files downloaded:")
			fmt.Println("\t", strings.Join(filenames, "\n\t "))
		}
	} else if *resume == true {
		if api, err = atlas.ParseURI(flag.Arg(0)); err != nil {
			log.Fatal(err)
		}
		api.SetVerbose(*verbose)
		var str string
		if str, err = api.Resume(); err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	} else if *pause == true {
		if api, err = atlas.ParseURI(flag.Arg(0)); err != nil {
			log.Fatal(err)
		}
		api.SetVerbose(*verbose)
		var str string
		if str, err = api.Pause(); err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	}
}
