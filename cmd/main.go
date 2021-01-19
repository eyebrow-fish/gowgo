package main

import (
	"github.com/eyebrow-fish/gowgo"
	"log"
)

func main() {
	err := gowgo.RenderTemplate("tutorial.html", "bin/hello-world.html", map[string]string{
		"lesson": "Hello, World!",
	})
	if err != nil {
		log.Fatal(err)
	}
}
