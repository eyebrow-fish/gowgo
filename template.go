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
	template = injectVars(template, attr, varMatches)
	pathMatches, err := findMatches("path", string(inputData))
	if err != nil {
		return err
	}
	template, err = injectPaths(template, output, pathMatches)
	if err != nil {
		return err
	}
	htmlMatches, err := findMatches("html", string(inputData))
	if err != nil {
		return err
	}
	template = injectHtml(template, htmlTemps, htmlMatches)
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

func injectPaths(template, output string, pathMatches []string) (string, error) {
	for _, match := range pathMatches {
		fileName := strings.TrimFunc(
			strings.SplitAfter(match, "=")[1],
			func(r rune) bool { return r == ' ' || r == '}' },
		)
		file, err := os.Stat(fileName)
		if err != nil {
			return "", err
		}
		parent := strings.Repeat("../", strings.Count(output, "/"))
		template = strings.ReplaceAll(template, match, parent+file.Name())
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
	varExp, err := regexp.Compile(fmt.Sprintf("{{\\s*%s\\s*=\\s*[A-z0-9\\.\\(\\)\\,\\s]+\\s*}}", typeName))
	if err != nil {
		return nil, err
	}
	varMatches := varExp.FindAllString(inputData, -1)
	return varMatches, nil
}
