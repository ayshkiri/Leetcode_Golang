package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)
	fmt.Print(string(s[0]))
	if len(s) != len(t) {
		//return false
		fmt.Println("false")
	}
	for len(s) > 0 {
		t = strings.Replace(t, string(s[0]), "", 1)
		s = strings.Replace(s, string(s[0]), "", 1)
	}
	if len(s) == len(t) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
