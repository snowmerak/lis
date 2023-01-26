package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ballast := make([]byte, 2<<10)

	init := flag.String("init", "", "init [build option name]: create [build option name].json")
	name := flag.String("make", "", "make [build option name]: compile with info of [build option name].json")
	flag.Parse()

	if *init != "" {
		initoption(*init)
		return
	}

	if *name == "" {
		fmt.Println("build option file name is not space")
		return
	}
	conf := load(*name)

	if err := buildAll(conf); err != nil {
		fmt.Println(err)
	}

	_ = ballast
}
