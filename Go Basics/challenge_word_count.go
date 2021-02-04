// challenge: word count
package main

import (
	"fmt";
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`

	/*
	word_count_map := map[string]int{}
	
	// tokenize
	word_list := strings.Fields(text)
	fmt.Println(word_list)

	// add into maps
	for _, word := range word_list {
		//fmt.Println(word)
		temp_word := strings.ToLower(word)
		value, ok := word_count_map[temp_word]
		if !ok {
			word_count_map[temp_word] = 1
		} else {
			word_count_map[temp_word] = value + 1
		}
	}
	*/

	// shorter version

	words := strings.Fields(text)
	counts := map[string]int{}
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	for key, value := range counts {
		fmt.Println(key, "->", value)
	}


}