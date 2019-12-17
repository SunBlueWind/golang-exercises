// Lissajous generates green/black GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var green = color.RGBA{0, 0xff, 0, 0xff}
var red = color.RGBA{0xff, 0, 0, 0xff}
var blue = color.RGBA{0, 0, 0xff, 0xff}
var palette = []color.Color{color.Black, green, red, blue}

const (
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycle   = 5     // number of complex x oscillator revolutions
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
		for t := 0.0; t < cycle*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(i%(len(palette)-1) + 1)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
