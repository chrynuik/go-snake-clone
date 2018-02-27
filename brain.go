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
	Start := Point{X: Head.X, Y: Head.Y}
	Tail := Body[len(Body)-1]

	Health := Us.Health

	AllSnakes := data.Snakes

	closestFood := PriorityQueue{}
	closestEnemy := PriorityQueue{}

	for _, Morsel := range Food {
		newItem := &Item{
			priority: hueristic(Point{X: Head.X, Y: Head.Y}, Point{X: Morsel.X, Y: Morsel.Y}),
			point:    Point{X: Morsel.X, Y: Morsel.Y},
		}

		heap.Push(&closestFood, newItem)
	}

	nextFood := heap.Pop(&closestFood).(*Item)

	board := Graph{}
	board.create(data)

	enemyHeads := getEnemyHeads(AllSnakes, Us)
	attackableEnemies := getAttackableEnemies(enemyHeads, 6)

	fmt.Println("ALL ENEMIES", enemyHeads)
	fmt.Println("ATTACKABLE ENEMIES", attackableEnemies)

	fmt.Println(board)

	for _, Enemy := range enemyHeads {
		EnemyCoords := Enemy.Coords

		newItem := &Item{
			priority: hueristic(Point{X: Head.X, Y: Head.Y}, Point{X: EnemyCoords.X, Y: EnemyCoords.Y}),
			point:    Point{X: EnemyCoords.X, Y: EnemyCoords.Y},
		}

		heap.Push(&closestEnemy, newItem)
	}

	nextEnemy := heap.Pop(&closestFood).(*Item)

	Goal := Point{}

	if nextFood.priority > nextEnemy.priority {
		Goal = nextEnemy.point
	} else {
		Goal = nextFood.point
	}

	if Health < 99 && Health > 50 && hueristic(Start, Goal) > 6 {
		Goal = Point{X: Tail.X, Y: Tail.Y}
	}

	path := astar(board, Point{X: Head.X, Y: Head.Y}, Point{X: Goal.X, Y: Goal.Y})

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
