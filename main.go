package main

import (
	"flag"
	"fmt"
	"lisb/macro"
	"os"
	"os/exec"
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

	conf := load()

	if err := buildAll(conf); err != nil {
		fmt.Println(err)
	}

	macro.End(files)

	if conf.WillRun {
		fmt.Println("--- AutoRun ---")
		cmd := exec.Command(fmt.Sprintf("./%s/%s-%s-%s", conf.BinPath, conf.BinName, runtime.GOOS, runtime.GOARCH))
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
