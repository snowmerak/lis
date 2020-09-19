package macro

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var interpreters map[rune]func([]rune) string

func init() {
	interpreters = map[rune]func([]rune) string{}
	interpreters['"'] = evalString
}

func evalString(t []rune) string {
	sb := strings.Builder{}
	opend := 0
	last := 0
	t = t[1 : len(t)-1]
	sb.WriteString("fmt.Sprint(")
	for i := range t {
		if opend == 0 && t[i] == '{' && (i == 0 || t[i-1] != '\\') {
			if last != 0 {
				sb.WriteString(fmt.Sprint("\"", string(t[last:i]), "\""))
				sb.WriteString(", ")
			}
			opend = i + 1
		} else if opend != 0 && t[i] == '}' {
			sb.WriteString(string(t[opend:i]))
			sb.WriteString(", ")
			opend = 0
			last = i + 1
		}
	}
	sb.WriteString(fmt.Sprint("\"", string(t[last:len(t)]), "\""))
	sb.WriteString(")")

	return sb.String()
}

func translate(filename string) bool {
	defer runtime.GC()
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return false
	}

	tname := fmt.Sprintf("%s.go", strings.TrimSuffix(filename, "-lis.go"))

	if _, err := os.Stat(tname); !os.IsNotExist(err) {
		if err := os.Remove(tname); err != nil {
			log.Fatal(err)
			return false
		}
	}

	tf, err := os.Create(tname)
	if err != nil {
		log.Fatal(err)
		return false
	}

	runes := []rune(string(bytes))
	opend := 0
	last := 0
	var function func([]rune) string
	for i := range runes {
		if opend == 0 && runes[i] == '`' && runes[i-1] != '\\' {
			if v, ok := interpreters[runes[i+1]]; ok {
				function = v
				tf.WriteString(string(runes[last:i]))
				opend = i + 1
			}
		} else if opend != 0 && runes[i] == '`' && runes[i-1] != '\\' {
			tf.WriteString(function(runes[opend:i]))
			opend = 0
			last = i + 1
		}
	}
	tf.WriteString(string(runes[last:len(runes)]))

	goimports := exec.Command(filepath.Join(build.Default.GOPATH, "bin", "goimports"), tname)
	if importsResult, err := goimports.Output(); err != nil {
		fmt.Print(err)
	} else {
		ioutil.WriteFile(tname, importsResult, os.ModePerm)
	}

	if err := os.Rename(filename, strings.TrimSuffix(filename, ".go")); err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

func rollbackTranslated(t []string) error {
	for _, f := range t {
		if err := os.Remove(fmt.Sprintf("%s.go", strings.TrimSuffix(f, "-lis.go"))); err != nil {
			return err
		}
	}

	return nil
}
