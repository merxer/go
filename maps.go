package main

import "fmt"

func main() {
	prefixMap := map[string]string {
		"pat" : "Mr. ",
		"kung" : "Mrs. ",
	}

	name := "kung"
	message := prefixMap[name] + " " + name
	fmt.Println(message)

}
