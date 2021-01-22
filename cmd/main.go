package main

import (
	"github.com/eyebrow-fish/gowgo"
	"io/ioutil"
	"log"
)

func main() {
	helloWorldCode, err := gowgo.ReadCode("hello-world")
	if err != nil {
		log.Fatal(err)
	}
	helloWorldOverview, err := ioutil.ReadFile("cmd/hello-world/overview.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = gowgo.RenderTemplate(
		"tutorial.html",
		"bin/hello-world.html",
		map[string]string{
			"lesson":   "Hello, World!",
			"overview": string(helloWorldOverview),
			"code":     helloWorldCode,
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
	ifOverview, err := ioutil.ReadFile("cmd/if/overview.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = gowgo.RenderTemplate(
		"tutorial.html",
		"bin/if.html",
		map[string]string{
			"lesson":   "If Statements",
			"overview": string(ifOverview),
			"code":     ifCode,
		},
		map[string]*gowgo.Html{
			"prev": {"prev", "hello-world.html", "Hello, World!"},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
