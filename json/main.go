package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {

	// struct -> json(string) -> stuct

	// encoding
	fmt.Println("this is json encoding and decoding")
	person := Person{Name: "sharad", Age: 20, IsAdult: true}
	fmt.Println("person: ", person)
	fmt.Printf("type: %T\n", person) // main.Person

	// convert person into JSON encoding: Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println("person: ", string(jsonData))
	fmt.Printf("type: %T\n", jsonData) // []uint8
	fmt.Println("person struct name: ", string(jsonData))

	// decoding
	var a_person Person
	// it takes array of bytes and address where we have to store data
	err = json.Unmarshal(jsonData, &a_person)
	if err!=nil {
		panic(err)
	}
	fmt.Println(a_person.Name)
	fmt.Println(a_person.Age)
	fmt.Println(a_person.IsAdult)

}
