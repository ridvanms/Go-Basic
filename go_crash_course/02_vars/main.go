package main

import "fmt"

func main(){

	// Using var
	// var name = "John"
	var age = 23
	const isCool = true;
	var size float32 = 2.4

	// Shorthand
	// name := "Brad"
	// email := "brad@abv.bg"

	name, email := "Brad","brad@abv.bg"

	fmt.Println(name,age,isCool,email)
	fmt.Printf("%T\n",size) // printing the type of age

	
}
