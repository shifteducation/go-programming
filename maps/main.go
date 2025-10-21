package main

import "fmt"

type User struct {
	Name string
	Host string
}

func main() {
	cache := make(map[string]string, 100)
	fmt.Println(len(cache))
	cache["a"] = "b"
	fmt.Println(len(cache))

	users := []User{
		{"John", "127.0.0.1"},
		{"Elvis", "196.198.0.1"},
		{"Antony", "196.198.0.2"},
		{"Samuel", "196.198.0.2"},
	}
	hostUsers := make(map[string][]string)
	for _, user := range users {
		hostUsers[user.Host] = append(hostUsers[user.Host], user.Name)
	}
	fmt.Println(hostUsers)

	iterateOverMap()
}

func iterateOverMap() {
	m := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
	}

	for key := range m {
		if key%2 == 0 {
			delete(m, key)
		}
	}

	fmt.Println(m)
}
