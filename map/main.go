package main

import (
	"fmt"
)

func main() {

	// var colors map[string]string

	// colors := make(map[string]string)
	// fruits := make(map[int]string)

	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "ffffff",
	}

	// fruits[1] = "Melon"
	// fmt.Println(fruits)
	
	colors["white"] = "#fffff"

	// fmt.Println(colors)
	// fmt.Println(colors["white"])

	printMap(colors)

	// delete(colors, "white")
	// fmt.Println(colors)
	
}

func printMap(c map[string]string) {

	for color, hex := range c {
		// fmt.Println(color)
		// fmt.Println(hex)
		fmt.Println("Hex code for ", color , "is", hex)
	}

}
