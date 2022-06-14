package main

import "fmt"

func main() {
	// // Define map
	// emails := make(map[string]string)

	// // Assign kv

	// emails["Bob"] = "bob@abv.bg"
	// emails["Loly"] = "Loly@abv.bg"

	// fmt.Println(emails)

	// // Delete from map
	// delete(emails, "Bob")
	// fmt.Println(emails)

	emails := map[string]string{"Bob":"bob@abv.bg","Sharon":"sharon@gmail.com"}
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
}