package main

import "container/heap"

// TileCost is tile cost
type TileCost struct {
	tile Tile
	cost int
}

// Determines the approximate cost going from one cord to another
func hueristic(start Tile, end Tile) int {
	xVal := start.X - end.X
	yVal := start.Y - end.Y

	if xVal < 0 {
		xVal = -1 * xVal
	}

	if yVal < 0 {
		yVal = -1 * yVal
	}

	heuristic := xVal + yVal

	return heuristic
}

func astar(g Graph, start Tile, goal Tile) []Tile {
	toVisit := PriorityQueue{}
	heap.Init(&toVisit)

	item := &Item{
		priority: 0,
		tile:     start,
	}

	heap.Push(&toVisit, item)

	cameFrom := make(map[Tile]Tile)
	costSoFar := make(map[Tile]int)

	costSoFar[start] = 0

	for toVisit.Len() > 0 {
		item := heap.Pop(&toVisit).(*Item)

		if item.tile == goal {
			break
		}

		neighbors := g.neighbors(item.tile)

		for _, neighbor := range neighbors {
			newCost := costSoFar[item.tile] + g.cost(neighbor)
			costSoFarNeighbor, ok := costSoFar[neighbor]

			if !ok || newCost < costSoFarNeighbor {
				costSoFar[neighbor] = newCost
				newItem := &Item{
					priority: newCost + hueristic(neighbor, goal),
					tile:     neighbor,
				}

				heap.Push(&toVisit, newItem)
				cameFrom[neighbor] = item.tile
			}
		}
	}

	tile := goal
	path := make([]Tile, 0)

	for tile != start {
		if val, ok := cameFrom[tile]; ok {
			path = append(path, tile)
			tile = val
		} else {
			path = make([]Tile, 0)
			break
		}
	}

	return path
}

func contains(list []TileCost, a Tile) bool {
	for _, b := range list {
		if b.tile == a {
			return true
		}
	}
	return false
}

func getCost(list []TileCost, a Tile) int {
	for _, b := range list {
		if b.tile == a {
			return b.cost
		}
	}
	return 0
}
