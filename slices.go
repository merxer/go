package main

import "fmt"

func main() {
	s := []int {1,2,3}

	fmt.Printf("%v\n", s[1:2])
	fmt.Printf("%v\n", s[1:])
	fmt.Printf("%v\n", s[2:])
	fmt.Printf("%v\n", s[:3])
}
