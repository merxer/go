package main

import (
	"time"
)

func main() {
	godur, _ := time.ParseDuration("10ms")
	go func() {
		for i := 0; i < 100; i++ {
			println(i, "Hello")
			time.Sleep(godur)
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			println(i, "Go")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
