package main

import "fmt"

func printLoop(message string,loop int) {
	for i := 0; i< loop; i++ {
	        fmt.Println(message)
	}

}

func main() {
	printLoop("Hello", 10)
}
