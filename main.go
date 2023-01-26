package main

import (
	"flag"
	"strings"
)

func main() {
	initFile := flag.String("init", "", "initialize build or test file  {{file_name}}.yaml")
	testFile := flag.String("test", "", "run test file {{file_name}}.yaml")
	benchFile := flag.String("bench", "", "run bench file {{file_name}}.yaml")
	makeFile := flag.String("build", "", "build project with {{file_name}}.yaml")
	flag.Parse()

	if initFile != nil && *initFile != "" {
		if !strings.HasSuffix(*initFile, ".yaml") || !strings.HasSuffix(*initFile, ".yml") {
			*initFile += ".yaml"
		}
		initTemplate(*initFile)
	}

	if testFile != nil && *testFile != "" {
		if !strings.HasSuffix(*testFile, ".yaml") || !strings.HasSuffix(*testFile, ".yml") {
			*testFile += ".yaml"
		}
		testWith(*testFile)
	}

	if benchFile != nil && *benchFile != "" {
		if !strings.HasSuffix(*benchFile, ".yaml") || !strings.HasSuffix(*benchFile, ".yml") {
			*benchFile += ".yaml"
		}
		benchWith(*benchFile)
	}

	if makeFile != nil && *makeFile != "" {
		if !strings.HasSuffix(*makeFile, ".yaml") || !strings.HasSuffix(*makeFile, ".yml") {
			*makeFile += ".yaml"
		}
		buildWith(*makeFile)
	}
}
