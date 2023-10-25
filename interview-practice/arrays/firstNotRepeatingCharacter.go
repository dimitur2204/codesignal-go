package main

import "fmt"

// a = [2, 1, 3, 5, 3, 2]
func solution(s string) string {
	m := make(map[byte]int)
	for i := len(s) - 1; i >= 0; i-- {
		_, ok := m[s[i]]
		if ok {
			m[s[i]]++
			continue
		}
		m[s[i]] = 1
	}

	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return string(s[i])
		}
	}
	return string('_')
}

func main() {
	input := "xaacdabad"
	fmt.Println(solution(input))
}
