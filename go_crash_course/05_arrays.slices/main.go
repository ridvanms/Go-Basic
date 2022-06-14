package main

import "fmt"

func main() {

	// Declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple","Orange","Grape"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])

}