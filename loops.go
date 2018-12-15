package main

import "fmt"

func printLoop(message string,loop int) {
	i := 0
	for i< loop {
	        fmt.Println(message)
		i++
	}

}

func main() {
	printLoop("Hello", 10)
}
