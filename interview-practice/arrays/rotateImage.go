package main

import "fmt"

func solution(matrix [][]int) [][]int {
	newMatrix := make([][]int, 0)

	for i, a := range matrix {
		newRow := make([]int, 0)
		for j := range a {
			newRow = append([]int{matrix[j][i]}, newRow...)
		}
		newMatrix = append(newMatrix, newRow)
	}
	return newMatrix
}

func main() {
	input := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(fmt.Sprint(solution(input)))
}
