package main

type Simulator struct {
	tick      uint32
	locations []*Location
}

func NewSimulator(locations []*Location) *Simulator {
	return &Simulator{tick: 0, locations: locations}
}
