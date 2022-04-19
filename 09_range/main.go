package main

import "fmt"

func main() {
	ids := []int{33,76,53,23,1}

	// Loop through ids
	for i,id := range ids {
		fmt.Printf("%d - %d\n",i,id)
	}

	// Not using index
	for _,id := range ids {
		fmt.Printf("ID : %d\n",id)
	}

	// Add ids together
	sum := 0
	for _,id := range ids {
		sum += id
	}

	fmt.Printf("Total is %d\n\n", sum)

	// Range with map
	emails := map[string]string{"Bob":"bob@gmail.com", "Sharon":"sharon@gmail.com", "Mike":"mike@gmail.com"}

	for k, v := range emails {
		fmt.Printf("Key: %s - Value: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}
}

