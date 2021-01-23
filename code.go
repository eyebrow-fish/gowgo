package gowgo

import (
	"fmt"
	"io/ioutil"
	"strconv"

	//"regexp"
	"strings"
)

func ReadCode(filename string) (string, error) {
	codeBytes, err := ioutil.ReadFile(fmt.Sprintf("cmd/%s/main.go", filename))
	if err != nil {
		return "", err
	}
	var replaced string
	for _, line := range strings.Split(string(codeBytes), "\n") {
		line = line + "\n"
		if strings.HasPrefix(strings.TrimSpace(line), "//") {
			replaced += fmt.Sprintf(`<span class="comment">%s</span>`, line)
			continue
		}
		var strs string
		var inStr bool
		var inNum bool
		var inTrue bool
		//var inFalse bool
		for i, c := range line {
			var prefix, suffix string
			if c == '"' {
				if !inStr {
					prefix = `<span class="string">`
					inStr = true
				} else {
					suffix = "</span>"
					inStr = false
				}
			} else if len(line) >= i+5 && line[i:i+5] == "true" {
				suffix = `<span class="prim">`
				inTrue = false
			} else if inTrue && strings.ContainsRune("true", c) {
				suffix = "</span>"
				inTrue = false
			} else if _, err := strconv.Atoi(string(c)); err == nil || c == ' ' {
				if !inNum && err == nil {
					prefix = `<span class="prim">`
					inNum = true
				} else if inNum && err != nil {
					suffix = "</span>"
					inNum = false
				}
			}
			strs += prefix + string(c) + suffix
		}
		packages := strings.ReplaceAll(strs, "package", `<span class="keyword">package</span>`)
		imports := strings.ReplaceAll(packages, "import", `<span class="keyword">import</span>`)
		funcs := strings.ReplaceAll(imports, "func", `<span class="keyword">func</span>`)
		vars := strings.ReplaceAll(funcs, "var", `<span class="keyword">var</span>`)
		consts := strings.ReplaceAll(vars, "const", `<span class="keyword">const</span>`)
		ifs := strings.ReplaceAll(consts, "if", `<span class="keyword">if</span>`)
		elses := strings.ReplaceAll(ifs, "else", `<span class="keyword">else</span>`)
		replaced += elses
	}
	return replaced, err
}
