package main

import "fmt"

type User struct {
	name string
	age  uint
}

func (user User) String() string {
	return fmt.Sprintf("[name: %s, age: %d]", user.name, user.age)
}

func main() {
	user := User{"John", 34}
	fmt.Printf("%v", user)
}
