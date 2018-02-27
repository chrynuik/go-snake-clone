package main

import (
	"strconv"
)

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

	Snake := data.You
	Body := Snake.Body
	Tail := Body[len(Body)-1]

	for _, Snake := range data.Snakes {
		for _, Point := range Snake.Body {
			if Tail != Point {
				g.Grid[Point.Y][Point.X] = 9
			}
		}
	}

}

func (g Graph) isPointAccessible(point Point) bool {
	// does it exist? on the board?
	if point.X < 0 || point.X >= g.Width {
		return false
	} else if point.Y < 0 || point.Y >= g.Height {
		return false
	} else {
		return true
	}
}

func (g Graph) neighbors(point Point) []Point {
	// iterate over neighbors, is accesssible? return array of accessible neighbors
	directions := [4]Point{}
	directions[0] = Point{X: 1, Y: 0}
	directions[1] = Point{X: 0, Y: 1}
	directions[2] = Point{X: -1, Y: 0}
	directions[3] = Point{X: 0, Y: -1}

	results := []Point{}

	for _, direction := range directions {
		neighbor := Point{X: point.X + direction.X, Y: point.Y + direction.Y}

		if g.isPointAccessible(neighbor) && g.cost(neighbor) != 9 {
			results = append(results, neighbor)
		}
	}

	return results
}

func (g Graph) cost(direction Point) int {
	// the value in the point

	if g.isPointAccessible(direction) {
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
