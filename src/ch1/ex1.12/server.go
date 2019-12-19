// Server is a web server that displays a Lissajous GIF image
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

var cycle = 5 // number of complex x oscillator revolutions
var palette = []color.Color{color.Black, color.RGBA{0, 0xff, 0, 0xff}}

const (
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles, ok := r.Form["cycles"]
	if !ok || len(cycles) == 0 {
		fmt.Fprintf(w, "Please provide a query param 'cycles'.")
		return
	}
	if cycle, err = strconv.Atoi(cycles[0]); err != nil {
		fmt.Fprint(w, err)
		return
	}
	lissajous(w)
}

func lissajous(out io.Writer) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas convers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms unites
	)
	freq := rand.Float64() * 3.0 // relative frequency of y osillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycle)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
