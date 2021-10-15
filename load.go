package main

import (
	"encoding/json"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func load(fileName string) *config {
	conf := &config{}
	b := true
	f, err := os.Open(fileName + ".yaml")
	if err != nil {
		f, err = os.Open(fileName + ".json")
		if err != nil {
			log.Fatal(err)
		}
		b = false
	}
	switch b {
	case true:
		decoder := yaml.NewDecoder(f)
		if err := decoder.Decode(conf); err != nil {
			log.Fatal(err)
		}
	case false:
		decoder := json.NewDecoder(f)
		if err := decoder.Decode(conf); err != nil {
			log.Fatal(err)
		}
	}
	if _, err := os.Stat(conf.BinPath); os.IsNotExist(err) {
		if err := os.Mkdir(conf.BinPath, 0770); err != nil {
			log.Fatal(err)
		}
	}
	return conf
}
