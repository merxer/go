package main

import "fmt"

const (
	PI = 3.14
	Language = "go"
)

const (
	A = iota
	B
	C
)

func main() {
	fmt.Println(A, B, C)
	fmt.Println(PI)
	fmt.Println(Language)
}
