package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func load() (*config, string) {
	conf := &config{}
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, ""
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
	return conf, gogc
}
