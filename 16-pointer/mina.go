package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main(){

	address := Address{"Subang", "Jawabarat", ""}
	ChangeAddressToIndonesia(&address)
	fmt.Println(address)

	number1 := 5
	number2 := &number1

	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(*number2)
	
	fmt.Println("--------------------------------")
	
	*number2 = 8
	fmt.Println(number2)
	fmt.Println("Number2: ", *number2)
	fmt.Println("Number1: ",  number1)
	
	fmt.Println("-------------Pointer var-------------------")
	
	var numberA int = 10
	var numberB *int = &numberA
	
	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)

	fmt.Println("--------------------------------")
	
	numberA = 20
	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)

	fmt.Println("--------------------------------")
	
	number := 5
	fmt.Println("Nilai awal:", number)
	change(&number, 100)
	fmt.Println("Nilai akhir:", number)
}

func change(old *int, new int){
	fmt.Println("Di dalam function:", *old)
	*old = new
	fmt.Println("Di dalam function:", *old)
}