package main

import (
	"encoding/json"
	"log"
	"os"
)

func load(fileName string) *config {
	conf := &config{}
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(conf); err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(conf.BinPath); os.IsNotExist(err) {
		if err := os.Mkdir(conf.BinPath, 0770); err != nil {
			log.Fatal(err)
		}
	}
	return conf
}
