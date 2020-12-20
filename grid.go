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

// GenColorArray generates an interpolated color array from a palette
func GenColorArray(palette *[]color) *[]*sdl.Color {
	var newColors []*sdl.Color

	// iterate the palette
	for i, color := range *palette {

		// color 1 and 2 defined as the current color, and the next one (if it exists)
		c1, c2 := color, (*palette)[int(math.Min(float64(i+1), float64(len(*palette)-1)))]

		// steps is the last 8 bits of
		// the first of the two colors
		steps := c1 & 0xff

		// array of interpolated colors
		colors := make([]*sdl.Color, steps)

		// generate a color for each of the steps
		for i := 0; i < int(steps); i++ {
			var color [3]uint8

			// interpolate each color channel of both colors,
			// here we ignore the steps channel hence j = 1
			for j := 1; j < 4; j++ {
				// isolate the channel from each colors
				cn1 := (c1 >> (8 * j)) & 0xff
				cn2 := (c2 >> (8 * j)) & 0xff

				// ratio each of the channels according to the step count
				r := float64(i) / float64(steps)

				color[j-1] = uint8((1-r)*float64(cn1) + r*float64(cn2))
			}

			// the final interpolated color for step i
			colors[i] = &sdl.Color{
				R: color[2],
				G: color[1],
				B: color[0],
				A: 255,
			}
		}

		// append the iterpolated colors to a flattened array
		newColors = append(newColors, colors...)
	}

	return &newColors
}
