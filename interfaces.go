package main

import "fmt"

type Salutation struct {
	Name string
	Greeting string
}

type Salutations []Salutation

type Printer func(Salutation, int)

type Renameable interface{
	Rename(newName string)
}

func (salutation *Salutation) Rename (newName string) {
	salutation.Name = newName
}

func (salutions Salutations)Greeting(do Printer,times int) {
	for _, value := range salutions {
		do(value,times)
	}
}

func RenameToFrog(r Renameable) {
	r.Rename("Frog")
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

	//friends[0].Rename("P@t")
	RenameToFrog(&friends[0])

	friends.Greeting(PrintLine, 2)
}
