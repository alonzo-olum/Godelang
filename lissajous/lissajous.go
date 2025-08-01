package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.RGBA{0x66, 0xff, 0x66, 0xff}, color.RGBA{0x00, 0x80, 0x80, 0}, color.RGBA{0xff, 0xff, 0x00, 0x00}}

const (
	whiteIndex = 0
	blackIndex = 1
)

func lissajous(out io.Writer, cycles int) {
	const (
		//cycles = 5	number of complete x oscillator revs
		res = 0.001	// angular resolution
		size = 100	// image canvas covers[-size..+size]
		nframes = 64	// animation frames
		delay = 8	// delay between frames
	)

	freq := rand.Float64() * 3.0	// relative frequency of y-oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	if cycles == 0 {
		cycles = 5
	}
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(3) % 3))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
