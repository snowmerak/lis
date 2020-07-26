package macro

import (
	"log"
	"os"
	"strings"
)

//Run ...
func Run() []string {
	files := readalldir()
	for i, f := range files {
		if !translate(f) {
			if err := rollbackTranslated(files[:i+1]); err != nil {
				log.Fatal(err)
			}
			break
		}
	}
	return files
}

//End ...
func End(t []string) {
	for _, f := range t {
		if err := os.Rename(strings.TrimSuffix(f, ".go"), f); err != nil {
			log.Fatal(err)
		}
	}
	rollbackTranslated(t)
}
