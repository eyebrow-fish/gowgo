package gowgo

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func RenderTemplate(input, output string, attr map[string]string) error {
	inputData, err := ioutil.ReadFile(input)
	if err != nil {
		return err
	}
	varExp, err := regexp.Compile("{{\\s*var\\s*=\\s*[A-z0-9]+\\s*}}")
	if err != nil {
		return err
	}
	varMatches := varExp.FindAllString(string(inputData), -1)
	template := string(inputData)
	for _, match := range varMatches {
		key := strings.TrimFunc(
			strings.SplitAfter(match, "=")[1],
			func(r rune) bool { return r == ' ' || r == '}' },
		)
		template = strings.ReplaceAll(string(template), match, attr[key])
	}
	paths := strings.SplitAfter(output, string(os.PathSeparator))
	err = os.Mkdir(strings.Join(paths[:len(paths)-1], string(os.PathSeparator)), 0755)
	if err != nil {
		if strings.Contains(err.Error(), "exists") {
			if err := os.RemoveAll(output); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	outputFile, err := os.Create(output)
	if err != nil {
		return err
	}
	_, err = outputFile.WriteString(template)
	return err
}
