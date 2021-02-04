// parameter passing
package main

import (
	"fmt"
)

func doubleAt(values []int, i int) {
	values[i] *= 2
}

func double(n int) {
	n *= 2 
}

func doublePtr(n *int) {
	*n *= 2
}

func main() {
	values := []int{1,2,3,4} // pass by reference
	doubleAt(values, 2)
	fmt.Println(values)

	val := 10
	double(val) // pass by copy value
	fmt.Println(val)

	doublePtr(&val)
	fmt.Println(val)
}