package main

import "fmt"

func main() {
	var prefixMap map[string]string
	prefixMap = make(map[string]string)
	prefixMap["pat"] = "Mr "
	prefixMap["kung"] = "Mrs "

	name := "kung"
	message := prefixMap[name] + " " + name
	fmt.Println(message)

}
