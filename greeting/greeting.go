package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	if prefix := "Mr. "; isFormal {
		do(prefix + " " + message)
	} else {
		do(alternate)
	}
}

func CreateMessage(name string, greeting string) (message string, alternate string) {
	fmt.Println(len(greeting))
	message = greeting + " " + name
	alternate = "Hey! " + name
	return message, alternate
}

func Print(s string) {
	fmt.Print(s)
}
func PrintLine(s string) {
	fmt.Println(s)
}

func CreateCustomFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
