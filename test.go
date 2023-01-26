package main

import (
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v2"
)

func testWith(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	decoder := yaml.NewDecoder(f)
	conf := &TestConfig{}
	if err := decoder.Decode(conf); err != nil {
		log.Fatal(err)
	}

	commands := append([]string{"test"}, conf.TestFlags...)

	for _, target := range conf.Targets {
		for _, test := range target.Test {
			cmd := exec.Command("go", append(commands, target.Package, "-run", test)...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func benchWith(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	decoder := yaml.NewDecoder(f)
	conf := &BenchConfig{}
	if err := decoder.Decode(conf); err != nil {
		log.Fatal(err)
	}

	commands := append([]string{"test"}, conf.BenchFlags...)

	for _, target := range conf.Targets {
		for _, bench := range target.Bench {
			cmd := exec.Command("go", append(commands, target.Package, "-bench", bench)...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
