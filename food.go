package main

import (
	"container/heap"
)

func getFoodPath(food PointList, us Snake, pathToTail []Point) *Item {
	closestFood := PriorityQueue{}

	for _, morsel := range food {
		newItem := &Item{
			priority: hueristic(us.Body[0], morsel),
			point:    morsel,
		}

		if len(pathToTail) > 0 && us.Body[len(us.Body)-1] != us.Body[1] {
			heap.Push(&closestFood, newItem)
		}
	}

	if len(closestFood) > 0 {
		return heap.Pop(&closestFood).(*Item)
	}

	return nil
}
