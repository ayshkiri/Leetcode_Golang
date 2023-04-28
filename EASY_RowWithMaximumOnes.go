package main

import "fmt"

func main() {
	mat := [][]int{{0, 1, 0}, {1, 1, 0}}
	result := []int{0, 0}
	maxOnes := 0
	for i := 0; i < len(mat); i++ {
		maxOnes = 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				maxOnes++
			}
		}
		if maxOnes > result[1] {
			result = []int{i, maxOnes}
		}
	}
	fmt.Println("Row = ", result[0], ", max count of ones = ", result[1])
}
