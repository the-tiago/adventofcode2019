package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coords struct {
	x int
	y int
}

type Cell struct {
	coords Coords
	first  bool
	second bool
}

type Grid struct {
	cells         map[string]Cell
	interceptions map[string]Cell
}

func main() {

}

func Day3(wire1, wire2 string) int {
	grid := Grid{map[string]Cell{}, map[string]Cell{}}
	grid.AddWire(wire1, true)
	grid.AddWire(wire2, false)

	closestIterception := 0
	root := Coords{0, 0}
	for key, cell := range grid.interceptions {
		distance := distance(root, cell.coords)
		fmt.Printf("interception @%s, distance %d\n", key, distance)
		if closestIterception == 0 || distance < closestIterception {
			closestIterception = distance
		}
	}
	//fmt.Printf("%#v", grid.interceptions)
	return closestIterception
}

func distance(c1, c2 Coords) int {
	return int(math.Abs(float64((c2.x - c1.x))) + math.Abs(float64((c2.y - c1.y))))
}

func (c Coords) HashMap() string {
	return strconv.Itoa(c.x) + "," + strconv.Itoa(c.y)
}

func (g *Grid) AddWire(wire string, first bool) {
	wireComponents := strings.Split(wire, ",")
	x, y := 0, 0
	for _, s := range wireComponents {
		direction := s[0:1]
		count, _ := strconv.Atoi(s[1:])
		switch direction {
		case "U":
			for i := 0; i < count; i++ {
				y++
				g.addOrUpdateCell(x, y, first)
			}
		case "D":
			for i := 0; i < count; i++ {
				y--
				g.addOrUpdateCell(x, y, first)
			}
		case "R":
			for i := 0; i < count; i++ {
				x++
				g.addOrUpdateCell(x, y, first)
			}
		case "L":
			for i := 0; i < count; i++ {
				x--
				g.addOrUpdateCell(x, y, first)
			}
		}
	}
}

func (g *Grid) addOrUpdateCell(x, y int, first bool) {
	key := strconv.Itoa(x) + "," + strconv.Itoa(y)
	cell, ok := g.cells[key]
	if !ok {
		cell = Cell{Coords{x, y}, false, false}
	}
	cell.first = first || cell.first
	cell.second = !first || cell.second
	g.cells[key] = cell

	if cell.first && cell.second {
		g.interceptions[key] = cell
	}
}
