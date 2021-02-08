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
		packages := highlightKeywords(hi{text: types}, "package")
		imports := highlightKeywords(packages, "import")
		funcs := highlightKeywords(imports, "func")
		vars := highlightKeywords(funcs, "var")
		consts := highlightKeywords(vars, "const")
		ifs := highlightKeywords(consts, "if")
		elses := highlightKeywords(ifs, "else")
		switches := highlightKeywords(elses, "switch")
		cases := highlightKeywords(switches, "case")
		defaults := highlightKeywords(cases, "default")
		fallthroughs := highlightKeywords(defaults, "fallthrough")
		typeWords := highlightKeywords(fallthroughs, "type")
		structs := highlightKeywords(typeWords, "struct")
		interfaces := highlightKeywords(structs, "interface")
		selects := highlightKeywords(interfaces, "selects")
		stringTypes := highlightKeywords(selects, "string")
		boolTypes := highlightKeywords(stringTypes, "bool")
		defers := highlightKeywords(boolTypes, "defer")
		returns := highlightKeywords(defers, "return")
		replaced += returns.text
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
				prefix = "</span>"
				inTrue = false
				inFalse = false
				inNum = false
			}
		}
		types += prefix + string(c) + suffix
	}
	return types
}

type hi struct {
	text    string
	comment bool
}

func highlightKeywords(src hi, word string) hi {
	if src.comment {
		return src
	}
	trailing := []string{":", "("}
	var output []string
	var inComment bool
	for i, v := range strings.Split(src.text, " ") {
		trimmedL := strings.TrimSpace(v)
		if trimmedL == "//" {
			output = append(output, `<span class="comment">`+src.text[i:len(src.text)-1]+"</span>\n")
			inComment = true
		} else if trimmedL == word {
			output = append(output, `<span class="keyword">`+v+"</span>")
		} else {
			var foundSuffixed bool
			for _, t := range trailing {
				if len(trimmedL) > len(word) && trimmedL[:len(word)+1] == word+t {
					leftMargin := strings.Repeat("\t", strings.Count(v, "\t"))
					rightMargin := v[len(word)+len(leftMargin):]
					output = append(output, leftMargin+`<span class="keyword">`+word+"</span>"+rightMargin)
					foundSuffixed = true
				}
			}
			if strings.Contains(trimmedL, `class="comment"`) {
				inComment = true
			}
			if !foundSuffixed {
				output = append(output, v)
			}
		}
	}
	return hi{strings.Join(output, " "), inComment}
}
