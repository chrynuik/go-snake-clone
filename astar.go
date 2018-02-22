package main

import (
	"fmt"
	"math"
)

// TileCost is tile cost
type TileCost struct {
	tile Tile
	cost int
}

// Determines the approximate cost going from one cord to another
func hueristic(start Tile, end Tile) int {
	return int(math.Abs(float64(start.X-end.X)) + math.Abs(float64(start.Y-end.Y)))
}

func astar(g Graph, start Tile, goal Tile) []Tile {
	pq := make(PriorityQueue, 0)
	currentTile := start
	visitedTiles := []TileCost{}

	firstItem := &Item{
		priority: 0,
		index:    0,
		tile:     start,
	}

	// heap.Push(&pq, firstItem)
	pq.Push(firstItem)

	fmt.Println("HELLO?")
	fmt.Println(pq.Len())

	// fmt.Println("queue", pq[0])
	for pq.Len() > 0 {
		currentTile = pq[0].tile

		if currentTile == goal {
			break
		}

		possibleDirections := g.neighbors(currentTile)
		fmt.Println("POSSIBLE DIRECTION", possibleDirections)

		for index, i := range possibleDirections {

			fmt.Println("==============")
			fmt.Println("index", index)

			// newCost := g.cost(i) + g.cost(goal)
			newCost := g.cost(currentTile) + g.cost(i)
			fmt.Println("newCost", newCost)
			fmt.Println("queue length", pq.Len())

			// if !contains(visitedTiles, i) || newCost < getCost(visitedTiles, i) {
			fmt.Println("WERE IN THE LOOP")
			visitedTiles = append(visitedTiles, TileCost{tile: i, cost: newCost})
			priority := newCost + hueristic(currentTile, i)
			fmt.Println("priority", priority)

			toPush := &Item{
				tile:     i,
				priority: priority,
			}
			// heap.Push(&pq, toPush)
			pq.Push(toPush)

			// }
			// fmt.Println("start", start)
			// fmt.Println(g.cost(i))
			// fmt.Println(newCost)
		}

	}

	path := make([]Tile, 1)
	path[0] = Tile{
		X: 1,
		Y: 1,
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
