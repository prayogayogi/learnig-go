package main

import (
	"fmt"
	"hello-word/calculation"
)
func main() {
	fmt.Println("Hello Word")

	result := calculation.Add(4, 5)
	fmt.Println(result)
}