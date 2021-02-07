package gowgo

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type Html struct {
	Id        string
	Href      string
	InnerHtml string
}

func RenderTemplate(input, output string, attr map[string]string, htmlTemps map[string]*Html) error {
	inputData, err := ioutil.ReadFile(input)
	if err != nil {
		return err
	}
	varMatches, err := findMatches("var", string(inputData))
	if err != nil {
		return err
	}
	template := string(inputData)
	vars := injectVars(template, attr, varMatches)
	pathMatches, err := findMatches("path", string(inputData))
	if err != nil {
		return err
	}
	paths, err := injectPaths(vars, pathMatches)
	if err != nil {
		return err
	}
	htmlMatches, err := findMatches("html", string(inputData))
	if err != nil {
		return err
	}
	htmls := injectHtml(paths, htmlTemps, htmlMatches)
	outputFile, err := os.Create(output)
	if err != nil {
		return err
	}
	_, err = outputFile.WriteString(htmls)
	return err
}

func injectVars(template string, attr map[string]string, varMatches []string) string {
	for _, match := range varMatches {
		key := strings.TrimFunc(
			strings.SplitAfter(match, "=")[1],
			func(r rune) bool { return r == ' ' || r == '}' },
		)
		template = strings.ReplaceAll(template, match, attr[key])
	}
	return template
}

func injectPaths(template string, pathMatches []string) (string, error) {
	for _, match := range pathMatches {
		fileName := strings.TrimFunc(
			strings.SplitAfter(match, "=")[1],
			func(r rune) bool { return r == ' ' || r == '}' },
		)
		file, err := os.Stat("bin/"+fileName)
		if err != nil {
			return "", err
		}
		template = strings.ReplaceAll(template, match, file.Name())
	}
	return template, nil
}

func injectHtml(template string, htmlTemps map[string]*Html, htmlMatches []string) string {
	for _, match := range htmlMatches {
		key := strings.TrimFunc(
			strings.SplitAfter(match, "=")[1],
			func(r rune) bool { return r == ' ' || r == '}' },
		)
		htmlTemp := htmlTemps[key]
		var htmlText string
		htmlText = ""
		if htmlTemp != nil {
			htmlText = fmt.Sprintf(`<a id="%s" href="%s">%s</a>`, htmlTemp.Id, htmlTemp.Href, htmlTemp.InnerHtml)
		}
		template = strings.ReplaceAll(template, match, htmlText)
	}
	return template
}

func findMatches(typeName, inputData string) ([]string, error) {
	varExp, err := regexp.Compile(fmt.Sprintf("{{\\s*%s\\s*=\\s*[A-z0-9\\/\\.\\(\\)\\,\\s]+\\s*}}", typeName))
	if err != nil {
		return nil, err
	}
	varMatches := varExp.FindAllString(inputData, -1)
	return varMatches, nil
}
