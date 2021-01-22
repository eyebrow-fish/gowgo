package gowgo

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func ReadCode(filename string) (string, error) {
	codeBytes, err := ioutil.ReadFile(fmt.Sprintf("cmd/%s/main.go", filename))
	if err != nil {
		return "", err
	}
	code := string(codeBytes)
	code, err = highlightRegexAs(code, "\".*\"", "string")
	if err != nil {
		return "", err
	}
	code, err = highlightRegexAs(code, "[0-9]+", "prim")
	if err != nil {
		return "", err
	}
	code = highlightKeyword(code, "package")
	code = highlightKeyword(code, "import")
	code = highlightKeyword(code, "func")
	code = highlightKeyword(code, "var")
	code = highlightKeyword(code, "const")
	code = highlightKeyword(code, "if")
	code = highlightKeyword(code, "else")
	code = strings.TrimRight(code, "\n ")
	return code, err
}

func highlightKeyword(code, word string) string {
	return strings.ReplaceAll(code, word, fmt.Sprintf(`<span class="keyword">%s</span>`, word))
}

func highlightRegexAs(code, regex, class string) (string, error) {
	pattern, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	for _, s := range pattern.FindAllString(code, -1) {
		code = strings.Replace(code, s, fmt.Sprintf(`<span class="%s">%s</span>`, class, s), 1)
	}
	return code, nil
}
