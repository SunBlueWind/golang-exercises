// Newton's emits a PNG of a fractal to the complex function z^4 - 1 = 0
package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

var colors = []color.RGBA{
	{0x6a, 0x2c, 0x70, 0xff},
	{0x08, 0xd9, 0xd6, 0xff},
	{0xf0, 0x8a, 0x5d, 0xff},
	{0xf9, 0xed, 0x69, 0xff},
}

var colorMap = make(map[complex128]color.RGBA)

func main() {
	const (
		xmin, xmax, ymin, ymax = -2, 2, -2, 2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for px := 0; px < width; px++ {
		x := float64(px)/width*(xmax-xmin) + xmin
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			z := complex(x, y)
			img.Set(px, py, newtons(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// f(z) = z^4 - 1
// f'(z) = 4*z^3
// x_{n+1} = x_n - f(z) / f'(z)
func newtons(z complex128) color.Color {
	const iterations = 200
	for n := 0; n < iterations; n++ {

		if cmplx.Abs(f(z)) < 1e-7 {
			approxZ := complex(math.Round(real(z)*100)/100, math.Round(imag(z)*100)/100)
			c, ok := colorMap[approxZ]
			if !ok {
				c = colors[0]
				colorMap[approxZ] = c
				colors = colors[1:]
			}
			y, cb, cr := color.RGBToYCbCr(c.R, c.G, c.B)
			scale := math.Log(float64(n)) / math.Log(iterations)
			y -= uint8(float64(y) * scale)
			return color.YCbCr{y, cb, cr}
		}
		z = z - f(z)/fp(z)
	}
	return color.Black
}

func f(z complex128) complex128 {
	return z*z*z*z - 1
}

func fp(z complex128) complex128 {
	return 4 * z * z * z
}
