// ex3.1 prints an svg image, ignoring non-finite vertexes.
package main

import (
	"fmt"
	"math"
)

const (
	width_1, height_1 = 600, 320                // canvas size in pixels
	cells_1           = 40                      // number of grid cells_1
	xyrange_1         = 30.0                    // axis ranges (-xyrange_1..+xyrange_1)
	xyscale_1         = width_1 / 2 / xyrange_1 // pixels per x or y unit
	zscale_1          = height_1 * 0.4          // pixels per z unit
	angle_1           = math.Pi / 6             // angle_1 of x, y axes (=30°)
)

var sin30_1, cos30_1 = math.Sin(angle_1), math.Cos(angle_1) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width_1: 0.7' "+
		"width_1='%d' height_1='%d'>", width_1, height_1)
	for i := 0; i < cells_1; i++ {
		for j := 0; j < cells_1; j++ {
			ax, ay := corner_eggbox(i+1, j)
			bx, by := corner_eggbox(i, j)
			cx, cy := corner_eggbox(i, j+1)
			dx, dy := corner_eggbox(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner_eggbox(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange_1 * (float64(i)/cells_1 - 0.5)
	y := xyrange_1 * (float64(j)/cells_1 - 0.5)

	// Compute surface height_1 z.
	z := eggbox(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width_1/2 + (x-y)*cos30_1*xyscale_1
	sy := height_1/2 + (x+y)*sin30_1*xyscale_1 - z*zscale_1
	return sx, sy
}

func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}
