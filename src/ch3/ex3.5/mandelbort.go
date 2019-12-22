// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ingoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	const (
		freqRed, phaseRed     = 1, 2
		freqBlue, phaseBlue   = 3, 4
		freqGreen, phaseGreen = 5, 6
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := uint8((math.Sin(float64(n*freqRed)+phaseRed) + 1) / 2 * 256)
			g := uint8((math.Sin(float64(n*freqGreen)+phaseGreen) + 1) / 2 * 256)
			b := uint8((math.Sin(float64(n*freqBlue)+phaseBlue) + 1) / 2 * 256)
			return color.RGBA{r, g, b, 0xff}
		}
	}
	return color.Black
}
