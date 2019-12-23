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
		widthP, heightP        = width * 2, height * 2
	)
	imgP := image.NewRGBA(image.Rect(0, 0, widthP, heightP))
	for py := 0; py < heightP; py++ {
		y := float64(py)/heightP*(ymax-ymin) + ymin
		for px := 0; px < widthP; px++ {
			x := float64(px)/widthP*(xmax-xmin) + xmin
			z := complex(x, y)
			imgP.Set(px, py, mandelbrot(z))
		}
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			r1, g1, b1, _ := imgP.At(2*px, 2*py).RGBA()
			r2, g2, b2, _ := imgP.At(2*px+1, 2*py).RGBA()
			r3, g3, b3, _ := imgP.At(2*px, 2*py+1).RGBA()
			r4, g4, b4, _ := imgP.At(2*px+1, 2*py+1).RGBA()
			img.Set(px, py, color.RGBA{
				uint8((r1 + r2 + r3 + r4) / 4),
				uint8((g1 + g2 + g3 + g4) / 4),
				uint8((b1 + b2 + b3 + b4) / 4),
				0xff,
			})
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
