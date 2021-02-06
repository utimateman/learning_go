// Example of "switch" statement

package main

import(
	"fmt"
)

func main() {
	x := 7

	switch x {
	case 1:
		fmt.Println("one")
	
	case 2:
		fmt.Println("two")
	
	case 3:
		fmt.Println("three")
	
	case 4:
		fmt.Println("four")

	default:
		fmt.Println("many")
	}

	switch {
	case x > 100:
		fmt.Println("x is very big")
	case x > 10:
		fmt.Println("x is big")
	default:
		fmt.Println("x is small") 
	}
	
}