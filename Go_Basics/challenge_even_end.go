// challenge: even-end
package main

import (
	"fmt"
)

func main() {
	count_even_end := 0
	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			result := i * j
			s := fmt.Sprintf("%d", result)
			if s[0] == s[len(s) - 1] {
				count_even_end = count_even_end + 1
			}
		}
	}

	fmt.Printf("Total even-ended: %d\n", count_even_end)
}