package main

import (
	"15-struct/management"
	"fmt"
)

func main() {
	// Struct [Slice of struct]
	peoples := []management.User{
		{1,"Piter", "piter", "piter@gmail.com", true},
		{2,"Park", "park", "pakr@gmail.com", true},
		{3,"Ogik", "ogik", "ogik@gmail.com", false},
	}

	for _, val := range peoples {
		fmt.Println(val)
	}

	user := management.User{1,"Piter", "park", "piter@gmail.com", true}
	user2 := management.User{2,"Park", "park", "park@gmail.com", true}

	// fmt.Println(displayUser(user2))
	fmt.Println(user.Display())

	users := []management.User{user, user2}
	group := management.Group{"Aji", user, users, true}
	
	// displayGroup(group)
	group.Group()
}
