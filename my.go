package main

import "fmt"

func removeIndexFromArr(arr *[]Position, index int) {
	*arr = append((*arr)[:index], (*arr)[index+1:]...)
}

type Position struct {
	x int
	y int
}

func main() {
	positions := make([]Position, 9)

	for i := 0; i < 9; i++ {
		positions[i].x = i % 3
		positions[i].y = i / 3
	}

	fmt.Println(positions)

	removeIndexFromArr(&positions, 0)

	fmt.Println(positions)
}
