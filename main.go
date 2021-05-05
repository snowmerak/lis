package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	init := flag.String("init", "", "init [build option name]: create [build option name].json")
	name := flag.String("make", "", "make [build option name]: compile with info of [build option name].json")
	flag.Parse()

	if *init != "" {
		*init = fmt.Sprintf("%s.json", *init)
		initoption(*init)
		return
	}

	if *name == "" {
		fmt.Println("build option file name is not space")
		return
	}
	*name = fmt.Sprintf("%s.json", *name)
	conf := load(*name)

	if err := buildAll(conf); err != nil {
		fmt.Println(err)
	}

	if conf.AutoRun {
		fmt.Println("--- AutoRun ---")
		cmd := exec.Command(fmt.Sprintf("./%s/%s-%s-%s", conf.BinPath, conf.Name, runtime.GOOS, runtime.GOARCH))
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}
