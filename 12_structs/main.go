package main

import (
	"fmt"
	"strconv"
)

// Define person structs
type Person struct {
	firstName,
	lastName,
	city,
	gender string
	age int
}

func(p *Person) hasBirthday() {
	p.age++
}

func(p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return 
	} else {
		p.lastName = spouseLastName
	}
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func main() {
		// init Person struct
		person1 := Person{
			firstName:"Marcello", 
			lastName:"Silva", 
			city:"Goiânia", 
			gender:"F",
			age:23,
		}

		// person2 := Person{"Teste","Silva","Goiânia","F",23}

		// fmt.Println(person1)
		// fmt.Println(person1.firstName)
		// fmt.Println(person2)
		person1.hasBirthday()
		person1.getMarried("Paquetá")
		fmt.Println(person1.greet())
}
