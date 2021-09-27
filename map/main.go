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
	}

	// fruits[1] = "Melon"
	// fmt.Println(fruits)
	
	colors["white"] = "#fffff"

	fmt.Println(colors)
	fmt.Println(colors["white"])

	// delete(colors, "white")
	// fmt.Println(colors)
	
}
