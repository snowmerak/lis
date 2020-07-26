package macro

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func readalldir() []string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	rs := readdir(path)

	return rs
}

func readdir(path string) []string {
	fl, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	rs := []string{}

	for _, f := range fl {
		if f.IsDir() {
			rs = append(rs, readdir(filepath.Join(path, f.Name()))...)
			continue
		}
		if strings.HasSuffix(filepath.Base(f.Name()), "-lib.go") {
			rs = append(rs, filepath.Join(path, f.Name()))
		}
	}

	return rs
}
