package main

import "fmt"

func main() {
	// result := printMyResult("Piter")
	// fmt.Println(result)

	// result := add(10, 20)
	// fmt.Println(result)

	luas, keliling:= calculate(10, 2)
	fmt.Println(luas)
	fmt.Println(keliling)
}

func calculate(panjang, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return
}

// func add(number1, number2 int) int {
// 	return number1 + number2
// }

// func printMyResult(sentence string) string {
// 	result := sentence + " sedang belajar Go"
// 	return result
// }

// Function itu pasti ada
// 1. Input
// 2. Proses
// 3. Output