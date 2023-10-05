package main

import (
	"github.com/parnurzeal/gorequest"
	"log"
)

func main() {
	// url := "http://localhost:1000/movie/search?movie_name=Blue Beetle"
	url := "http://localhost:1000/movie/crawl"
	// url := "http://localhost:1000/movie/list"
	request := gorequest.New()

	resp, body, errs := request.Get(url).End()
	if errs != nil {
		log.Fatal(errs)
	}
	if resp.StatusCode == 200 {
		log.Println(body)
	} else {
		log.Println(body)
	}
}