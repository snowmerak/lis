package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func buildWith(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	conf := &BuildConfig{}
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(conf); err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(conf.BinPath); os.IsNotExist(err) {
		if err := os.Mkdir(conf.BinPath, 0770); err != nil {
			log.Fatal(err)
		}
	}

	if err := buildAll(conf); err != nil {
		log.Fatal(err)
	}
}

func buildAll(conf *BuildConfig) error {
	errChan := make(chan error)
	max := 0
	gogc := fmt.Sprintf("GOGC=%d", conf.GOGC)
	if conf.GOGC == 0 {
		gogc = "GOGC=off"
	}
	mod := "GO111MODULE=on"
	if !conf.Module {
		mod = "GO111MODULE=off"
	}
	for k, archs := range conf.Target {
		for _, arch := range archs {
			max++
			go func(gogc, arch, k string) {
				fmt.Printf("start to compile for %s %s\n", k, arch)
				cmd := exec.Command("env", mod, gogc, fmt.Sprintf("GOOS=%s", k), fmt.Sprintf("GOARCH=%s", arch), "go", "build")
				if conf.ToPlugin {
					cmd.Args = append(cmd.Args, "-buildmode=plugin")
				}
				cmd.Args = append(cmd.Args, "-o", filepath.Join(conf.BinPath, fmt.Sprintf("%s-%s-%s", conf.Name, k, arch)))
				if o, err := cmd.Output(); err != nil {
					errChan <- errors.New(fmt.Sprint(o, "\n", err.Error()))
				} else {
					errChan <- nil
					fmt.Printf("compiled for %s %s\n", k, arch)
				}
			}(gogc, arch, k)
		}
	}

	num := 0
	for err := range errChan {
		num++
		if err != nil {
			close(errChan)
			return err
		}
		if num == max {
			close(errChan)
			return nil
		}
	}
	return nil
}
