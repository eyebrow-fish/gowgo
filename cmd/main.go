package main

import (
	"github.com/eyebrow-fish/gowgo"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type tutorial struct {
	name string
	dir  string
}

func main() {
	gen([]tutorial{
		{name: "Hello, World!", dir: "hello-world"},
		{name: "If Statements", dir: "if"},
	})
}

func gen(tuts []tutorial) {
	for i, tut := range tuts {
		code, err := gowgo.ReadCode(tut.dir)
		if err != nil {
			log.Fatal(err)
		}
		var lines string
		for i := range strings.Split(code, "\n") {
			lines += strconv.Itoa(i + 1) + "\n"
		}
		overview, err := ioutil.ReadFile("cmd/" + tut.dir + "/overview.txt")
		if err != nil {
			log.Fatal(err)
		}
		var prev *gowgo.Html = nil
		if i > 0 {
			prev = &gowgo.Html{Id: "prev", Href: tuts[i-1].dir + ".html", InnerHtml: tuts[i-1].name}
		}
		var next *gowgo.Html = nil
		if i < len(tuts)-1 {
			next = &gowgo.Html{Id: "next", Href: tuts[i+1].dir + ".html", InnerHtml: tuts[i+1].name}
		}
		err = gowgo.RenderTemplate(
			"tutorial.html",
			"bin/"+tut.dir+".html",
			map[string]string{
				"lesson":   tut.name,
				"lines":    lines,
				"overview": string(overview),
				"code":     code,
			},

			map[string]*gowgo.Html{"prev": prev, "next": next},
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}
