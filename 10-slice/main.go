package main

import "fmt"

func main(){
	// Tanpa iteration atau langsung
	gamingConsole := []string{"playStation","Cod"}
	gamingConsole2 := []string{"Gta", "PES"}

	result := append(gamingConsole, gamingConsole2...)

	// Dengan mengunakan append
	// var gamingConsole []string
	// gamingConsole = append(gamingConsole, "PlayStation4")
	// gamingConsole = append(gamingConsole, "Cod")
	// gamingConsole = append(gamingConsole, "Point Blank")

	for _, console := range result{
		fmt.Println("Game:", console)
	}
}