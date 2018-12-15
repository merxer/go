package main

import "fmt"

const (
	PI = 3.14
	Language = "go"
	A = iota
	B = iota
	C = iota
)

func main() {
	fmt.Println(A, B, C)
	fmt.Println(PI)
	fmt.Println(Language)
}
