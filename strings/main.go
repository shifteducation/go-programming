package main

import (
	"fmt"
	"strings"
)

type User struct {
	name string
	age  uint
}

func (user User) String() string {
	return fmt.Sprintf("[name: %s, age: %d]", user.name, user.age)
}

func main() {
	//  проверка  наличия  подстроки
	//Contains(s,  substr  string)  bool

	// строка  начинается  с  ?
	//HasPrefix(s,  prefix  string)  bool

	//  склейка  строк
	//Join(a  []string,  sep  string)  string

	//  разбиение  по  разделителю
	//Split(s,  sep  string)  []string
	str := "Hello world"
	fmt.Println(strings.Contains(str, "llo"))
	fmt.Println(strings.HasPrefix(str, "He"))
	fmt.Println(strings.Join([]string{str, "!"}, ""))
	fmt.Println(strings.Split(str, ""))
	fmt.Println(len(str))

	user := User{"John", 34}
	fmt.Printf("%v", user)
}
