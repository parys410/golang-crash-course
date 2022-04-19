package main

import "fmt"

func main() {
	x := 15
	y := 15

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal with %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	color := "green"
	switch(color) {
	case "blue":
		fmt.Println("Color is Blue")
	case "red":
		fmt.Println("Color is Red")
	default:
		fmt.Println("Color is Other")
	}
}

