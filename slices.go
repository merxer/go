package main

import "fmt"

func main() {
	var s []int
	s = make([]int, 3)
	s[0] = 1
	s[1] = 2
	s[2] = 3

	fmt.Printf("%v\n", s)
}
