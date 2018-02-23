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

	Food := data.Food[0]
	Snake := data.You
	Body := Snake.Body
	Head := Body[0]

	board := Graph{}
	board.create(data)

	fmt.Println(board)
	path := astar(board, Tile{X: Head.X, Y: Head.Y}, Tile{X: Food.X, Y: Food.Y})

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
