package management

import "fmt"

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func (user User) Display() string {
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}

func DisplayUser(user User) string {
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvialable bool
}

func (group Group) Group() {
	fmt.Printf("Name: %s\n", group.Name)
	fmt.Printf("Member Count: %d\n", len(group.Users))

	for _, val := range group.Users {
		fmt.Println(val.FirstName)
	}
}

func DisplayGroup(group Group) {
	fmt.Printf("Name: %s\n", group.Name)
	fmt.Printf("Member Count: %d\n", len(group.Users))

	for _, val := range group.Users {
		fmt.Println(val.FirstName)
	}
}