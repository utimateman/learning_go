package main

import (
	"fmt"
	//"os" //ver.2.0
)

func safeValue(vals []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	}()

	return vals[index]
}

func main() {
	// ver.1.0
	/*
	vals := []int{1,2,3}
	// this will cause panic
	v := vals[10]
	fmt.Println(v)
	*/

	// ver.2.0
	/*
	file, err := os.Open("no-such-file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("file opened")
	*/

	// ver.3.0 to wrap untrusted code
	v := safeValue([]int{1,2,3}, 10)
	fmt.Println(v)
}