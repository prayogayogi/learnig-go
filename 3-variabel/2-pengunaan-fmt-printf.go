package main

import "fmt"

func main(){
	var firstName string = "Yogik"

	var lastName string
	lastName = "Piter"

	fmt.Printf("Hallo Yogik Piter \n")
	fmt.Printf("Hallo %s %s!\n", firstName, lastName)
	fmt.Println("Hallo", firstName, lastName)
}