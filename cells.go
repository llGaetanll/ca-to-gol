package main

// this file contains the logic to generate any cellular automata rule

func isAlive(val uint) bool {
	return val == 0
}

// Note: only the bottom most row actually generates cells

// GetCellState calculates the state of the cell given its rule number
func GetCellState(x, y int, rule uint8) bool {
	// neighbors - all dead by default
	n := make([]bool, 3)

	// WRAP AROUND
	// only check left neighbor if this isn't an edge cell
	if x > 0 {
		n[0] = isAlive(cells[x-1][y])
	} else {
		// wrap around and check last cell
		n[0] = isAlive(cells[cellCountX-1][y])
	}

	// check middle neighbor
	n[1] = isAlive(cells[x][y])

	// only check right neighbor if this isn't an edge cell
	if x < cellCountX-1 {
		n[2] = isAlive(cells[x+1][y])
	} else {
		// wrap around and check first cell
		n[2] = isAlive(cells[0][y])
	}

	// now that we know the state of all 3 neighbors,
	// calculate the state of the cell based on the rule count

	i := 0
	// iterate over all bits of the rule
	// his checks every sub-rule of the rule
	for rule > 0 {

		// if the current bit is 1, check ith subrule
		if rule&1 == 1 {
			// c is the current configuration
			// p is a copy of i to shift over
			c, p := i, 0

			// by default, rule matches
			match := true

			// eval current subrule
			// eval bits of i to see what neighbors to check
			for range n {
				// if the current bit is 1 but neighbor is dead,
				// if the current but is 0 but neighbor is alive,
				// subrule fails
				if c&1 == 1 && !n[p] || c&1 == 0 && n[p] {
					match = false
				}

				// shift config over to check next bit
				c >>= 1
				p++
			}

			// if subrule still matches at this point, and all the neighbors have been checked, we are done
			// remember that as long as one subrule matches, the rule passes
			if match {
				return true
			}
		}

		// shift rule over
		rule >>= 1
		i++
	}

	// if we reach this point, it means we have checked every subrule of the current rule, and none matche
	return false
}
