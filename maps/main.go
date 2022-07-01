package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"white": "#fff",
	// 	"black": "#000000",
	// }
	// fmt.Println(colors)

	// colors := make(map[string]string)
	// colors["white"] = "#fff"
	// colors["red"] = "#ff0000"
	// delete(colors, "red")
	// fmt.Println(colors)

	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#fff",
		"black": "#000000",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
