package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "sadbutsad"
	needle := "but"
	result := strings.Index(haystack, needle)
	fmt.Println(result)
	// var check int
	// for i := 0; i < len(haystack); i++ {
	// 	if haystack[i] == needle[0] {
	// 		check = 0
	// 		for j := 1; j < len(needle); j++ {
	// 			if haystack[i+j] == needle[j] {
	// 				check++
	// 			}
	// 		}
	// 		if check+1 == len(needle) {
	// 			fmt.Println(i)
	// 			break
	// 		}
	// 	}
	// }
	// fmt.Println(-1)
}
