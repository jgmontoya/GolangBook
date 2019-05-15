// Server1 is a minimal "echo" server.
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
)

var palette = []color.Color{
	color.White,
	color.RGBA{26, 188, 156, 1.0},
	color.RGBA{230, 126, 34, 1.0},
	color.RGBA{52, 152, 219, 1.0},
	color.RGBA{155, 89, 182, 1.0},
	color.RGBA{52, 73, 94, 1.0},
}

func main() {
	const ( // const declared inside a function are only visible within that function
		_cycles  = 5      // number of complete x oscillator revolutions
		_res     = 0.0001 // angular resolution
		_size    = 100    // image canvas covers [-size..+size]
		_nframes = 128    // number of animation frames
		_delay   = 4      // delay between frames in 10ms units
	)
	options := make(map[string]float64)
	// http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		options["cycles"] = _cycles
		options["res"] = _res
		options["size"] = _size
		options["nframes"] = _nframes
		options["delay"] = _delay
		for option, value := range r.Form {
			options[option], _ = strconv.ParseFloat(value[0], 64)
		}
		lissajous(w, int(options["cycles"]), options["res"], int(options["size"]), int(options["nframes"]), int(options["delay"]))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles int, res float64, size int, nframes int, delay int) {
	freq := rand.Float64() * 3.0        // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} // struct of type gif.GIF, all other fields besides LoopCount have the zero value of their type
	phase := 0.0                        // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			//colorIndex := rand.Intn(5) + 1
			//img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colorIndex))
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(t/res)%5+1)
		}
		// phase += 0.1
		phase += rand.Float64() * 0.2
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: Ignoring encoding errors
}
