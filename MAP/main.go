package main

import "fmt"

func main() {

	mapDeclarationAndBasicMethods()

}

func mapDeclarationAndBasicMethods() {
	// Types of declaring map in go

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	// the output of this var println will be map[] as it will take zero values means empty map as map doesn't have any key value pair as we have seen this earlier
	// var colors map[string]string

	// colors := make(map[string]string)

	// This is the syntax to add the key value pair into existing map
	colors["white"] = "#FFFFFF"

	printMap(colors)
	//This is an existing function of go that takes a map and key as an argument and delete that key from the map
	delete(colors, "white")

	fmt.Println(colors)
}

func printMap(col map[string]string) {

	for color, hexCode := range col {
		fmt.Println("Hex code for ", color, "is", hexCode)
	}
}
