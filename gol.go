package main

// this file contains the logic for conway's game of life

// GetNeighborCount finds the number of alive neighbors
func getNeighborCount(i, j int) int {
	liveCount := 0

	// check in a 3x3 radius around the point
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x >= 0 && y >= 0 && x < cellCountX && y < cellCountY && !(x == i && y == j) && cells[x][y] == 0 {
				liveCount++
			}
		}
	}

	return liveCount
}

// LogicGOL given a cell at coordinates x, y as well as its status,
// calculates if that cell should die or live in the next iteration of the board
func LogicGOL(x, y int, cellVal uint) uint {
	numNeighbors := getNeighborCount(x, y)

	// if the cell is alive and has too few or too many neighbors, it dies next frame
	if cellVal == 0 && (numNeighbors < 2 || numNeighbors > 3) {
		return 1
	}

	// if the cell is dead
	if cellVal != 0 {
		// if it has 3 neighbors
		if numNeighbors == 3 {
			// lives next frame
			return 0
		} else if cellVal < ^uint(0) {
			// stays dead, no rollovers
			return cells[x][y] + 1
		}
	}

	return cells[x][y]
}
