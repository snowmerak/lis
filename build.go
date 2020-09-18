package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

func buildAll(conf *config) error {
	errChan := make(chan error)
	max := 0
	gogc := fmt.Sprintf("GOGC=%d", conf.GOGC)
	if conf.GOGC == 0 {
		gogc = "GOGC=off"
	}
	for k, archs := range conf.Target {
		for _, arch := range archs {
			max++
			go func(gogc, arch, k string) {
				fmt.Printf("start to compile for %s %s\n", k, arch)
				cmd := exec.Command("env", gogc, fmt.Sprintf("GOOS=%s", k), fmt.Sprintf("GOARCH=%s", arch), "go", "build")
				if conf.IsPlugin {
					cmd.Args = append(cmd.Args, "-buildmode=plugin")
				}
				cmd.Args = append(cmd.Args, "-o", filepath.Join(conf.BinPath, fmt.Sprintf("%s-%s-%s", conf.BinName, k, arch)))
				if _, err := cmd.Output(); err != nil {
					errChan <- err
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
