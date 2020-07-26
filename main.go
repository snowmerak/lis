package main

import (
	"flag"
	"lisb/macro"
	"runtime"
)

const fileName = "build_option.json"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	init := flag.String("init", "", "init [project name]: to create build_option.json with project name")
	flag.Parse()

	if *init != "" {
		initoption(*init)
		return
	}

	files := macro.Run()

	buildAll(load())

	macro.End(files)
}
