package main

import (
	"fmt"
)

func main() {
	fmt.Println(chessboardGame(5, 2))
}

func chessboardGame(x int32, y int32) string {
	x = x % 4
	y = y % 4
	if (y == 0) || (y == 3) || (x == 0) || (x == 3) {
		return "First"
	}
	return "Second"
}

func chessboardGame2(x int32, y int32) string {
	player := 1
	allowMove := true

	for allowMove {
		x, y, allowMove = move(x, y)
		if player == 1 {
			player = 2
		} else {
			player = 1
		}
	}
	if player == 1 {
		return "First"
	}
	return "Second"
}

func move(x int32, y int32) (int32, int32, bool) {
	if x-2 >= 0 && y+1 <= 15 {
		return x - 2, y + 1, true
	} else if x-2 >= 0 && y-1 >= 0 {
		return x - 2, y - 1, true
	} else if x+1 <= 15 && y-2 >= 0 {
		return x + 1, y - 2, true
	} else if x-1 >= 0 && y-2 >= 0 {
		return x - 1, y - 2, true
	}

	return 0, 0, false
}
