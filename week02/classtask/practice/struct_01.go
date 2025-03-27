package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func newPerson(name string, age int, email string) *Person {
	person := Person{
		Name:  name,
		Age:   age,
		Email: email,
	}
	return &person
}

func printPerson(person *Person) {
	fmt.Printf("%s %d %s\n", person.Name, person.Age, person.Email)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func main() {
	person := newPerson("李华", 18, "2039114086@qq.com")
	printPerson(person)
}
