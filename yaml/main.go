package main

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"log"
)

type Config struct {
	Db DBConfig `yaml:"db"`
}

type DBConfig struct {
	Url      string `yaml:"url"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

//db:
//	url: string
//	user: string
//	password: string

func main() {
	marshal()
	unmarshal()
}

func marshal() {
	config := Config{
		Db: DBConfig{
			Url:      "postgres://localhost:5432",
			User:     "user",
			Password: "password",
		},
	}

	bytes, err := yaml.Marshal(config)
	if err != nil {
		log.Printf("error while marshalling: %s\n", err)
	}
	fmt.Println(string(bytes))
}

func unmarshal() {
	bytes := []byte(
		`db:
  url: postgres://localhost:5432
  user: user
  password: password`)
	var config *Config
	err := yaml.Unmarshal(bytes, config)
	if err != nil {
		log.Printf("error while unmarshalling: %s\n", err)
	}
	fmt.Println(string(bytes))
}
