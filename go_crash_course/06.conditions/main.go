package main

import "fmt"

func main() {

	x := 10
	y := 10

	if x<=y {
		fmt.Printf("%d is less than or equal to  %d\n",x,y)
		}else {
			fmt.Printf("%d is less than %d\n",y,x)
	}


	color := "red"
	
	if color == "red"{
		fmt.Printf("color is red")
	}else if color =="blue"{
		fmt.Printf("color is blue")
	}else {
		fmt.Printf("color is not blue or red")
	}

	// Switch
	switch color{
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is NOT red or blue")
	}
}