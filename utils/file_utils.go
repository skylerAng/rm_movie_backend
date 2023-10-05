package utils

import (
    "bytes"
    "image/jpeg"
    "image"
    "log"
    "os"
	"path/filepath"
	"errors"
    "encoding/base64"
    "io/ioutil"
)

func CreateFolder(path string) {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
}

func SaveImage(imgByte []byte, image_path string) {
    img, _, err := image.Decode(bytes.NewReader(imgByte))
    if err != nil {
        log.Fatalln(err)
    }

    out, _ := os.Create(image_path)
    defer out.Close()

    var opts jpeg.Options
    opts.Quality = 1

    err = jpeg.Encode(out, img, &opts)
    //jpeg.Encode(out, img, nil)
    if err != nil {
        log.Println(err)
    }
}

func FileExits(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func ImagetoBase64(filepath string) (string, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println(err)
        return "", err
	}

	base64Encoding := ""
	base64Encoding += ToBase64(bytes)
	return base64Encoding, nil
}