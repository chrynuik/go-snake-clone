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

	Move := ""

	X := Food.X - Head.X
	Y := Food.Y - Head.Y

	if X == 0 {
		if Y > 0 {
			Move = directions[1]
		} else if Y < 0 {
			Move = directions[0]
		}
	} else {
		if X > 0 {
			Move = directions[3]
		} else if X < 0 {
			Move = directions[2]
		}
	}

	board := Graph{}
	board.create(data)

	fmt.Println(board)
	fmt.Println(astar(board, Tile{X: Head.X, Y: Head.Y}, Tile{X: Food.X, Y: Food.Y}))
	return Move
}
