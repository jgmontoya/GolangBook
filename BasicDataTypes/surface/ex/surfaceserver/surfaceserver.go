// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

type gridpoint struct {
	ax, ay, bx, by, cx, cy, dx, dy, z float64
}

var sin30, cos30 = math.Sin(math.Pi / 6), math.Cos(math.Pi / 6) // sin(30°), cos(30°)
var min = math.Inf(1)
var max = math.Inf(-1)

func main() {
	const (
		_width, _height = 600, 320               // canvas size in pixels
		_cells          = 100                    // number of grid cells
		_xyrange        = 30.0                   // axis ranges (-xyrange..+xyrange)
		_xyscale        = _width / 2 / _xyrange  // pixels per x or y unit
		_zscale         = float64(_height) * 0.4 // pixels per z unit
		_angle          = math.Pi / 6            // angle of x, y axes (30°)
	)
	options := make(map[string]float64)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		r.ParseForm()
		options["width"] = _width
		options["height"] = _height
		options["cells"] = _cells
		options["xyrange"] = _xyrange
		options["xyscale"] = _xyscale
		options["zscale"] = _zscale
		options["angle"] = _angle
		for option, value := range r.Form {
			options[option], _ = strconv.ParseFloat(value[0], 64)
		}
		surface(w, int(options["width"]), int(options["height"]), int(options["cells"]), float64(options["xyrange"]), float64(options["xyscale"]), float64(options["zscale"]), float64(options["angle"]))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(out io.Writer, width, height, cells int, xyrange, xyscale, zscale, angle float64) {
	grid := make([][]gridpoint, cells) // We need dynamically allocated grid
	for row := range grid {
		grid[row] = make([]gridpoint, cells)
	}
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height, cells, xyrange, xyscale, zscale)
			bx, by := corner(i, j, width, height, cells, xyrange, xyscale, zscale)
			cx, cy := corner(i, j+1, width, height, cells, xyrange, xyscale, zscale)
			dx, dy := corner(i+1, j+1, width, height, cells, xyrange, xyscale, zscale)
			if areFinite(ax, ay, bx, by, cx, cy, dx, dy) {
				z := gridF(i, j, cells, xyrange)
				switch {
				case z < min:
					min = z
				case z > max:
					max = z
				}
				grid[i][j] = gridpoint{ax, ay, bx, by, cx, cy, dx, dy, z}
			}
		}
	}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			point := grid[i][j]
			if areFinite(point.ax, point.ay, point.bx, point.by, point.cx, point.cy, point.dx, point.dy) {
				printPolyPoint(out, point)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j, width, height, cells int, xyrange, xyscale, zscale float64) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

// Variadic function explained in chapter 5: Functions can receive arbitrary number of parameters
func areFinite(points ...float64) bool {
	for _, point := range points {
		if !isFinite(point) {
			return false
		}
	}
	return true
}

func isFinite(x float64) bool {
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return false
	}
	return true
}

func colorMap(num float64) string {
	num -= min
	num /= (max - min)
	r := int(math.Round(num * 255))
	b := int(math.Round((1 - num) * 255))
	return fmt.Sprintf("rgb(%d, 00, %d)", r, b)
}

func gridF(i, j, cells int, xyrange float64) float64 {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	return f(x, y)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0.0)
	return math.Sin(r) / r
}

func printPolyPoint(out io.Writer, point gridpoint) {
	fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' "+
		"style='fill: %s' />\n",
		point.ax, point.ay, point.bx, point.by, point.cx, point.cy, point.dx, point.dy, colorMap(point.z))
}
