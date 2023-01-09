# Elementary Cellular Automata to Conway's Game of Life

![ca-to-gol](https://i.imgur.com/4mapZRt.png "Cellular Automata fed to Conway's Game of Life")

Any Cellular Automaton rule fed as input to Conway's Game of Life. This project is inspired by [this amazing video](https://www.youtube.com/watch?v=IK7nBOLYzdE) by Elliot Waite.

This version of the program is not limited to rule 30, and can generate any of the 256 rules possible from [elementary cellular automata](https://en.wikipedia.org/wiki/Elementary_cellular_automaton).

## Running the Code
Make sure you have Go installed on your system

```
go run *.go
```

This should take some time as go collects the various files used by `sdl`, which is the rendering engine used in the program.

## Parameters

You can tweak the following output parameters. These can be found at the top of the `main.go` file.

- `cellSize`: The sidelength in pixels of the cells of the output grid.
- `cellCountX`: The width of the output grid in cells
- `cellCountY`: The height of the output grid in cells
- `divisionY`: The layer at which the board rules change from Elementary Cellular Automata to Conway's Game of Life
- `frameSleepTime`: How long the program should halt between each frame. Allows to speed up or slow down the output
- `rule`: The Elementary Cellular Automaton rule number to generate at the bottom of the grid

Note that you'll have to restart the program everytime you want these changes to take effect.

## Notes

The colors can also be changed, the list can be found at the top of the `main` function in `main.go`

```go
func main() {
  p := []color{
    0xffffff01, // first color is a live cell
    0x711C9108, // ...everything here is dead
    0xea00d91b, // and solely for aesthetic
    0x0adbc640, // purposes.
    0x133ea47d,
    0x00000001, // last color is black with step=1
  }
  
  // ...
}
```
Each color is specified as a hex value. The given channels are `RGBA`.
