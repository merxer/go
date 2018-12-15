package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	if prefix := GetPrefix(salutation.Name); isFormal {
		do(prefix + " " + message)
	} else {
		do(alternate)
	}
}

func GetPrefix(name string) (prefix string) {
	switch {
	case name == "Bob":
		prefix = "Mr "
	case name == "Joe",name == "Amy", len(name) == 10:
		prefix = "Dr "
	default:
		prefix = "Dude "
	}
	return
}

func TypeSwitchTest(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Salutation:
		fmt.Println("salutation")
	default:
		fmt.Println("unknown")
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
