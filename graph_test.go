package main

import "testing"

var board = Graph{
	Width:  20,
	Height: 20,
	Grid:   make([][]int, 20),
}

func TestIsTileAccessible(t *testing.T) {
	tilesToTest := []Tile{{X: 1, Y: 1}, {X: 5, Y: 5}, {X: -1, Y: 0}, {X: 10, Y: 10}}
	expectedResults := []bool{true, true, false, true}

	for i, tile := range tilesToTest {
		result := board.isTileAccessible(tile)
		expected := expectedResults[i]

		if result != expected {
			t.Errorf("Expected:\n%#v\nGot:\n%#v", expected, result)
		}
	}
}
