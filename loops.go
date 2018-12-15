package main

import "fmt"

type LoopMessage struct {
	message string
	times int
}

func printLoop(message string, times int) {
	i := 0
	for {
		if i >= times { break }
	        fmt.Println(message)
		i++
	}

}

func main() {
	slice := []LoopMessage {
		{"Pat", 10},
		{"Kung", 7},
	}
	for _, v := range slice {
		printLoop(v.message, v.times)
}
}
