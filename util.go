package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

// this file contains helpers to do with the grid & colors

// DrawCell draws a cell on the grid at coordinates x, y
func DrawCell(r *sdl.Renderer, x, y int, state uint) {
	cell := sdl.Rect{
		X: int32(x) * cellSize,
		Y: int32(y) * cellSize,
		W: cellSize,
		H: cellSize,
	}

	// map color index to [0, len(colors) - 1)
	color := colors[int(math.Min(float64(state), float64(len(colors)-1)))]

	r.SetDrawColor(
		color.R,
		color.G,
		color.B,
		255,
	)

	r.FillRect(&cell)
}

// GenColorArray generates an interpolated color array from the basic palette
func GenColorArray(colors *[]color) *[]*sdl.Color {
	var newColors []*sdl.Color

	// interpolate between 2 colors and return an array of length steps with the interpolated colors
	interColors := func(a, b color, steps uint) []*sdl.Color {
		colors := make([]*sdl.Color, steps)

		for i := 0; i < int(steps); i++ {
			ratio := func(c1, c2 uint8) uint8 {
				r := float64(i) / float64(a.S)

				return uint8((1-r)*float64(c1) + r*float64(c2))
			}

			colors[i] = &sdl.Color{
				R: ratio(a.R, b.R),
				G: ratio(a.G, b.G),
				B: ratio(a.B, b.B),
				A: 255,
			}
		}

		return colors
	}

	for i, color := range *colors {
		colors := interColors(color, (*colors)[int(math.Min(float64(i+1), float64(len(*colors)-1)))], color.S)
		newColors = append(newColors, colors...)
	}

	return &newColors
}
