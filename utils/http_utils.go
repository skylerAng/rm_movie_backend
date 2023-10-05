package utils

import (
	"github.com/parnurzeal/gorequest"
	"log"
)

func GetImage(image_path string) ([]byte, string, error) {
	img_url := "https://image.tmdb.org/t/p/original" + image_path
	request := gorequest.New()	

	resp, body, errs := request.Get(img_url).End()
	if errs != nil {
		return nil, img_url, errs[0]
	}

	if resp.StatusCode == 200 {
		log.Println("Image Received")
	}

	return []byte(body),  img_url, nil
}