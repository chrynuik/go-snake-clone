package main

import "strconv"

// Graph is a struct
type Graph struct {
	Width  int
	Height int
	Grid   [][]int
}

func (g *Graph) create(width int, height int) {
	g.Width = width
	g.Height = height
	g.Grid = make([][]int, height)

	for i := range g.Grid {
		g.Grid[i] = make([]int, width)

		for j := range g.Grid[i] {
			g.Grid[i][j] = 1
		}
	}
}

func (g Graph) isTileAccessible(tile Tile) bool {
	// does it exist? on the board?
}

func (g Graph) neighbors(tile Tile) []Tile {
	// iterate over neighbors, is accesssible? return array of accessible neighbors
}

func (g Graph) cost(tile Tile, direction Tile) {
	// the value in the tile
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
