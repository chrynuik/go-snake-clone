package main

import (
	"strconv"
)

// Tile is a struct
type Tile struct {
	X int
	Y int
}

// Graph is a struct
type Graph struct {
	Width  int
	Height int
	Grid   [][]int
}

func (g *Graph) create(data *MoveRequest) {
	g.Width = data.Width
	g.Height = data.Height
	g.Grid = make([][]int, g.Height)

	for i := range g.Grid {
		g.Grid[i] = make([]int, g.Width)

		for j := range g.Grid[i] {
			g.Grid[i][j] = 1
		}
	}

	for _, Snake := range data.Snakes {
		for _, Point := range Snake.Body {
			g.Grid[Point.Y][Point.X] = 9
		}
	}

}

func (g Graph) isTileAccessible(tile Tile) bool {
	// does it exist? on the board?
	if tile.X < 0 || tile.X >= g.Width {
		return false
	} else if tile.Y < 0 || tile.Y >= g.Height {
		return false
	} else {
		return true
	}
}

func (g Graph) neighbors(tile Tile) []Tile {
	// iterate over neighbors, is accesssible? return array of accessible neighbors
	directions := [4]Tile{}
	directions[0] = Tile{X: 1, Y: 0}
	directions[1] = Tile{X: 0, Y: 1}
	directions[2] = Tile{X: -1, Y: 0}
	directions[3] = Tile{X: 0, Y: -1}

	results := []Tile{}

	for _, direction := range directions {
		neighbor := Tile{X: tile.X + direction.X, Y: tile.Y + direction.Y}

		if g.isTileAccessible(neighbor) && g.cost(neighbor) != 9 {
			results = append(results, neighbor)
		}
	}

	return results
}

func (g Graph) cost(direction Tile) int {
	// the value in the tile

	if g.isTileAccessible(direction) {
		return g.Grid[direction.Y][direction.X]
	}

	return 9
}

func (g Graph) String() string {
	graph := ""

	for i := range g.Grid {
		for j := range g.Grid[i] {
			graph += strconv.Itoa(g.Grid[i][j])
			graph += ""
		}
		graph += "\n"
	}

	return graph
}
