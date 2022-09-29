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
	for _, value := range myMap {
		fmt.Println(value)
	}

	// Slice of map
	students := []map[string]string{
		{"name": "Piter", "age": "24"},
		{"name": "Yogi", "age": "25"},
	}
	for _, val := range students{
		fmt.Println(val["name"])
	}
}