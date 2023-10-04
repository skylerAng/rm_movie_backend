package main

import (
	"gorequest"

	"fmt"
)

func main() {
	testreq := gorequest.New()
	resp, body, errs := testreq.Get("https://image.tmdb.org/t/p/original/1syW9SNna38rSl9fnXwc9fP7POW.jpg").End()

	structtest := struct {
		name string
	}
}
