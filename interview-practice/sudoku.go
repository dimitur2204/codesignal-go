package main

import "fmt"

func checkBox(startRow int, startCol int, matrix [][]string) bool {
	boxEntries := make(map[string]string)
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			_, ok := boxEntries[matrix[i][j]]
			if ok {
				return true
			}
			if matrix[i][j] != "." {
				boxEntries[matrix[i][j]] = ""
			}
		}
	}
	return false
}

func solution(matrix [][]string) bool {

	for row := 0; row < len(matrix); row++ {
		rowEntries := make(map[string]string)
		colEntries := make(map[string]string)
		for col := 0; col < len(matrix); col++ {
			isStart := col%3 == 0 && row%3 == 0
			if isStart && checkBox(row, col, matrix) {
				return false
			}
			_, okCol := colEntries[matrix[col][row]]
			if okCol {
				return false
			}
			if matrix[col][row] != "." {
				colEntries[matrix[col][row]] = ""
			}
			_, okRow := rowEntries[matrix[row][col]]
			if okRow {
				return false
			}
			if matrix[row][col] != "." {
				rowEntries[matrix[row][col]] = ""
			}
		}

	}
	return true
}

func main() {

	input := [][]string{
		{"1", ".", ".", "8", "4", ".", ".", "2", "."},
		{".", "9", "6", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "1", ".", ".", ".", ".", ".", "."},
		{".", "6", "7", ".", ".", ".", ".", ".", "9"},
		{".", ".", ".", ".", ".", ".", "8", "1", "."},
		{".", "3", ".", ".", ".", ".", ".", ".", "6"},
		{".", ".", ".", ".", ".", "7", ".", ".", "."},
		{".", ".", ".", "5", ".", ".", ".", "7", "."},
	}
	// input2 := [][]string{
	// 	{".", ".", ".", ".", "2", ".", ".", "9", "."},
	// 	{".", ".", ".", ".", "6", ".", ".", ".", "."},
	// 	{"7", "1", ".", ".", "7", "5", ".", ".", "."},
	// 	{".", "7", ".", ".", ".", ".", ".", ".", "."},
	// 	{".", ".", ".", ".", "8", "3", ".", ".", "."},
	// 	{".", ".", "8", ".", ".", "7", ".", "6", "."},
	// 	{".", ".", ".", ".", ".", "2", ".", ".", "."},
	// 	{".", "1", ".", "2", ".", ".", ".", ".", "."},
	// 	{".", "2", ".", ".", "3", ".", ".", ".", "."},
	// }
	fmt.Println(fmt.Sprint(solution(input)))
}
