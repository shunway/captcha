package main

import (
	"image/png"
	"net/http"

	"github.com/afocus/captcha"
)

var cap *captcha.Captcha

func main() {

	cap = captcha.New()

	if err := cap.SetFont("comic.ttf"); err != nil {
		panic(err.Error())
	}

	cap.SetSize(128, 64)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(captcha.Color{255, 255, 255})
	cap.SetBkgColor(captcha.Color{255, 0, 0}, captcha.Color{0, 0, 255}, captcha.Color{0, 153, 0})

	http.HandleFunc("/r", func(w http.ResponseWriter, r *http.Request) {
		img, str := cap.Create(6, captcha.ALL)
		png.Encode(w, img)
		println(str)
	})

	http.HandleFunc("/c", func(w http.ResponseWriter, r *http.Request) {
		str := r.URL.RawQuery
		img := cap.CreateCustom(str)
		png.Encode(w, img)
	})

	http.ListenAndServe(":8085", nil)

}
