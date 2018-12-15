package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func main() {
	var s = Salutation{}
	s.name = "Joe"
	s.greeting = "Hi"

	fmt.Println(s.name)
	fmt.Println(s.greeting)
}
