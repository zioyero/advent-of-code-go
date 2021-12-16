package days

import (
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// Eleven is Day 11 of Advent of Code!
type Eleven struct {
}

// PartOne counts the flashes that happen over 100 steps
func (e Eleven) PartOne(lines []string) int {
	input := parseInput(lines)
	garden := octopusGarden{
		optopedes: input,
	}
	garden.print()
	sum := 0
	for x := 0; x < 100; x++ {
		flashes := garden.step()
		sum += flashes
		garden.print()
	}
	return sum
}

func parseInput(lines []string) [][]int {
	input := make([][]int, 0, len(lines))
	for _, line := range lines {
		energyLevelsRaw := strings.Split(line, "")
		energyLevels := make([]int, 0, len(energyLevelsRaw))
		for _, rawLevel := range energyLevelsRaw {
			level, _ := strconv.Atoi(rawLevel)
			energyLevels = append(energyLevels, level)
		}
		input = append(input, energyLevels)
	}
	return input
}

// PartTwo calculates the first step where the octopedes are synchronized
func (e Eleven) PartTwo(lines []string) int {
	input := parseInput(lines)
	garden := octopusGarden{
		optopedes: input,
	}
	garden.print()
	found := false
	x := 0
	for !found {
		flashes := garden.step()
		garden.print()
		x++
		found = flashes == 100
	}
	return x
}

type octopusGarden struct {
	optopedes [][]int
}

type position struct {
	x int
	y int
}

func (og *octopusGarden) print() {
	flashed := color.New(color.FgHiWhite)
	ready := color.New(color.FgHiYellow)
	normal := color.New(color.FgHiCyan)
	for _, row := range og.optopedes {
		for _, v := range row {
			if v == 0 {
				flashed.Printf("%d", v)
			} else if v >= 10 {
				ready.Printf("*")
			} else {
				normal.Printf("%d", v)
			}
		}
		normal.Printf("\n")
	}
}

func (og *octopusGarden) step() int {
	// increment the energy levels for all octopedes by 1
	og.incrementEnergyLevels()
	og.print()
	// check for octopedes that are ready to flash
	flashable := og.flashReadyOctopedes()
	if len(flashable) == 0 {
		return 0
	}
	// flash these and propagate the energy
	return og.flash(flashable, make([]position, 0))
}

func (og *octopusGarden) incrementEnergyLevels() {
	for x, row := range og.optopedes {
		for y := range row {
			og.optopedes[x][y]++
		}
	}
}

func (og *octopusGarden) flashReadyOctopedes() []position {
	positions := make([]position, 0)
	for x, row := range og.optopedes {
		for y := range row {
			if og.optopedes[x][y] > 9 {
				positions = append(positions, position{x, y})
			}
		}
	}
	return positions
}

func (og *octopusGarden) flash(flashable []position, alreadyFlashed []position) int {
	// first, remove any flashable position that has already flashed
	positionsThatCanFlash := filterOut(flashable, alreadyFlashed)
	// if there are no more flashable positions, we can return
	if len(positionsThatCanFlash) == 0 {
		return len(alreadyFlashed)
	}
	// set each flashable position back to 0
	for _, f := range positionsThatCanFlash {
		og.optopedes[f.x][f.y] = 0
	}
	// add these flashed positions to the already flashed list
	alreadyFlashed = append(alreadyFlashed, positionsThatCanFlash...)
	// propagate the energy
	og.propagateFlashEnergyFrom(positionsThatCanFlash, alreadyFlashed)
	// check if this propagation has put any more octopedes in a flash-ready state
	next := og.flashReadyOctopedes()
	if len(next) == 0 {
		return len(alreadyFlashed)
	}
	return og.flash(next, alreadyFlashed)
}

func filterOut(from []position, badElements []position) []position {
	filtered := make([]position, 0, len(from))
	for _, p := range from {
		if !positionsContains(badElements, p) {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func positionsContains(positions []position, p position) bool {
	for _, c := range positions {
		if p.x == c.x && p.y == c.y {
			return true
		}
	}
	return false
}

func (og *octopusGarden) propagateFlashEnergyFrom(positions []position, alreadyFlashed []position) {
	for _, p := range positions {
		// get the surrounding area for each position
		neighbors := p.neighbors()
		// filter out the ones that have already flashed
		filteredNeighbors := filterOut(neighbors, alreadyFlashed)
		// increment their level by 1
		for _, n := range filteredNeighbors {
			og.optopedes[n.x][n.y]++
		}
	}
}

func (p position) neighbors() []position {
	candidates := []position{
		// LEFT
		{x: p.x - 1, y: p.y},
		{x: p.x - 1, y: p.y - 1},
		{x: p.x - 1, y: p.y + 1},
		// RIGHT
		{x: p.x + 1, y: p.y},
		{x: p.x + 1, y: p.y - 1},
		{x: p.x + 1, y: p.y + 1},
		// TOP
		{x: p.x, y: p.y - 1},
		// BOTTOM
		{x: p.x, y: p.y + 1},
	}
	// filter out the out of bounds neighbors
	viable := make([]position, 0, len(candidates))
	for _, c := range candidates {
		if c.x >= 0 && c.x < 10 && c.y >= 0 && c.y < 10 {
			viable = append(viable, c)
		}
	}
	return viable
}
