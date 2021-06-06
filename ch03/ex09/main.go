// ch03 / ex09 is a server that reads parameter values from URLs and generates images of the corresponding Mandelbrot fractals.
// Corresponds to the following parameters through URL queries:
// - x    : Of the center point x Coordinate
// - y    : Of the center point y Coordinate
// - zoom : magnification (2 ^ (zoom - 1) Times)
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		x := parseFirstFloat64OrDefault(r.Form["x"], 0)
		y := parseFirstFloat64OrDefault(r.Form["y"], 0)
		zoom := parseFirstFloat64OrDefault(r.Form["zoom"], 0)
		renderPNG(w, x, y, zoom)
	}
	http.HandleFunc("/", handler)

	fmt.Println("Listening at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func renderPNG(out io.Writer, x, y, zoom float64) {
	const (
		width, height = 1024, 1024
	)

	m := math.Exp2(1 - zoom)
	xmin, xmax := x-m, x+m
	ymin, ymax := y-m, y+m

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// parseFirstFloat64OrDefault parses and returns the first element of the given array of strings to a floating point number.
// If none of the elements can be parsed, it returns the given default value.
func parseFirstFloat64OrDefault(array []string, defaultValue float64) float64 {
	if len(array) < 1 {
		return defaultValue
	}
	value, err := strconv.ParseFloat(array[0], 64)
	if err != nil {
		return defaultValue
	}
	return value
}
