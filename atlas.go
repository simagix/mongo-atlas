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
	clusters := flag.Bool("clusters", false, "list atlas clusters")
	logs := flag.Bool("logs", false, "download and analyze logs")
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
	if *clusters == true {
		if api, err = atlas.ParseURI(flag.Arg(0)); err != nil {
			log.Fatal(err)
		}
		api.SetVerbose(*verbose)
		var str string
		if str, err = api.GetClustersSummary(); err != nil {
			log.Fatal(err)
		}
		fmt.Println(str)
	} else if *logs == true {
		var atl *atlas.Log
		if atl, err = atlas.ParseLogURI(flag.Arg(0)); err != nil {
			log.Fatal(err)
		}
		atl.SetVerbose(*verbose)
		var filenames []string
		if filenames, err = atl.Download(); err != nil {
			log.Fatal(err)
		}
		fmt.Println(strings.Join(filenames, "\n"))
	}
}
