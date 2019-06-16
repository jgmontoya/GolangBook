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

const iterations = 200

var hues [iterations + 1]float64

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	histogram := make(map[float64]int)
	values := make(map[Pair]float64, width*height)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			count := mandelbrot(z)

			values[Pair{px, py}] = count
			if _, ok := histogram[math.Floor(count)]; ok && count < iterations {
				histogram[math.Floor(count)]++
			} else if count < iterations {
				histogram[math.Floor(count)] = 1
			}

		}
	}
	total := 0
	for _, val := range histogram {
		total += val
	}

	h := 0.0
	index := 0
	for i := 0.0; i < iterations; i++ {
		if val, ok := histogram[i]; ok {
			h += float64(val) / float64(total)
		}
		hues[index] = h
		index++
	}
	hues[index] = h

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {

			count := values[Pair{px, py}]

			img.Set(px, py, colorize(count))
		}
	}
	png.Encode(os.Stdout, img) // Note: ignoring errors
}

// Pair models (X, Y) as a tuple to be used as a key in a map
type Pair struct {
	X, Y int
}

func mandelbrot(z complex128) float64 {
	// mandelbrot returns the renormalized iteration count
	const contrast = 15

	var v complex128
	var n int
	for n = 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return float64(n)
		}
	}
	if n == iterations {
		return iterations
	}
	return float64(n) + 1 - math.Log10(math.Log(cmplx.Abs(v)))
}

func colorize(count float64) color.Color {
	hue := float64(1 - linearInterpolation(hues[int(math.Floor(count))], hues[int(math.Ceil(count))], math.Mod(count, 1.0)))
	saturation := 1.0
	value := 0.0
	if count < iterations {
		value = 1.0
	}
	return hsv2rgb(hue, saturation, value)
}

func hsv2rgb(hue, sat, val float64) color.Color {
	var r, g, b float64

	i := math.Floor(hue * 6)
	f := hue*6 - i
	p := val * (1 - sat)
	q := val * (1 - f*sat)
	t := val * (1 - (1-f)*sat)

	switch int(i) % 6 {
	case 0:
		r, g, b = val, t, p
	case 1:
		r, g, b = q, val, p
	case 2:
		r, g, b = p, val, t
	case 3:
		r, g, b = p, q, val
	case 4:
		r, g, b = t, p, val
	case 5:
		r, g, b = val, p, q
	}

	return color.RGBA{uint8(255 * r), uint8(255 * g), uint8(255 * b), 255}
}

func linearInterpolation(color1, color2, t float64) float64 {
	return color1*(1-t) + color2*t
}
