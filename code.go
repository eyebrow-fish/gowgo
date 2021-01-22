package gowgo

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type code string

func ReadCode(filename string) (string, error) {
	codeBytes, err := ioutil.ReadFile(fmt.Sprintf("cmd/%s/main.go", filename))
	if err != nil {
		return "", err
	}
	c := code(codeBytes)
	packages := c.highlightKeyword("package")
	imps := packages.highlightKeyword("import")
	funcs := imps.highlightKeyword("func")
	vars := funcs.highlightKeyword("var")
	consts := vars.highlightKeyword("const")
	ifs := consts.highlightKeyword("if")
	elses := ifs.highlightKeyword("else")
	prims, err := elses.highlightRegexAs("([0-9]+)|true|false", "prim")
	if err != nil {
		return "", err
	}
	strs, err := prims.highlightRegexAs("\".*\"", "string")
	if err != nil {
		return "", err
	}
	comments, err := strs.highlightRegexAs("\\/\\/.*\n", "comment")
	if err != nil {
		return "", err
	}
	trimmed := comments.removeLastLine()
	return string(trimmed), err
}

func (c code) highlightKeyword(word string) code {
	return code(strings.ReplaceAll(string(c), word, fmt.Sprintf("<span class='keyword'>%s</span>", word)))
}

func (c code) highlightRegexAs(regex, class string) (code, error) {
	pattern, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	full := string(c)
	for _, s := range pattern.FindAllString(full, -1) {
		full = strings.ReplaceAll(full, s, fmt.Sprintf("<span class='%s'>%s</span>", class, s))
	}
	return code(full), nil
}

func (c code) removeLastLine() code {
	return code(strings.TrimRight(string(c), "\n "))
}
