package main

import "fmt"

type Salutation struct {
	name string
	greeting string
}

func Greet(salutation Salutation) {
	_ , alternate := CreateMessage(salutation.name, salutation.greeting)
	fmt.Println(alternate)
}

func CreateMessage(name, greeting string) (string, string) {
	return greeting + " " + name, "Hey! " + name
}


func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s)
}
