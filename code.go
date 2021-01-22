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
	code, err = highlightStrings(code)
	if err != nil {
		return "", err
	}
	code = highlightKeyword(code, "package")
	code = highlightKeyword(code, "import")
	code = highlightKeyword(code, "func")
	code = highlightKeyword(code, "var")
	code = highlightKeyword(code, "const")
	code = strings.TrimRight(code, "\n ")
	return string(code), err
}

func highlightKeyword(code, word string) string {
	return strings.ReplaceAll(code, word, fmt.Sprintf(`<span class="keyword">%s</span>`, word))
}

func highlightStrings(code string) (string, error) {
	str, err := regexp.Compile("\".*\"")
	if err != nil {
		return "", err
	}
	for _, s := range str.FindAllString(code, -1) {
		code = strings.Replace(code, s, fmt.Sprintf(`<span class="string">%s</span>`, s), 1)
	}
	return code, nil
}
