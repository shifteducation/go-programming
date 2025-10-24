package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	marshal()
	unmarshal()
}

func marshal() {
	user := User{
		//Name:        "John",
		Age:         12,
		PhoneNumber: "5555",
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("error while marshalling user: %s", err)
	}

	log.Printf("user: %s\n", string(bytes))
}

func unmarshal() {
	bytes, err := os.ReadFile("data/test.json")
	if err != nil {
		log.Fatal(err)
	}

	var user User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		log.Printf("error while unmarshalling user: %s", err)
	}

	fmt.Println(user)
}
