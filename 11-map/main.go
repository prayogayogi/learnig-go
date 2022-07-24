package main

import "fmt"

func main() {
	// Bikin map dengan menginisialisai varibel nya dulu
	// var myMap map[string]int
	// myMap = map[string]int{}
	// myMap["golang"] = 9
	// myMap["php"] = 8
	// myMap["javascript"] = 9

	// Bikin map sekaligus diisi isinya
	myMap := map[string]string{
		"php": "is intermediate", 
		"go": "super fast",
	}
	fmt.Println(myMap)
}