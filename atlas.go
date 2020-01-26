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

var version = "self-built"

func main() {
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
	if *loginfo == true {
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
		os.Exit(0)
	}

	if api, err = atlas.ParseURI(flag.Arg(0)); err != nil {
		log.Fatal(err)
	}
	api.SetVerbose(*verbose)
	var str string
	if *info == true {
		if str, err = api.GetClustersSummary(); err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	} else if *ftdc == true {
		if str, err = api.DownloadFTDC(); err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	} else if *resume == true {
		if str, err = api.Do("PATCH", `{ "paused": false }`); err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	} else if *pause == true {
		if str, err = api.Do("PATCH", `{ "paused": true }`); err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	} else if *request != "" {
		data := "{}"
		if len(flag.Args()) > 1 {
			data = flag.Arg(1)
		}
		if str, err = api.Do(*request, data); err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	}
}
