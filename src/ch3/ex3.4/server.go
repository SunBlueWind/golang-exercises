// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

var (
	width, height = 600, 320                     // canvas size in pixels
	cells         = 100                          // number of grid cells
	xyrange       = 30.0                         // axis ranges (-xyrange..+xyrange)
	xyscale       = float64(width) / 2 / xyrange // pixels per x or y unit
	zscale        = float64(height) * 0.4        // pixels per z unit
	angle         = math.Pi / 6                  // angle of x, y axes (=30°)
	color         = "#ffffff"
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

var maxHeight = 1.
var minHeightX = math.Pi * 1.5
var minHeight = math.Sin(minHeightX) / minHeightX
var midHeight = (maxHeight + minHeight) / 2

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	if heightParam, ok := r.Form["height"]; ok {
		if parsedHeight, err := strconv.Atoi(heightParam[0]); err == nil {
			height = parsedHeight
		}
	}
	if widthParam, ok := r.Form["width"]; ok {
		if parsedWidth, err := strconv.Atoi(widthParam[0]); err == nil {
			width = parsedWidth
		}
	}
	if colorParam, ok := r.Form["color"]; ok {
		color = colorParam[0]
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w)
}

func surface(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, okA := corner(i+1, j)
			bx, by, okB := corner(i, j)
			cx, cy, okC := corner(i, j+1)
			dx, dy, okD := corner(i+1, j+1)
			if !okA || !okB || !okC || !okD {
				continue
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)
	// Compute surface height z.
	z := f(x, y)
	if math.IsInf(z, 0) {
		return 0, 0, false
	}
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
