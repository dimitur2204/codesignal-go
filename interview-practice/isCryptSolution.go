package main

import (
	"fmt"
	"strconv"
)

func decrypt(word string, solution [][]string) string {
	arr := ""
	for _, letter := range word {
		for _, solutionTuple := range solution {
			if string(letter) == solutionTuple[0] {
				arr += solutionTuple[1]
			}
		}
	}
	return arr
}

func solution(crypt []string, solution [][]string) bool {
	sumArr := make([]int, 0)

	for _, word := range crypt {
		decryptedWord := decrypt(word, solution)
		if len(decryptedWord) > 1 && decryptedWord[0] == '0' {
			return false
		}
		num, _ := strconv.Atoi(string(decryptedWord))
		sumArr = append(sumArr, num)
	}
	return sumArr[0]+sumArr[1] == sumArr[2]
}

func main() {
	solutionArr := [][]string{{"O", "0"},
		{"M", "1"},
		{"Y", "2"},
		{"E", "5"},
		{"N", "6"},
		{"D", "7"},
		{"R", "8"},
		{"S", "9"}}
	crypt := []string{"SEND", "MORE", "MONEY"}
	fmt.Println(solution(crypt, solutionArr))
}
