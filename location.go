package main

type Position struct {
	x, y uint32
}

type Location struct {
	name             string
	positions        map[Position]*Person
	limit_x, limit_y uint32
}

func NewLocation(name string, limit_x, limit_y uint32) *Location {
	return &Location{name: name, limit_x: limit_x, limit_y: limit_y, positions: make(map[Position]*Person)}
}

func (location *Location) SetPosition(person *Person, position Position) bool {
	if position.x > location.limit_x || position.y > location.limit_y {
		return false
	}

	if _, ok := location.positions[position]; ok {
		return false
	}

	location.positions[position] = person
	return true
}

func (location *Location) GetPosition(person *Person) *Position {
	for key, value := range location.positions {
		if value == person {
			return &key
		}
	}
	return nil
}
