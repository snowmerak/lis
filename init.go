package main

import (
	"log"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v2"
)

func initTemplate(fileName string) {
	const (
		BUILD = "build"
		TEST  = "test"
		BENCH = "bench"
	)

	if !strings.HasSuffix(fileName, ".yaml") && !strings.HasSuffix(fileName, ".yml") {
		fileName += ".yaml"
	}

	target := ""
	if err := survey.AskOne(&survey.Select{
		Message: "choose init type: ",
		Options: []string{BUILD, TEST, BENCH},
	}, &target, survey.WithValidator(survey.Required)); err != nil {
		log.Fatal(err)
	}

	switch target {
	case BUILD:
		initBuild(fileName)
	case TEST:
		initTest(fileName)
	case BENCH:
		initBench(fileName)
	default:
		log.Fatal("unknown type")
	}
}

func initBuild(fileName string) {
	conf := &BuildConfig{
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

func initTest(fileName string) {
	conf := &TestConfig{
		Targets: []struct {
			Package string   `json:"package" yaml:"package"`
			Test    []string `json:"test" yaml:"test"`
		}{},
		TestFlags: []string{},
	}

	for {
		isEnd := false
		if err := survey.AskOne(&survey.Confirm{
			Message: "end to input test target? ",
			Default: false,
		}, &isEnd, survey.WithValidator(survey.Required)); err != nil {
			log.Fatal(err)
		}
		if isEnd {
			break
		}
		packageName := ""
		if err := survey.AskOne(&survey.Input{
			Message: "input package name: ",
		}, &packageName, survey.WithValidator(survey.Required)); err != nil {
			log.Fatal(err)
		}
		testNames := []string(nil)
		testName := ""
		for {
			isEnd := false
			if err := survey.AskOne(&survey.Confirm{
				Message: "end to input test name? ",
				Default: false,
			}, &isEnd, survey.WithValidator(survey.Required)); err != nil {
				log.Fatal(err)
			}
			if isEnd {
				break
			}
			if err := survey.AskOne(&survey.Input{
				Message: "input test name: ",
			}, &testName, survey.WithValidator(survey.Required)); err != nil {
				log.Fatal(err)
			}
			testNames = append(testNames, testName)
		}
		conf.Targets = append(conf.Targets, struct {
			Package string   `json:"package" yaml:"package"`
			Test    []string `json:"test" yaml:"test"`
		}{
			Package: packageName,
			Test:    testNames,
		})
	}

	for {
		isEnd := false
		if err := survey.AskOne(&survey.Confirm{
			Message: "end to input test flag? ",
			Default: false,
		}, &isEnd, survey.WithValidator(survey.Required)); err != nil {
			log.Fatal(err)
		}
		if isEnd {
			break
		}
		flagValue := ""
		if err := survey.AskOne(&survey.Input{
			Message: "input test flag: ",
		}, &flagValue, survey.WithValidator(survey.Required)); err != nil {
			log.Fatal(err)
		}
		conf.TestFlags = append(conf.TestFlags, flagValue)
	}

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	encoder := yaml.NewEncoder(f)
	if err := encoder.Encode(conf); err != nil {
		log.Fatal(err)
	}
}

func initBench(fileName string) {
	conf := &BenchConfig{
		Targets: []struct {
			Package string   `json:"package" yaml:"package"`
			Bench   []string `json:"bench" yaml:"bench"`
		}{},
		BenchFlags: []string{},
	}

	for {
		isEnd := false
		if err := survey.AskOne(&survey.Confirm{
			Message: "end to input bench target? ",
			Default: false,
		}, &isEnd, survey.WithValidator(survey.Required)); err != nil {
			log.Fatal(err)
		}
		if isEnd {
			break
		}
		packageName := ""
		if err := survey.AskOne(&survey.Input{
			Message: "input package name: ",
		}, &packageName, survey.WithValidator(survey.Required)); err != nil {
			log.Fatal(err)
		}
		benchNames := []string(nil)
		benchName := ""
		for {
			isEnd := false
			if err := survey.AskOne(&survey.Confirm{
				Message: "end to input bench name? ",
				Default: false,
			}, &isEnd, survey.WithValidator(survey.Required)); err != nil {
				log.Fatal(err)
			}
			if isEnd {
				break
			}
			if err := survey.AskOne(&survey.Input{
				Message: "input bench name: ",
			}, &benchName, survey.WithValidator(survey.Required)); err != nil {
				log.Fatal(err)
			}
			benchNames = append(benchNames, benchName)
		}
		conf.Targets = append(conf.Targets, struct {
			Package string   `json:"package" yaml:"package"`
			Bench   []string `json:"bench" yaml:"bench"`
		}{
			Package: packageName,
			Bench:   benchNames,
		})
	}

	for {
		isEnd := false
		if err := survey.AskOne(&survey.Confirm{
			Message: "end to input bench flag? ",
			Default: false,
		}, &isEnd, survey.WithValidator(survey.Required)); err != nil {
			log.Fatal(err)
		}
		if isEnd {
			break
		}
		flagValue := ""
		if err := survey.AskOne(&survey.Input{
			Message: "input bench flag: ",
		}, &flagValue, survey.WithValidator(survey.Required)); err != nil {
			log.Fatal(err)
		}
		conf.BenchFlags = append(conf.BenchFlags, flagValue)
	}

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	encoder := yaml.NewEncoder(f)
	if err := encoder.Encode(conf); err != nil {
		log.Fatal(err)
	}
}
