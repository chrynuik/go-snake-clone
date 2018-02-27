package main

import "container/heap"

// PointCost is point cost
type PointCost struct {
	point Point
	cost  int
}

// Determines the approximate cost going from one cord to another
func hueristic(start Point, end Point) int {
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

func astar(g Graph, start Point, goal Point) []Point {
	toVisit := PriorityQueue{}
	heap.Init(&toVisit)

	item := &Item{
		priority: 0,
		point:    start,
	}

	heap.Push(&toVisit, item)

	cameFrom := make(map[Point]Point)
	costSoFar := make(map[Point]int)

	costSoFar[start] = 0

	for toVisit.Len() > 0 {
		item := heap.Pop(&toVisit).(*Item)

		if item.point == goal {
			break
		}

		neighbors := g.neighbors(item.point)

		for _, neighbor := range neighbors {
			newCost := costSoFar[item.point] + g.cost(neighbor)
			costSoFarNeighbor, ok := costSoFar[neighbor]

			if !ok || newCost < costSoFarNeighbor {
				costSoFar[neighbor] = newCost
				newItem := &Item{
					priority: newCost + hueristic(neighbor, goal),
					point:    neighbor,
				}

				heap.Push(&toVisit, newItem)
				cameFrom[neighbor] = item.point
			}
		}
	}

	point := goal
	path := make([]Point, 0)

	for point != start {
		if val, ok := cameFrom[point]; ok {
			path = append(path, point)
			point = val
		} else {
			path = make([]Point, 0)
			break
		}
	}

	return path
}

func contains(list []PointCost, a Point) bool {
	for _, b := range list {
		if b.point == a {
			return true
		}
	}
	return false
}

func getCost(list []PointCost, a Point) int {
	for _, b := range list {
		if b.point == a {
			return b.cost
		}
	}
	return 0
}
