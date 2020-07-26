package macro

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	sb.WriteString("fmt.Sprint(")
	for i := range t {
		if opend == 0 && t[i] == '{' && (i == 0 || t[i-1] != '\\') {
			sb.WriteString(string(t[last:i]))
			sb.WriteString(", ")
			opend = i + 1
		} else if opend != 0 && t[i] == '}' {
			sb.WriteString(string(t[opend:i]))
			opend = 0
			last = i + 1
		}
	}
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
		if opend == 0 && runes[i] == '`' && runes[i-1] == '\\' {
			if v, ok := interpreters[runes[i+1]]; ok {
				function = v
				fmt.Fprint(tf, string(runes[last:i]))
				opend = i + 1
			}
		} else if opend != 0 && runes[i-1] != '\\' && runes[i] == '`' {
			fmt.Fprint(tf, function(runes[opend:i]))
			opend = 0
			last = i + 1
		}
	}
	fmt.Fprint(tf, runes[last:len(runes)])

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
