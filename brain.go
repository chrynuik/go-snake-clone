package main

import (
	"container/heap"
	"fmt"
)

func handleMove(data *MoveRequest) string {
	directions := []string{
		"up",
		"down",
		"left",
		"right",
	}

	Food := data.Food

	Us := data.You
	Body := Us.Body
	Head := Body[0]
	Tail := Body[len(Body)-1]

	Health := Us.Health

	AllSnakes := data.Snakes

	closestFood := PriorityQueue{}

	for _, Morsel := range Food {
		newItem := &Item{
			priority: hueristic(Tile{X: Head.X, Y: Head.Y}, Tile{X: Morsel.X, Y: Morsel.Y}),
			tile:     Tile{X: Morsel.X, Y: Morsel.Y},
		}

		heap.Push(&closestFood, newItem)
	}

	Goal := heap.Pop(&closestFood).(*Item).tile

	if Health < 99 && Health > 50 {
		Goal = Tile{X: Tail.X, Y: Tail.Y}
	}

	board := Graph{}
	board.create(data)

	enemyHeads := getEnemyHeads(AllSnakes, Us)
	attackableEnemies := getAttackableEnemies(enemyHeads, 6)

	fmt.Println("ALL ENEMIES", enemyHeads)
	fmt.Println("ATTACKABLE ENEMIES", attackableEnemies)

	fmt.Println(board)
	path := astar(board, Tile{X: Head.X, Y: Head.Y}, Tile{X: Goal.X, Y: Goal.Y})

	nextMove := path[len(path)-1]
	differenceX := nextMove.X - Head.X
	differenceY := nextMove.Y - Head.Y

	if differenceY == -1 {
		return directions[0]
	}

	if differenceY == 1 {
		return directions[1]
	}

	if differenceX == -1 {
		return directions[2]
	}

	if differenceX == 1 {
		return directions[3]
	}

	return directions[0]
}
