package gowgo

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func ReadCode(filename string) (string, error) {
	codeBytes, err := ioutil.ReadFile(fmt.Sprintf("cmd/%s/main.go", filename))
	if err != nil {
		return "", err
	}
	var replaced string
	for _, line := range strings.Split(string(codeBytes), "\n") {
		line = line + "\n"
		types := highlightTypes(line)
		packages := highlightKeywords(types, "package")
		imports := highlightKeywords(packages, "import")
		funcs := highlightKeywords(imports, "func")
		vars := highlightKeywords(funcs, "var")
		consts := highlightKeywords(vars, "const")
		ifs := highlightKeywords(consts, "if")
		elses := highlightKeywords(ifs, "else")
		replaced += elses
	}
	return strings.TrimSpace(replaced), err
}

func highlightTypes(src string) string {
	var types string
	var inStr bool
	var inTrue bool
	var inFalse bool
	var inNum bool
	for i, c := range src {
		var prefix, suffix string
		if c == '/' && i < len(src)-1 && src[i+1] == '/' {
			types += `<span class="comment">` + src[i:len(src)-1] + "</span>\n"
			break
		} else if c == '"' {
			if !inStr {
				prefix = `<span class="string">`
				inStr = true
			} else {
				suffix = "</span>"
				inStr = false
			}
		} else if !inStr {
			if !inTrue && len(src) >= i+4 && src[i:i+4] == "true" {
				prefix = `<span class="prim">`
				inTrue = true
			} else if !inFalse && len(src) >= i+5 && src[i:i+5] == "false" {
				prefix = `<span class="prim">`
				inFalse = true
			} else if !inNum && unicode.IsDigit(c) {
				prefix = `<span class="prim">`
				inNum = true
			} else if inTrue && !strings.ContainsRune("true", c) ||
				inFalse && !strings.ContainsRune("false", c) ||
				inNum && !unicode.IsDigit(c) {
				suffix = "</span>"
				inTrue = false
				inFalse = false
				inNum = false
			}
		}
		types += prefix + string(c) + suffix
	}
	return types
}

func highlightKeywords(src, word string) string {
	var output []string
	for i, v := range strings.Split(src, " ") {
		if v == "//" {
			output = append(output, `<span class="comment">`+src[i:len(src)-1]+"</span>\n")
			break
		} else if v == word {
			output = append(output, `<span class="keyword">`+v+"</span>")
		} else {
			output = append(output, v)
		}
	}
	return strings.Join(output, " ")
}
