package main

import "fmt"

// a = [2, 1, 3, 5, 3, 2]
func solution(a []int) int {
	m := make(map[int]int)
	for _, v := range a {
		_, ok := m[v]
		if ok {
			return v
		}
		m[v] = 1
	}
	return -1
}

func main() {
	input := []int{2, 1, 3, 5, 3, 2}
	fmt.Println(solution(input))
}
