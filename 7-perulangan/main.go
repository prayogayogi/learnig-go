package main

import "fmt"

func main() {
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("Bismilah Belajar golang :", i)
	// }
	
	// i := 1
	// for i <= 20{
	// 	fmt.Println("Bismilah Belajar golang :", i)
	// 	i++
	// }

	title := "Bismilah Belajar golang"
	for _, letter := range title{
		fmt.Println("leeter: ", string(letter))
	}
}