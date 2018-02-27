package main

// EnemyHead is a struct
type EnemyHead struct {
	Coords Point
	Length int
}

func getEnemyHeads(allSnakes SnakeList, us Snake) []EnemyHead {
	var heads []EnemyHead

	for _, snake := range allSnakes {
		if snake.ID != us.ID {
			heads = append(heads, EnemyHead{
				Coords: snake.Body[0],
				Length: snake.Length,
			})
		}
	}
	return heads
}

func getAttackableEnemies(snakes []EnemyHead, ourLength int) []Point {
	var weakEnemies []Point

	for _, snake := range snakes {
		if snake.Length < ourLength {
			weakEnemies = append(weakEnemies, Point{
				X: snake.Coords.X,
				Y: snake.Coords.Y,
			})
		}
	}
	return weakEnemies
}
