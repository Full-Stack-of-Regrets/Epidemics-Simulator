package main

import "fmt"

type Gender string

const (
	MALE   Gender = "MALE"
	FEMALE Gender = "FEMALE"
)

type Person struct {
	age        uint
	gender     Gender
	vaccinated bool
}

func NewPerson(age uint, gender Gender, vaccinated bool) *Person {
	return &Person{age: age, gender: gender, vaccinated: vaccinated}
}

func (person *Person) Recover() {

}

func (person *Person) Infect(target *Person) {
	fmt.Printf("%v: I am infecting %v", person, target)
}
