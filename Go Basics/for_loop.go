// for loop

package main

import(
	"fmt"
)

func main() {
	fmt.Println("--- forloop ---")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("--- break ---")

	for i := 0; i < 3; i++ {
		if i > 1 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("--- continue ---")

	for i := 0; i < 3; i++ {
		if i < 1 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("--- while ---")

	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	fmt.Println("--- while 1 ---")

	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}
	
} 