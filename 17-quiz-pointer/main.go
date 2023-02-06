package main

import "fmt"

type Gamer struct {
	Name string
	Games []string
}

func (gamer *Gamer) AddGame(name string){
	gamer.Games = append(gamer.Games, name)
}

func main() {
	gamer := Gamer{Name:"Piter"}
	gamer.AddGame("Mario")
	gamer.AddGame("Pes 2020")
	for _, val := range gamer.Games {
		fmt.Println(val)
	}
}