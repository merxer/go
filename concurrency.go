package main

func main() {
	go func() {
		println("Hello")
	}()
	go func() {
		println("Go")
	}()
}
