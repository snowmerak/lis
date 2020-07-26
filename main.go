package main

import (
	"flag"
	"fmt"
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

	if err := buildAll(load()); err != nil {
		fmt.Println(err)
	}

	macro.End(files)
}
