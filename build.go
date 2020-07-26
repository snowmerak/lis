package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"sync"
)

func buildAll(conf *config, gogc string) {
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
