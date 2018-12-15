package main

import "fmt"

func main() {
	s := []int {1,2,3}

	fmt.Printf("%v\n", s[1:2])
	fmt.Printf("%v\n", s[1:])
	fmt.Printf("%v\n", s[2:])
	fmt.Printf("%v\n", s[:3])

	fmt.Println(s)
	s1 := append(s, 4)
	fmt.Println(s1)
}
