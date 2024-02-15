package main

import (
	"fmt"
)

func main() {
	// Syntax 1
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors)
	// Syntax 2
	var colors2 map[string]string
	fmt.Println(colors2)

	// Syntax 3
	colors3 := make(map[string]string)
	colors3["white"] = "#ffffff"
	fmt.Println(colors3)

	//TODO: Is there any difference between different syntax
}
