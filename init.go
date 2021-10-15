package main

import (
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v2"
)

func initoption(fileName string) {
	conf := &config{
		BinPath: "bin",
		Name:    "temp",
		Target: map[string][]string{
			"windows": {"amd64", "386"},
			"darwin":  {"amd64", "arm64"},
			"linux":   {"amd64", "386", "arm64"},
		},
		GOGC:     150,
		ToPlugin: false,
		Module:   false,
	}
	qs := []*survey.Question{
		{
			Name: "Name",
			Prompt: &survey.Input{
				Message: "input project name: ",
				Default: "demo",
			},
		},
		{
			Name: "BinPath",
			Prompt: &survey.Input{
				Message: "input path of compiled file: ",
				Default: "bin",
			},
		},
		{
			Name: "GOGC",
			Prompt: &survey.Select{
				Message: "choose gogc: ",
				Default: "150",
				Options: []string{"off", "50", "100", "150", "200", "250"},
			},
		},
		{
			Name: "AutoRun",
			Prompt: &survey.Confirm{
				Message: "do you want to run app after compile? ",
				Default: false,
			},
		},
		{
			Name: "ToPlugin",
			Prompt: &survey.Confirm{
				Message: "do you want to compile to plugin? ",
				Default: false,
			},
		},
		{
			Name: "Module",
			Prompt: &survey.Confirm{
				Message: "will you use go111module? ",
				Default: true,
			},
		},
	}
	if err := survey.Ask(qs, conf); err != nil {
		log.Fatal(err)
	}
	conf.GOGC *= 50
	fileName += ".yaml"
	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		if err := os.Remove(fileName); err != nil {
			log.Fatal(err)
		}
	}
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	conf.Etc = make([]string, 0)
	encoder := yaml.NewEncoder(f)
	if err := encoder.Encode(conf); err != nil {
		log.Fatal(err)
	}
}
