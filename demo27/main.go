package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name    string
	Age     int
	Emails  []string
	Extra   map[string]string
	UserKey string `mapstructure:"user-key"`
	UserID  string `mapstructure:"user-id"`
}

func main() {
	toMap()
	toStruct()
}

// map -> struct
func toMap() {
	person := Person{
		Name:    "2233",
		Age:     11,
		Emails:  []string{"a@a.com", "b@b.com", "c@c.com"},
		Extra:   map[string]string{"A": "1", "B": "2", "C": "3"},
		UserKey: "ssssss",
		UserID:  "qw8cji28xjw2",
	}
	output := make(map[string]interface{})
	err := mapstructure.Decode(person, &output)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)

}

// struct -> map
func toStruct() {
	input := map[string]interface{}{
		"name":     "lili",
		"age":      18,
		"emails":   []string{"a@a.com", "b@b.com", "c@c.com"},
		"extra":    map[string]string{"A": "1", "B": "2", "C": "3"},
		"user-key": "ssssss",
		"userID":   "qw8cji28xjw2",
	}
	var person Person
	err := mapstructure.Decode(input, &person)
	if err != nil {
		panic(err)
	}
	fmt.Println(person)
}
