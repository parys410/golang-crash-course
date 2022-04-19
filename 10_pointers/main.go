package main

import "fmt"

func main() {
	// Pointer basicly allows you point to memory address of location of the value
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n",a)
	fmt.Printf("%T\n",b)

	// Use * to read val from memory address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with Pointer
	*b = 10
	fmt.Println(a)
}

