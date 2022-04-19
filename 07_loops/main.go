package main

import "fmt"

func main() {
	// Long Method
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Short Method
	for j := 1; j <= 10; j++ {
		fmt.Printf("Number %d\n", j)
	}

	// FizzBuzz
	fmt.Printf("\n\n")
	for j := 1; j <= 100; j++ {
		if j % 3 == 0 && j % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if j % 3 == 0 {
			fmt.Println("Fizz")
		} else if j % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(j)
		}
	}
}

