package main

import (
	"encoding/json"
	"log"
	"os"
)

func initoption(name string) {
	conf := &config{
		BinPath: "bin",
		BinName: name,
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
}
