package main

import "fmt"

func main() {
	location := NewLocation("Chisinau", 10, 10)
	//simulator := NewSimulator([]*Location{location})

	person := NewPerson(18, MALE, false)
	location.SetPosition(person, Position{5, 3})
	fmt.Printf("Location for %v is %v", person, location.GetPosition(person))
}
