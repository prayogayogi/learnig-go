package main

import "fmt"

func main(){
	// Tanpa iteration atau langsung
	// gamingConsole := []string{"playStation","Cod"}

	// Dengan mengunakan append
	var gamingConsole []string
	gamingConsole = append(gamingConsole, "PlayStation4")
	gamingConsole = append(gamingConsole, "Cod")
	gamingConsole = append(gamingConsole, "Point Blank")

	for _, console := range gamingConsole{
		fmt.Println("Game: ", console)
	}
}