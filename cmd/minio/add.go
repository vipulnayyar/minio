package main

import (
	"log"
	"os"
	"path"

	"github.com/codegangsta/cli"
)

func add(c *cli.Context) {
	config, err := parseInput(c)
	if err != nil {
		log.Fatal(err)
	}
	var filePath, objectName string
	switch len(c.Args()) {
	case 1:
		objectName = path.Base(c.Args().Get(0))
		filePath = c.Args().Get(0)
	case 2:
		objectName = c.Args().Get(0)
		filePath = c.Args().Get(1)
	default:
		log.Fatal("Please specify a valid object name \n # erasure-demo put [OBJECTNAME] [FILENAME]")
	}
	inputFile, err := os.Open(filePath)
	defer inputFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Staging parity
	stagingConfig := config
	stagingConfig.k = 2
	stagingConfig.m = 1
	stagingConfig.rootDir = config.stagingDir

	if err := erasurePut(stagingConfig, objectName, inputFile); err != nil {
		log.Fatal(err)
	}

}
