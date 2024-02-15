package main

import (
	"fmt"
)

func main() {
	// Syntax 1
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(colors)
	// Syntax 2
	var colors2 map[string]string
	fmt.Println(colors2)

	// Syntax 3
	colors3 := make(map[string]string)
	colors3["white"] = "#ffffff"
	fmt.Println(colors3)

	delete(colors3, "white")
	fmt.Println(colors3)

	//TODO: Is there any difference between different syntax
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
	/**
	* Difference between Map and Struct
	* 1. in Map, all keys and values must have the same data type, while in struct, it can be different
	* 2. Struct is value type, while the Map is reference type
	* 3. Keys supports indexing in map, not in struct
	* 4. struct needs to know all the types during the compiling time
	 */
}
