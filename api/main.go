package main

import (
	"demo/assets"
	"fmt"
	"net/http"

	"github.com/h2non/bimg"
	"github.com/morelj/lambada"
	"github.com/yudai/pp"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	http.HandleFunc("/never-gonna", func(w http.ResponseWriter, r *http.Request) {
		watermark := bimg.Watermark{
			Text:       "Chuck Norris (c) 2315",
			Opacity:    0.25,
			Width:      200,
			DPI:        100,
			Margin:     150,
			Font:       "sans bold 12",
			Background: bimg.Color{255, 255, 255},
		}

		newImage, err := bimg.NewImage(assets.Rick).Watermark(watermark)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
			return
		}
		pp.Println(len(newImage))

		w.Header().Set("Content-Type", "image/jpeg")
		w.WriteHeader(http.StatusOK)
		w.Write(newImage)
	})

	lambada.ServeWithOptions(http.StripPrefix("/go", http.DefaultServeMux), lambada.WithDefaultBinary(true))
}
