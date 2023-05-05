package main

import "fmt"

func main() {
	s := " sdf "
	var result, lastWordLength int
	for i := 0; i < len(s); i++ {
		if string(s[i]) != " " {
			lastWordLength++
		} else if lastWordLength != 0 {
			result = lastWordLength
			lastWordLength = 0
		}
	}
	if lastWordLength != 0 {
		result = lastWordLength
	}
	fmt.Println("Length of last word 's' is", result)
}
