package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	FirstName, LastName string
	City, Gender string `json:",omitempty"`			// omitempty berfungsi sebagai optional value seoerti di typescript dengan tanda ?
	Age int `json:",omitempty"`									// tapi omitempty harus dibarengi dengan exported variable/key (diawali dengan huruf kapital)
}

// Greeting Method (value receiver); p = unique identifier similar to 'this' keyword
func (p Person) greet() string {
	return "Hello, my name is " + p.FirstName + " " + p.LastName + " and I am " + strconv.Itoa(p.Age) + " years old."
}

// Pointer Receiver
func (p *Person) setAge(newAge int) {
	p.Age = newAge
}

func main() {
	// Init perso using struct
	person1 := Person {
		FirstName: "Ary", 
		LastName: "Setiyawan",
		City: "Badung",
		Gender: "M",
		Age: 30,
	}

	// Alternative (tidak support json: omitempty)
	person2 := Person {
		"Ary",
		"Setiyawan",
		"Badung",
		"M",
		30,
	}

	person3 := Person {
		FirstName: "Dwija",
		LastName: "Shenda",
	}

	jPerson1, _ := json.Marshal(person1)
	jPerson3, _ := json.Marshal(person3)

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(string(jPerson1))
	fmt.Println(string(jPerson3))

	fmt.Println(person1.FirstName)
	person1.Age++
	person3.setAge(15)
	fmt.Println(person1)
	fmt.Println(person1.greet())
	fmt.Println(person3.greet())
}

