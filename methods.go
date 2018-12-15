package main

import "fmt"

type Salutation struct {
	Name string
	Greeting string
}

type Salutations []Salutation

type Printer func(Salutation, int)

func (salutions Salutations)Greeting(do Printer,times int) {
	for _, value := range salutions {
		do(value,times)
	}
}

func PrintLine(s Salutation, times int) {
	for  i := 0; i< times; i++ {
		fmt.Println(s.Greeting, s.Name)
	}
}

func main() {
	friends := Salutations {
		{"Pat", "Hello"},
		{"Kung", "Hi"},
	}

	friends.Greeting(PrintLine, 4)
}
