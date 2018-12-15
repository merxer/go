package main

import (
	"time"
)

func main() {
	go func() {
		for i := 0; i < 100; i++ {
			println(i, "Hello")
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			println(i, "Go")
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
