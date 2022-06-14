package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName, LastName, City, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet()string {
	return "Hello, my name is " + p.firstName +" " + p.LastName + " and i am " + strconv.Itoa(p.age) + " years!"
}

//  hasBirthday method (pointer reciever)
func(p *Person) hasBirthday() {
	p.age++
} 

// getMarried (pointer reciever)
func(p *Person) getMarried(spouseLastName string){
	if p.gender == "m"{
		return
	}else {
		p.LastName = spouseLastName;
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Bob", LastName: "Smith", City: "Boston", gender: "m", age: 23}
	person2 := Person{firstName: "Lucy", LastName: "Smith", City: "Boston", gender: "f", age: 20}
	
	// Alternative
	// person2 := Person{"Bob", "Smith","Boston", "m", 23}
	
	// fmt.Println(person1)
	// fmt.Println(person1.firstName)

	// person1.age++
	// fmt.Println(person1.age)

	fmt.Println(person1.greet())

	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	
	fmt.Println(person1.age)

	person1.getMarried("Williams")
	fmt.Println(person1.LastName)
	
	person2.getMarried("Williams")
	fmt.Println(person2.LastName)
	fmt.Println(person2.greet())
	
	
}