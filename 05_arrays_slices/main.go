package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	// Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and Assign
	fruitArr := [5]string{"Apple", "Orange"}
	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitArr, fruitSlice)
	fmt.Println(len(fruitArr), len(fruitSlice))
	fmt.Println(cap(fruitArr), cap(fruitSlice))
	fmt.Println(fruitSlice[3:4])
	fmt.Println()
	fmt.Println(fruitSlice[0:], len(fruitSlice[0:]), cap(fruitSlice[0:]))
	fmt.Println(fruitSlice[1:], len(fruitSlice[1:]), cap(fruitSlice[1:]))
	fmt.Println(fruitSlice[2:], len(fruitSlice[2:]), cap(fruitSlice[2:]))
	fmt.Println(fruitSlice[3:], len(fruitSlice[3:]), cap(fruitSlice[3:]))
	fmt.Println(fruitSlice[4:], len(fruitSlice[4:]), cap(fruitSlice[4:]))
	fmt.Println()
	fmt.Println(fruitSlice[:0], len(fruitSlice[:0]), cap(fruitSlice[:0]))
	fmt.Println(fruitSlice[:1], len(fruitSlice[:1]), cap(fruitSlice[:1]))
	fmt.Println(fruitSlice[:2], len(fruitSlice[:2]), cap(fruitSlice[:2]))
	fmt.Println(fruitSlice[:3], len(fruitSlice[:3]), cap(fruitSlice[:3]))
	fmt.Println(fruitSlice[:4], len(fruitSlice[:4]), cap(fruitSlice[:4]))
	fmt.Println()
	fmt.Println(fruitSlice[0:0], len(fruitSlice[0:0]), cap(fruitSlice[0:0]))
	fmt.Println(fruitSlice[0:1], len(fruitSlice[0:1]), cap(fruitSlice[0:1]))
	fmt.Println(fruitSlice[0:2], len(fruitSlice[0:2]), cap(fruitSlice[0:2]))
	fmt.Println(fruitSlice[0:3], len(fruitSlice[0:3]), cap(fruitSlice[0:3]))
	fmt.Println(fruitSlice[0:4], len(fruitSlice[0:4]), cap(fruitSlice[0:4]))
	fmt.Println()
	fmt.Println(fruitSlice[1:1], len(fruitSlice[1:1]), cap(fruitSlice[1:1]))
	fmt.Println(fruitSlice[1:2], len(fruitSlice[1:2]), cap(fruitSlice[1:2]))
	fmt.Println(fruitSlice[1:3], len(fruitSlice[1:3]), cap(fruitSlice[1:3]))
	fmt.Println(fruitSlice[1:4], len(fruitSlice[1:4]), cap(fruitSlice[1:4]))
	fmt.Println()
	fmt.Println(fruitSlice[2:2], len(fruitSlice[2:2]), cap(fruitSlice[2:2]))
	fmt.Println(fruitSlice[2:3], len(fruitSlice[2:3]), cap(fruitSlice[2:3]))
	fmt.Println(fruitSlice[2:4], len(fruitSlice[2:4]), cap(fruitSlice[2:4]))
	fmt.Println()
	fmt.Println(fruitSlice[3:3], len(fruitSlice[3:3]), cap(fruitSlice[3:3]))
	fmt.Println(fruitSlice[3:4], len(fruitSlice[3:4]), cap(fruitSlice[3:4]))
	fmt.Println()
	fmt.Println(fruitSlice[4:4], len(fruitSlice[4:4]), cap(fruitSlice[4:4]))

	/* Penting! Cara Baca Slice */
	/* nama_slice[dipotong_berapa:tampil_berapa = panjang - 1] */
	/*
	* fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
	*	fruitSlice[1:2]
	* 1: artinya fruitSlice dipotong 1 dari index 0 maka sisanya ["Orange", "Grape", "Cherry"]
	* :2 artinya tampilkan data sebanyak 2 - 1 = 1 dari ["Orange", "Grape", "Cherry"] maka yang tampil "Orange"  
	* len(fruitSlice[1:2]); Panjang / Banyaknya data = 1
	* cap(fruitSlice[1:2]); Kapasitas data = total_kapasitas (yaitu 4) - 1: (yaitu 1) maka 3
	*/

}
