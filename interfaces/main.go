package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorrila struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephard",
	}

	gorrila := Gorrila{
		Name:          "Jack",
		Color:         "gray",
		NumberOfTeeth: 38,
	}

	PrintInfo(&dog)
	PrintInfo(&gorrila)

}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "legs.")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorrila) Says() string {
	return "Ugh"
}

func (g *Gorrila) NumberOfLegs() int {
	return 2
}
