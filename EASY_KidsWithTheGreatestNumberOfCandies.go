package main

import "fmt"

func main() {
	var candies []int = []int{2, 3, 5, 1, 3}
	var extraCandies int = 3
	var result []bool
	var maxCountOfCandies int = candies[0]
	for i := 0; i < len(candies); i++ {
		if candies[i] > maxCountOfCandies {
			maxCountOfCandies = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= maxCountOfCandies {
			result = append(result, true)
			fmt.Println(candies[i]+extraCandies, "i = ", i)
		} else {
			result = append(result, false)
			fmt.Println(candies[i]+extraCandies, "i = ", i)
		}
	}
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}
