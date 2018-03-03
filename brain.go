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

	enemyHeads := getEnemyHeads(AllSnakes, Us)

	board := Graph{}
	board.create(data, enemyHeads)

	pathToTail := astar(board, Head, Tail)

	fmt.Println("Do we have a path to our tail", pathToTail)

	nextFood := getFoodPath(Food, Us, pathToTail)

	attackableEnemies := getAttackableEnemies(enemyHeads, Us.Length)

	// fmt.Println("ALL ENEMIES", enemyHeads)
	// fmt.Println("ATTACKABLE ENEMIES", attackableEnemies)

	fmt.Println(board)

	var isLongest bool

	for _, snake := range enemyHeads {
		if snake.Length < Us.Length {
			isLongest = true
		}
	}

	nextEnemy := getEnemyPath(attackableEnemies, Us)

	Goal := Point{}

	if nextFood != nil && nextEnemy != nil &&
		nextFood.priority > nextEnemy.priority {
		Goal = nextEnemy.point
	} else if nextFood != nil {
		if !isLongest || Health > 50 {
			Goal = nextFood.point
		}
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
		if len(pathToTail) > 0 {
			nextMove = pathToTail[len(pathToTail)-1]
		} else {
			nextMove = board.neighbors(Head)[0]
		}
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
