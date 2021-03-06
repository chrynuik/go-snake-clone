package main

import (
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

	fmt.Println("Body length:", len(Body))

	board := Graph{}
	board.create(data)

	pathToTail := astar(board, Head, Tail)

	fmt.Println("Do we have a path to our tail", pathToTail)

	nextFood := getFoodPath(Food, Us, pathToTail)

	enemyHeads := getEnemyHeads(AllSnakes, Us)
	attackableEnemies := getAttackableEnemies(enemyHeads, 6)

	fmt.Println("ALL ENEMIES", enemyHeads)
	fmt.Println("ATTACKABLE ENEMIES", attackableEnemies)

	fmt.Println(board)

	nextEnemy := getEnemyPath(attackableEnemies, Us)

	Goal := Point{}

	if nextFood != nil && nextEnemy != nil &&
		nextFood.priority > nextEnemy.priority {
		Goal = nextEnemy.point
	} else if nextFood != nil {
		Goal = nextFood.point
	}

	if Health < 99 && Health > 50 && hueristic(Head, Goal) > 4 {
		Goal = Tail
	}

	path := astar(board, Head, Goal)

	fmt.Println("PATH", len(path))

	var nextMove Point

	fmt.Println("Got to the neighbors")

	if len(path) > 0 {
		fmt.Println("There is food", path)
		nextMove = path[len(path)-1]
	} else {
		fmt.Println("No food but neighbours")
		nextMove = board.neighbors(Head)[0]
	}

	fmt.Println("NextMove", nextMove)

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

	fmt.Println("Got to the end")

	return directions[0]
}
