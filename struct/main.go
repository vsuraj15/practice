package main

import "fmt"

type User struct {
	name, email, role string
	age               int
}

// Here we are just copying variable User in each case
func (u User) salary() int {
	switch u.role {
	case "engineer":
		return 100
	case "developer":
		return 150
	case "manager":
		return 1000
	default:
		return 0
	}
}

// Here we are just copying variable User in each case, This will not able to modify the value of existing struct
func (u User) updateAge() {
	switch u.role {
	case "engineer":
		u.age = 40
	case "developer":
		u.age = 45
	case "manager":
		u.age = 50
	}
}

// Here we are using pointer, Which means we are able to modify the fields of existing structure.
func (u *User) updateRoleAge() {
	switch u.role {
	case "engineer":
		u.age = 40
	case "developer":
		u.age = 45
	case "manager":
		u.age = 50
	}
}

func main() {
	var user1 = User{name: "Suraj", email: "suraj@gmail.com", age: 33, role: "engineer"}
	var user2 = User{name: "Gagan", email: "gagan@gmail.com", age: 30, role: "developer"}
	var user3 = User{name: "Akash", email: "akash@gmail.com", age: 29, role: "manager"}
	fmt.Printf("User1: %+v, Salary: %+v\n", user1.name, user1.salary())
	fmt.Printf("User2: %+v, Salary: %+v\n", user2.name, user2.salary())
	fmt.Printf("User3: %+v, Salary: %+v\n", user3.name, user3.salary())
	user1.updateAge()
	user2.updateAge()
	user3.updateAge()
	fmt.Printf("User1: %+v\n", user1)
	fmt.Printf("User2: %+v\n", user2)
	fmt.Printf("User3: %+v\n", user3)
	user1.updateRoleAge()
	user2.updateRoleAge()
	user3.updateRoleAge()
	fmt.Printf("User1: %+v\n", user1)
	fmt.Printf("User2: %+v\n", user2)
	fmt.Printf("User3: %+v\n", user3)
}
