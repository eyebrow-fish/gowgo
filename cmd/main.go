package main

import (
	"github.com/eyebrow-fish/gowgo"
	"log"
)

func main() {
	helloWorldCode, err := gowgo.ReadCode("hello-world")
	if err != nil {
		log.Fatal(err)
	}
	err = gowgo.RenderTemplate(
		"tutorial.html",
		"bin/hello-world.html",
		map[string]string{
			"lesson": "Hello, World!",
			"code":   helloWorldCode,
		},
		map[string]*gowgo.Html{
			"next": {"next", "if.html", "If Statements"},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	ifCode, err := gowgo.ReadCode("if")
	if err != nil {
		log.Fatal(err)
	}
	err = gowgo.RenderTemplate(
		"tutorial.html",
		"bin/if.html",
		map[string]string{
			"lesson": "If Statements",
			"code": ifCode,
		},
		map[string]*gowgo.Html{
			"prev": {"prev", "hello-world.html", "Hello, World!"},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
