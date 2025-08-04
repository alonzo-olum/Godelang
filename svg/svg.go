package main

import (
	"fmt"
	"math"
	"io"
)

const (
	cells         = 100                   //number of grid cells
	xyrange       = 30.0                 //axis ranges (-xyrange..+xyrange) 
	angle         = math.Pi / 6       // angle of x, y axes (30 degrees)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)    // sin(30), cos(30)

func f(x, y float64) (float64, error) {
	r := math.Hypot(x, y)    // distance from (0,0)
	if math.IsInf(r, 1) || math.IsNaN(r) || r == 0{
		return float64(1), fmt.Errorf("non-finite value  of hypoteneuse!")
	}
	return math.Sin(r) / r, nil
}

func corner(i, j, width, height int, xyscale, zscale float64) (float64, float64, float64) {
	// Find point (x, y) at corner of cell(i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z
	z, err := f(x,y)
	if err != nil {
		return float64(width) / 2, float64(height) / 2, 0
	}
	// Project (x, y, z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func svg(w io.Writer, color string, height, width int) {
	xyscale       := float64(width) / 2 / xyrange // pixels per x or y unit
	zscale        := float64(height) * 0.4       // pixels per z unit
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
	"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
	"width='%d' height='%d'>", color, width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, width, height, xyscale, zscale)
			bx, by, bz := corner(i, j, width, height, xyscale, zscale)
			cx, cy, cz := corner(i, j+1, width, height, xyscale, zscale)
			dx, dy, dz := corner(i+1, j+1, width, height, xyscale, zscale)
			zAvg := (az + bz + cz + dz) / 4
			fill := "red"
			if zAvg < 0 {
				fill = "blue"
			}
			fmt.Fprintf(w, "<polygon points='%g, %g %g, %g %g, %g %g, %g' style='fill:%v;fill-opacity:%v'/>\n", ax, ay, bx, by, cx, cy, dx, dy, fill, math.Abs(zAvg/0.4))
		}
	}
	fmt.Fprintf(w, "</svg>")
}
