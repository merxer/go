package main

import "fmt"

func main() {
	s := []int {1,2,3}
	s1 := []int {5, 6}

	fmt.Printf("%v\n", s[1:2])
	fmt.Printf("%v\n", s[1:])
	fmt.Printf("%v\n", s[2:])
	fmt.Printf("%v\n", s[:3])

	fmt.Println(s)
	s = append(s, 4)
	fmt.Println(s)
	s = append(s, s1...)
	fmt.Println(s)
}
