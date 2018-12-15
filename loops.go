package main

import "fmt"

func printLoop(message string, times int) {
	i := 0
	for {
		if i >= times { break }
	        fmt.Println(message)
		i++
	}

}

func main() {
	printLoop("Infinity Loops", 5)
}
