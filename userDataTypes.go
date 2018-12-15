package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func main() {
	var s = Salutation{greeting: "Hello!",
		name: "Joe"}

	fmt.Println(s.name)
	fmt.Println(s.greeting)
}
