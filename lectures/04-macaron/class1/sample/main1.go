package main

import "github.com/Unknwon/macaron"

func main() {
	m := macaron.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
