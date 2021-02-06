// go string
package main

import (
	"fmt"
)

func main() {
	book := "The colour of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// uint8 is a byte

	// Strings in go are immutable
	// can't assign book[0] = 116

	// Slice (start, end), 0 based, 1/2 empty range
	fmt.Println(book[4:11])

	// Slice (no end)
	fmt.Println(book[4:])

	// slice (no start)
	fmt.Println(book[:4])
	
	//use + to concatenate strings
	fmt.Println("t" + book[1:])

	// unicode
	fmt.Println("It was 1/2 price!")

	// multi line
	poem := `
	The road goes ever on
	Down from teh door where it began
	...
	`

	fmt.Println(poem)
}