package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address,omitempty"`
}

func TestEmployee() {
	person := Person{
		ID:        123,
		FirstName: "Jerry",
	}
	data, err := json.Marshal(person)
	if err != nil {
		log.Panicf("fail to marshal, err = %s\n", err)
	}
	fmt.Printf("data => %s\n", data)

	var person2 Person
	err = json.Unmarshal(data, &person2)
	if err != nil {
		log.Panicf("fail to unmarshal, err = %s\n", err)
	}
	fmt.Printf("person2 => %+v\n", person2)
}
