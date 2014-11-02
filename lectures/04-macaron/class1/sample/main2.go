package main

import (
	"log"
	"net/http"

	"github.com/Unknwon/macaron"
)

func main() {
	m := macaron.Classic()
	m.Get("/", myHandler)

	log.Println("Server is running...")
	log.Println(http.ListenAndServe("0.0.0.0:4000", m))
}

func myHandler(ctx *macaron.Context) string {
	return "the request path is: " + ctx.Req.RequestURI
}
