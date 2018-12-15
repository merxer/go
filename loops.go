package main

import "fmt"

func printLoop(message string) {
	for {
	        fmt.Println(message)
	}

}

func main() {
	printLoop("Infinity Loops")
}
