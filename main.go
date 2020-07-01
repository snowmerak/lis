package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
)

func main() {
	const fileName = "build_option.json"

	runtime.GOMAXPROCS(runtime.NumCPU())

	init := flag.String("init", "", "init [project name]: to create build_option.json with project name")
	flag.Parse()

	if *init != "" {
		conf := &config{
			BinPath: "bin",
			BinName: *init,
			Target: map[string][]string{
				"windows": {"amd64", "386"},
				"darwin":  {"amd64", "arm64"},
				"linux":   {"amd64", "386", "arm64"},
			},
			GOGC: 150,
		}
		if _, err := os.Stat(fileName); !os.IsNotExist(err) {
			if err := os.Remove(fileName); err != nil {
				log.Fatal(err)
			}
		}
		f, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		encoder := json.NewEncoder(f)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(conf); err != nil {
			log.Fatal(err)
		}
		return
	}

	conf := &config{}
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(conf); err != nil {
		log.Fatal(err)
	}
	gogc := fmt.Sprintf("GOGC=%d", conf.GOGC)
	if conf.GOGC == 0 {
		gogc = "GOGC=off"
	}
	if _, err := os.Stat(conf.BinPath); os.IsNotExist(err) {
		if err := os.Mkdir(conf.BinPath, 0770); err != nil {
			log.Fatal(err)
		}
	}

	wg := sync.WaitGroup{}

	for k, archs := range conf.Target {
		for _, arch := range archs {
			wg.Add(1)
			go func(gogc, arch, k string, wg *sync.WaitGroup) {
				fmt.Printf("start to compile for %s %s\n", k, arch)
				cmd := exec.Command("env", gogc, fmt.Sprintf("GOOS=%s", k),
					fmt.Sprintf("GOARCH=%s", arch), "go", "build", "-o",
					filepath.Join(conf.BinPath, fmt.Sprintf("%s-%s-%s", conf.BinName, k, arch)))
				if _, err := cmd.Output(); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("compiled for %s %s\n", k, arch)
				wg.Done()
			}(gogc, arch, k, &wg)
		}
	}

	wg.Wait()
}
