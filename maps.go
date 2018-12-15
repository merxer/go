package main

import "fmt"

func main() {
	prefixMap := map[string]string {
		"pat" : "Mr. ",
		"kung" : "Mrs. ",
	}

	prefixMap["Odin"] = "Jr. "

	delete(prefixMap, "pat")

	name := "pat"
	message := prefixMap[name] + " " + name
	fmt.Println(message)

}
