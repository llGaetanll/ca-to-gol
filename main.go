package main

import (
	"fmt"
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	cellSize       = 3
	cellCountX     = 451
	cellCountY     = 251
	divisionY      = cellCountY - 2 // the Y level at which to divide GOL the generated CA. Note that this counts from the top of the screen
	frameSleepTime = time.Millisecond
	rule           = 1
)

type grid [cellCountX][cellCountY]uint

// colors are formatted as 0xRRGGBBSS (hex) where S is
// the number of interpolation steps with the next color
type color uint32

// cells of the grid
// if val == 0, cell is alive. any number after that is how many frames its been dead for
var cells grid
var colors []*sdl.Color

func main() {
	// store a color palette
	p := []color{
		0xffffff01, // first color is a live cell
		0x711C9108, // ...everything here is dead
		0xea00d91b, // and solely for aesthetic
		0x0adbc640, // purposes.
		0x133ea47d,
		0x00000001, // last color is black with step=1
	}

	// generate interpolated color array from the palette
	colors = *GenColorArray(&p)

	// initialize sdl
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	// create a window for the cell grid
	window, err := sdl.CreateWindow(
		"Cellular Automata",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		cellCountX*cellSize, cellCountY*cellSize,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()

	// create a renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	// blend colors to allow the alpha channel to work
	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	// fill cells slice with the uint max value by default
	for i, row := range cells {
		for j := range row {
			cells[i][j] = ^uint(0)
		}
	}

	// bottom middle cell is alive
	cells[cellCountX/2][cellCountY-1] = 0

	// generate random dots on first layer
	// for i := 0; i < cellCountX; i++ {
	// 	if isAlive := rand.Float64() > 0.96; isAlive {
	// 		cells[i][cellCountY-1] = 0
	// 	} else {
	// 		cells[i][cellCountY-1] = ^uint(0)
	// 	}
	// }

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				// stop the program when the user closes the window
				return
			}
		}

		// draw background as black
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		// next board state
		nextState := cells

		// iterate grid to draw,
		for i, row := range cells {
			for j := range row {
				// draw current cell
				DrawCell(renderer, i, j, cells[i][j])

				// both rulesets share the middle row

				// conway's game of life in upper half
				if float64(j) <= divisionY {
					// compute next state
					nextState[i][j] = LogicGOL(&cells, i, j, cells[i][j])
				}

				// shift up all rows in lower half by 1 each frame
				if float64(j) >= divisionY {
					nextState[i][j] = cells[i][int(math.Min(float64(j+1), float64(cellCountY-1)))]
				}
			}
		}

		// compute next generation of 1D CA on lowest row
		for i := 0; i < cellCountX; i++ {
			if isAlive := GetCellState(i, cellCountY-1, rule); isAlive {
				nextState[i][cellCountY-1] = 0
			} else {
				nextState[i][cellCountY-1] = ^uint(0)
			}
		}

		cells = nextState

		// render frame
		renderer.Present()

		// sleep until next frame
		time.Sleep(frameSleepTime)
	}
}
