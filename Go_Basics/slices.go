// slices
package main

import (
	"fmt"
)

func main() {
	sweets := []string{"cake", "ice cream","waffle"}
	fmt.Println("sweets = %v (type %T)\n", sweets, sweets)

	fmt.Println(len(sweets))

	fmt.Println(sweets[1])

	fmt.Println(sweets[1:])

	fmt.Println("--- access normal for --- ")
	for i := 0; i < len(sweets); i++ {
		fmt.Println(sweets[i])
	}

	fmt.Println("--- single value range --- ")
	for i := range sweets {
		fmt.Println(i)
	}

	fmt.Println("--- double value --- ")
	for i, name := range sweets {
			fmt.Printf("%s at %d\n", name, i)
	}

	fmt.Println("--- double value ignore  index --- ")
	for _, name := range sweets{
		fmt.Println(name)
	}

	fmt.Println("--- append ---")
	sweets = append(sweets, "candy")
	fmt.Println(sweets)
}