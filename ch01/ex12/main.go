// Exercise 1.12: Modify the Lissajous server to read parameter values from the URL. For example,
// you might arrange it so that a URL like http://localhost:8000/?cycles=20 sets the number of
// cycles to 20 instead of the default 5. Use the strconv.Atoi function to convert the string
// parameter into an integer. You can see its documentation with go doc strconv.Atoi.
// Is a server that reads parameter values from a URL and generates the corresponding Lissajous figure.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		cycles := parseFirstIntOrDefault(r.Form["cycles"], 5)
		res := parseFirstFloat64OrDefault(r.Form["res"], 0.001)
		size := parseFirstIntOrDefault(r.Form["size"], 100)
		nframes := parseFirstIntOrDefault(r.Form["nframes"], 64)
		delay := parseFirstIntOrDefault(r.Form["delay"], 8)
		lissajous(w, cycles, res, size, nframes, delay)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

// parseFirstFloat64OrDefault Parses the first element of the given array of strings into an integer and returns it.
// If no element can be parsed, it returns the default value given.
func parseFirstIntOrDefault(array []string, defaultValue int) int {
	if len(array) < 1 {
		return defaultValue
	}
	value, err := strconv.Atoi(array[0])
	if err != nil {
		return defaultValue
	}
	return value
}

// parseFirstFloat64OrDefault Parses the first element of the given array of strings into a floating point number and returns it.
// If no element can be parsed, it returns the default value given.
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
