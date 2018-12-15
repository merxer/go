package main

import "fmt"

type Salutation struct {
	name string
	greeting string
}

func Greet(salutation Salutation) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting)
	fmt.Println(message)
	fmt.Println(alternate)
}

func CreateMessage(name, greeting string) (string, string) {
	return greeting + " " + name, "Hey! " + name
}


func main() {
	var s = Salutation{"Bob", "Hello"}
	Greet(s)
}
