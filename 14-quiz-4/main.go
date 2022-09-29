package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total)

	result, err := calculate(10, 2, "=")
	if err != nil{
		fmt.Println("Terjadi kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

// Function sum
func sum(arr1 []int) int {
	var result int
	for _, val := range arr1 {
		result += val
	}
	return result
}

// Function calculate
func calculate(num1, num2 int, simbol string) (int, error) {
	var result int
	var errorResult error
	if simbol == "+" {
		result = num1 + num2
	} else if simbol == "-" {
		result = num1 - num2
	} else if simbol == "*" {
		result = num1 * num2
	} else if simbol == "/" {
		result = num1 / num2
	}else{
		errorResult = errors.New("Uknow Operation")
	}
	return result, errorResult
}
