package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var Palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":3454", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, request *http.Request) {

	cycles, err := strconv.ParseFloat(request.URL.Query().Get("cycles"), 64)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}
	res, err := strconv.ParseFloat(request.URL.Query().Get("res"), 64)
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}
	size, err := strconv.Atoi(request.URL.Query().Get("size"))
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}
	nframes, err := strconv.Atoi(request.URL.Query().Get("nframes"))
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}
	delay, err := strconv.Atoi(request.URL.Query().Get("delay"))
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
	}

	lissajous(w, cycles, res, size, nframes, delay)
}

func lissajous(out io.Writer, cycles float64, res float64, size int, nframes int, delay int) {

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, Palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors

}
