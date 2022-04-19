package main

import "fmt"

func main() {
	// Maps is Key Value Pair
	// emails := make(map[string]string)

	// Assign key Value
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "mike@gmail.com"

	// Declare map and add key value
	emails := map[string]string{"Bob":"bob@gmail.com", "Sharon":"sharon@gmail.com", "Mike":"mike@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Sharon"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}

