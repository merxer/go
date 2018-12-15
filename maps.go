package main

import "fmt"

func main() {
	prefixMap := map[string]string {
		"pat" : "Mr. ",
		"kung" : "Mrs. ",
	}

	prefixMap["Odin"] = "Jr. "

	name := "Odin"
	message := prefixMap[name] + " " + name
	fmt.Println(message)

}
