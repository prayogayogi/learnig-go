package main

import (
	"fmt"
)

func main(){
	// Array yang menalokasikan element
	// var languages [5]string
	// languages[0] = "Go"
	// languages[1] = "javaScript"
	// languages[2] = "php"
	// languages[3] = "Ruby"
	// languages[4] = "Python"

	// Inisialisasi Nilai Awal Array
	languages := [...]string{
		"Go", 
		"JavaScript", 
		"Php", 
		"Ruby", 
		"Python",
		"Kotlin",
	}
	for index, leng := range languages{
		fmt.Println("index: ", index, "language: ", leng)
	}
}
